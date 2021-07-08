extends Node2D

onready var hp_bar = $HPBar

var hp setget set_hp, get_hp

func set_hp(v):
	hp_bar.value = v
	
func get_hp():
	return hp_bar.value
