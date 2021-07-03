extends "res://scenes/unit.gd"

onready var friendly_bars = $friendly_bars
onready var rank_indicator = $rank
	
var unit_id: String

var _unit_type: String

export var unit_type: String setget set_unit_type, get_unit_type

func set_unit_type(t: String):
	_unit_type = t
	
func get_unit_type() -> String:
	return _unit_type

export var hp: int setget set_hp, get_hp

func set_hp(value: int):
	friendly_bars.hp = value
	
func get_hp() -> int:
	return friendly_bars.hp

export var mana: int setget set_mana, get_mana

func set_mana(value: int):
	friendly_bars.mana = value
	
func get_mana() -> int:
	return friendly_bars.mana 

export var rank: int setget set_rank, get_rank

func set_rank(v: int):
	rank_indicator.value = v

func get_rank():
	return rank_indicator.value
