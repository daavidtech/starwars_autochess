extends Node2D

onready var label = $label

export var value: int setget set_value, get_value

func set_value(v: int):
	label.text = String(v)
	
func get_value() -> int:
	return int(label.text)
