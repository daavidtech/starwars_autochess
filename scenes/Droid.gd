extends RigidBody


# Declare member variables here. Examples:
# var a = 2
# var b = "text"

var mouse_is_over = false
var drag_button_already_down = false
var dragging = false

onready var parent_coordinates = get_parent().translation

# Called when the node enters the scene tree for the first time.
func _ready():
	pass # Replace with function body.

func _physics_process(delta):
	if drag_button_already_down == false and mouse_is_over and Input.is_action_pressed("left_mouse_button"):
		if dragging == false:
			dragging = true
			#custom_integrator = true
			#mode = MODE_STATIC
			apply_central_impulse(Vector3(0, 1, 0))
			print("start dragging")
		
	if Input.is_action_pressed("left_mouse_button") == false and dragging == true:
		dragging = false
		drag_button_already_down = false
		apply_central_impulse(Vector3(0,0,0))
		#mode = MODE_RIGID
		#custom_integrator = false
		print("dragging stopped")
		
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


func _on_Droid_mouse_entered():
	mouse_is_over = true
	if Input.is_action_pressed("left_mouse_button"):
		drag_button_already_down = true
	print("tickling")


func _on_Droid_mouse_exited():
	mouse_is_over = false
	drag_button_already_down = false
	print("stop tickling")
