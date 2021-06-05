extends StaticBody

signal clicked

onready var cost_label = $Viewport/CostLabel

export var level = 1
export var unit_type = "droid"
export var cost = 10

var mouse_over = false

func _ready():
	cost_label.set_cost(cost)

func _process(delta):
	if Input.is_action_just_pressed("left_mouse_button") and mouse_over == true:
		emit_signal("clicked", self)

func _on_Spatial_mouse_entered():
	mouse_over = true

func _on_Spatial_mouse_exited():
	mouse_over = false
