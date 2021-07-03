extends Control

onready var name_label = $PlayerName

var player_id = "null123"

var player_name setget set_name, get_name

func set_name(name):
	name_label.text = name
	
func get_name():
	return name_label.text

