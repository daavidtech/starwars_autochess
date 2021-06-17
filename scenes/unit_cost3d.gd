extends Spatial

onready var cost = $Viewport/cost

export var value: int setget set_value, get_value

func set_value(v: int):
	cost.value = v
	
func get_value() -> int:
	return cost.value
