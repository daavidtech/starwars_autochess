extends Node

var Unit = preload("res://scripts/Unit.gd")

var x = null
var y = null

var units = {}

func get_units():
	return units.values()
	
func count_units():
	return units.size()

func add_unit(unit):
	print("adding unit ", unit.get_instance_id())
	
	units[unit.get_instance_id()] = unit

func remove_unit(unit):
	units.erase(unit.get_instance_id())

func set_size(x, y):
	x = x
	y = y
