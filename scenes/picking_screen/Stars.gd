extends Node2D

onready var label = $Label

export var stars: int setget set_stars, get_stars

func set_stars(v: int):
	var text = ""
	
	for i in range(v):
		text += "X"
		
	label.text = text

func get_stars():
	print("get_stars ", label.text, " ", label.text.length())
	
	return label.text.length()
