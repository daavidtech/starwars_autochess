extends "res://scenes/unit.gd"

onready var friendly_bars = $friendly_bars
onready var rank_indicator = $rank
onready var place_holder = $PlaceHolder

var location

var mouse_over = false
var dragging = false

signal drag_started(unit)
signal drag_finished(unit)
	
func _ready():
	place_holder.visible = false	

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


func _unhandled_input(event):
	if mouse_over and event.is_action_pressed("left_mouse_button"):
		emit_signal("drag_started", self)
		
	if dragging and event.is_action_released("left_mouse_button"):
		emit_signal("drag_finished", self)

func _on_Spatial_mouse_entered():
	mouse_over = true

func _on_Spatial_mouse_exited():
	mouse_over = false
