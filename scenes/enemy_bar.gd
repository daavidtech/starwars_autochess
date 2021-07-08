extends Spatial

onready var hp_bar = $Viewport/HPBar2D

var hp setget set_hp, get_hp

func set_hp(v):
	hp_bar.hp = v

func get_hp():
	return hp_bar.hp
