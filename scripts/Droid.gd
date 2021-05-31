extends "res://scripts/Unit.gd"

var mouse_is_over = false
var drag_button_already_down = false
var dragging = false

var direction = null

onready var parent_coordinates = get_parent().translation

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

func attack():
	print("attack")
