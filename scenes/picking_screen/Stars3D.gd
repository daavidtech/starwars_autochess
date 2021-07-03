extends Spatial

onready var stars2d = $Viewport/Stars2D

export var value: int setget set_value, get_value

func set_value(v: int):
	if stars2d:
		stars2d.stars = v

func get_value():
	return stars2d.stars
