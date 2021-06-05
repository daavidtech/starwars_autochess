extends Node

signal someone_won(team)
signal draw

var x = null
var y = null

var units = []

func get_units():
	return units
	
func count_units():
	return units.size()

func add_unit(unit):
	print("adding unit ", unit.get_instance_id())
	
	if units.has(unit):
		return
	
	units.push_back(unit)

func remove_unit(unit):
	units.erase(unit)
	
	if units.size() == 0:
		emit_signal("draw")
		
		return
		
	if is_there_one_team_left() == false:
		return
		
	var unit_left = units[0]
		
	emit_signal("someone_won", unit_left.team)

func is_there_one_team_left():
	var team = null
	
	for unit in units:
		if team == null:
			team = unit.team
			continue
		if team != unit.team:
			return false
			
	return true

func set_size(x, y):
	x = x
	y = y
