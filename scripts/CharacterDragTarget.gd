extends Area

onready var mouse_over_indicator = $MouseOverIndicator

# Declare member variables here. Examples:
# var a = 2
# var b = "text"


# Called when the node enters the scene tree for the first time.
func _ready():
	pass # Replace with function body.


# Called every frame. 'delta' is the elapsed time since the previous frame.
#func _process(delta):
#	pass

func show_location():
	print("show_location")

func _on_CharacterDragTarget_mouse_entered():
	mouse_over_indicator.visible = true


func _on_CharacterDragTarget_mouse_exited():
	mouse_over_indicator.visible = false
