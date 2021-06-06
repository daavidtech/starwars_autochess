extends Spatial

onready var stars2d = $Viewport/Stars2D

export var stars: int setget set_stars, get_stars

func set_stars(v: int):
	stars2d.stars = v

func get_stars():
	return stars2d.stars
