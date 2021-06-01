extends "res://scripts/Unit.gd"

onready var health_bar = $HealthBar

var mouse_is_over = false
var drag_button_already_down = false
var dragging = false

var direction = null

func _ready():
	self.connect("health_changed", health_bar, "set_value")

func attack():
	attack_closest_unit()
	
