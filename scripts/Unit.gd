extends KinematicBody

export var speed = 10
export var attack_range = 1

var move_destination = null
var move_direction = null

var game_grid = null

func set_game_grid(g):
	game_grid = g

# Called when the node enters the scene tree for the first time.
func _ready():
	pass # Replace with function body.

func _physics_process(delta):
	if move_direction != null:
		handle_move(delta)

func handle_move(delta):
	var step_size = delta * speed

	var stop_now = false
	
	#print("move_direction length ", move_direction.length())
	
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

	# print(
	# 	"x: ", 
	# 	move_destination.x, 
	# 	" x2: ",
	# 	transform.origin.x, 
	# 	" z: ", 
	# 	move_destination.z, 
	# 	" z2: ", 
	# 	transform.origin.z)

	# if move_destination.x <= transform.origin.x:
	# 	var inverse = Vector3(-move_destination.x, 0, 0)

	# 	apply_central_impulse(inverse)

	# 	move_destination.x = 0

	# if move_destination.z >= transform.origin.z:
	# 	var inverse = Vector3(0, 0, -move_destination.z)
		
	# 	apply_central_impulse(inverse)

	# 	move_destination.z = 0

	# print("move_destination", move_destination, transform.origin)

	# if move_destination.x == 0 and move_destination.z == 0:
	# 	print("Both are zero")

	# 	move_direction = null
	# 	move_destination = null

	# 	return

	# if move_destination.x == transform.origin.x and move_destination.z == transform.origin.z:
	# 	print("stop moving")

	# 	apply_central_impulse(move_direction.inverse() * step_size)


		
	# 	return


	# apply_central_impulse(move_direction * step_size)

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


