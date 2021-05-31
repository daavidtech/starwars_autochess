extends KinematicBody

export var health = 100
export var speed = 2
export var attack_range = 1
export var attack_rate = 1

signal death(unit)

var move_destination = null
var move_direction = null

var path_coordinator = null

var current_target = null

var team = null

var since_last_hit = 0

func set_path_coordinator(c):
	path_coordinator = c

func _physics_process(delta):
	since_last_hit += delta
	
	if current_target:
		var my_position = get_position()
		var target_position = current_target.get_position()
		
		if (my_position.x > target_position.x - 1 and my_position.x < target_position.x + 1
			and my_position.y > target_position.y - 1 and my_position.y < target_position.y + 1):
			stop()
			
			if since_last_hit > attack_rate:
				hit_target()
				since_last_hit = 0
			
			if current_target.is_death():
#				print("it is death")
				
				current_target = null
				
				attack_closest_unit()
		else:
			move(target_position.x, target_position.y)
	
	if move_direction != null:
		handle_move(delta)

func handle_move(delta):
	var step_size = delta * speed

	var stop_now = false

	if step_size > move_direction.length():
		print("unit reached destination")
		
		step_size = move_direction.length()
		stop_now = true
	
	if move_destination.x > translation.x and translation.x + step_size > move_destination.x:
		move_direction.x = 0
		
	if move_destination.z > translation.z and translation.z + step_size > move_destination.z:
		move_direction.z = 0
		
	if move_destination.x < translation.x and translation.x - step_size < move_destination.x:
		move_direction.x = 0
	
	if move_destination.z < translation.z and translation.z - step_size < move_destination.z:
		move_direction.z = 0
	
	translation += move_direction.normalized() * step_size

	if move_direction.x == 0 and move_direction.z == 0:
		move_direction = null
		move_destination = null

func stop():
	move_destination = null
	move_direction = null

func move(x, y):
#	print("moving to point")

	move_destination = Vector3(y, transform.origin.y, x)

	move_direction = move_destination - translation
	
func set_position(x, y):
	translation.x = y
	translation.z = x
	
func get_position():
	return {
		"x": translation.z,
		"y": translation.x
	}
	
func set_current_target(t):
	current_target = t
	
func find_closest_enemy():
	var direct_state = get_world().direct_space_state
	
	var query = PhysicsShapeQueryParameters.new()
	
	var s = SphereShape.new()
	s.radius = 100
	
	query.set_shape(s)

	var collisions = direct_state.intersect_ray(query)

	for collision in collisions:
		if collision.collider.has_group("enemy"):
			return collision.collider


func attack_closest_unit():
	var units = path_coordinator.get_units()
	
	for unit in units:
		if unit != self and unit.team != team:
			set_current_target(unit)
			return

func take_damage(amount: int):
	health -= amount
	
	if health <= 0:
		print("It is deadth")
		emit_signal("death", self)

func is_death():
	return health <= 0

func hit_target():
	if current_target:
		current_target.take_damage(30)
