extends "res://scripts/Unit.gd"

var mouse_is_over = false
var drag_button_already_down = false
var dragging = false

var direction = null

onready var parent_coordinates = get_parent().translation

func attack():
	attack_closest_unit()
	
