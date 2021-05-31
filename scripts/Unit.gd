extends KinematicBody

export var speed = 10
export var attack_range = 1

var move_destination = null
var move_direction = null

var game_grid = null

func set_game_grid(g):
	game_grid = g

func _physics_process(delta):
	if move_direction != null:
		handle_move(delta)

func handle_move(delta):
	var step_size = delta * speed

	var stop_now = false

	if step_size > move_direction.length():
		print("unit reached destination")
		
		step_size = move_direction.length()
		stop_now = true
		
	print("translation.x ", translation.x, " destination.x ", move_destination.x)	
	
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
	print("stopping now")

func move(x, y):
	print("moving to point")

	move_destination = Vector3(y, transform.origin.y, x)

	move_direction = move_destination - translation
	
	
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


