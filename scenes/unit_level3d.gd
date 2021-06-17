extends Spatial

onready var level = $Viewport/level

export var value: int setget set_value, get_value

func set_value(v: int):
	level.value = v
	
func get_value() -> int:
	return level.value
