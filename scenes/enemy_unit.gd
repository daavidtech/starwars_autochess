extends "res://scenes/unit.gd"

onready var enemy_bars = $enemy_bars
onready var rank_indicator = $rank
onready var place_holder = $PlaceHolder

func _ready():
	place_holder.visible = false

export var hp: int setget set_hp, get_hp

func set_hp(value: int):
	enemy_bars.hp = value
	
func get_hp() -> int:
	return enemy_bars.hp

export var rank: int setget set_rank, get_rank

func set_rank(v: int):
	rank_indicator.value = v

func get_rank():
	return rank_indicator.value
