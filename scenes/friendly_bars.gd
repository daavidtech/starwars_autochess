extends Spatial

onready var bars = $Viewport/bars

export var hp: int setget set_hp, get_hp

func set_hp(value: int):
	bars.hp = value

func get_hp() -> int:
	return bars.hp

export var mana: int setget set_mana, get_mana

func set_mana(value: int):
	bars.mana = value
	
func get_mana() -> int:
	return bars.mana
