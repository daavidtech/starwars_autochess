extends Node2D

onready var health_bar = $health_bar
onready var mana_bar = $mana_bar

var _max_hp: int

export var max_hp: int setget set_max_health, get_max_health

func set_max_health(value: int):
	_max_hp = value

func get_max_health():
	return _max_hp

export var hp: int setget set_hp, get_hp

func set_hp(value: int):
	health_bar.value = value
	
func get_hp() -> int:
	return health_bar.value

export var mana: int setget set_mana, get_mana

func set_mana(value: int):
	mana_bar.value = value
	
func get_mana() -> int:
	return mana_bar.value
