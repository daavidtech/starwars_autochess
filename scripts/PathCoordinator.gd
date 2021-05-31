extends Node

var Unit = preload("res://scripts/Unit.gd")

var x = null
var y = null

var units = {}

func get_units():
	return units.values()

func add_unit(unit):
	units[unit.get_instance_id()] = unit

func remove_unit(unit):
	units.erase(unit.get_instance_id())

func set_size(x, y):
	x = x
	y = y

func commit_path():
	pass
