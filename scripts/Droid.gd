extends "res://scripts/Unit.gd"

var mouse_is_over = false
var drag_button_already_down = false
var dragging = false

var direction = null

onready var parent_coordinates = get_parent().translation

# Called when the node enters the scene tree for the first time.
func _ready():
	pass # Replace with function body.

# func _physics_process(delta):
# 	if direction != null:
# 		var step_size = delta * speed
		
# 		apply_central_impulse(direction * step_size)
	
# 	if drag_button_already_down == false and mouse_is_over and Input.is_action_pressed("left_mouse_button"):
# 		if dragging == false:
# 			dragging = true
# 			#custom_integrator = true
# 			#mode = MODE_STATIC
# 			apply_central_impulse(Vector3(0, 1, 0))
# 			print("start dragging")
		
# 	if Input.is_action_pressed("left_mouse_button") == false and dragging == true:
# 		dragging = false
# 		drag_button_already_down = false
# 		apply_central_impulse(Vector3(0,0,0))
# 		#mode = MODE_RIGID
# 		#custom_integrator = false
# 		print("dragging stopped")
		
#		var viewport = get_viewport()
#
#		var camera = viewport.get_camera()
#		var mouse_position = viewport.get_mouse_position()
#
#		var from = camera.project_ray_origin(mouse_position)
#		var to = from + camera.project_ray_normal(mouse_position) * 200
#
##		print("from", from)
##		print("to", to)
#
#		var direct_state = get_world().direct_space_state
#
#		var collision = direct_state.intersect_ray(from, to)
#
#		if collision:
#
#			#print("collision", collision)
#
#			translation.x = collision.position.x
#			translation.z = collision.position.z
#			translation.y = collision.position.y + 0.01
		
#	if dragging:
#		var viewport = get_viewport()
#
#		var camera = viewport.get_camera()
#		var mouse_position = viewport.get_mouse_position()
#
##		var camera = get_viewport().get_camera()
##		print(mouse_position)
#
#		var from = camera.project_ray_origin(mouse_position)
#		var to = from + camera.project_ray_normal(mouse_position) * 200
#
##		print("from", from)
##		print("to", to)
#
#		var direct_state = get_world().direct_space_state
#
#		var collision = direct_state.intersect_ray(from, to)
#
#		if collision:
#
#			#print("collision", collision)
#
#			translation.x = collision.position.x
#			translation.z = collision.position.z
#			translation.y = collision.position.y + 1

#	var camera = get_viewport().get_camera()
#	var direct_state = get_world().direct_space_state
#	var collision = direct_state.intersect_ray(camera.translation, Vector3(0, 0, -20))
#
#	if collision:
#		print("collision", collision)

func _integrate_forces(state):
	if dragging:		
		var viewport = get_viewport()
		
		var camera = viewport.get_camera()
		var mouse_position = viewport.get_mouse_position()
		
		var from = camera.project_ray_origin(mouse_position)
		var to = from + camera.project_ray_normal(mouse_position) * 200
		
		var direct_state = get_world().direct_space_state
		
		var collision = direct_state.intersect_ray(from, to, [], 5)
		
		if collision:
			if collision.collider.is_in_group("Characters"):
				return
			
			var xform = state.get_transform()
			
			xform.origin.x = collision.position.x
			xform.origin.z = collision.position.z
			xform.origin.y = collision.position.y + 5
			
			state.set_transform(xform)
		

# Called every frame. 'delta' is the elapsed time since the previous frame.
#func _process(delta):
#	pass

func attack():
	var collision = find_closest_enemy()

	var direct_state = get_world().direct_space_state
	
	var query = PhysicsShapeQueryParameters.new()
	
	var s = SphereShape.new()
	s.radius = 100
	
	query.set_shape(s)
	query.collision_mask = 256
	
	var collisions = direct_state.intersect_shape(query)
	
	print("collisions ", collisions)
	
	var first_collision = collisions[2]

	if first_collision:
		print("first attack collision", first_collision)
		
		var collider = first_collision.collider
		
		print("collider ", collider.transform.origin)
		
#		move_to_point(collider.transform.origin.x, collider.transform.origin.z)
		

# func move_to_point(x, z):
# 	var destination = Vector3(x, transform.origin.y, z)
	
# 	direction = destination - translation

func _on_Droid_mouse_entered():
	mouse_is_over = true
	if Input.is_action_pressed("left_mouse_button"):
		drag_button_already_down = true
	print("tickling")


func _on_Droid_mouse_exited():
	mouse_is_over = false
	drag_button_already_down = false
	print("stop tickling")
