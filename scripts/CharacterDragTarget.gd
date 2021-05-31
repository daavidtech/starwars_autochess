extends Area

onready var mouse_over_indicator = $MouseOverIndicator

func show_location():
	print("show_location")

func _on_CharacterDragTarget_mouse_entered():
	mouse_over_indicator.visible = true


func _on_CharacterDragTarget_mouse_exited():
	mouse_over_indicator.visible = false
