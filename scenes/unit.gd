extends Spatial

export var attack_rate: int
export var move_speed: float

var moving: bool
var move_dest: Vector3
var direct: Vector3

func stop():
	moving = false
	move_dest = Vector3()
	direct = Vector3()

func start_move(dest: Vector3):
	print("start_move ", dest, translation)
	
	moving = true
	
	move_dest = dest
	
	direct = move_dest - translation
	
	print("driect ", direct)

func move_to_position(dest: Vector3):
	moving = false
	
	move_dest = Vector3()
	direct = Vector3()
	
	translation = dest

func _physics_process(delta):
	if moving:
		var step_size = delta * move_speed

		if step_size > direct.length():
			print("unit reached destination")
			
			step_size = direct.length()
			
			var step = direct.normalized() * move_speed
			
			moving = false
			move_dest = Vector3()
			direct = Vector3()
	
		if move_dest.x > translation.x and translation.x + step_size > move_dest.x:
			direct.x = 0
			
		if move_dest.z > translation.z and translation.z + step_size > move_dest.z:
			direct.z = 0
			
		if move_dest.x < translation.x and translation.x - step_size < move_dest.x:
			direct.x = 0
		
		if move_dest.z < translation.z and translation.z - step_size < move_dest.z:
			direct.z = 0
			
		#print("move_direction ", direct)

		translation += direct.normalized() * step_size
