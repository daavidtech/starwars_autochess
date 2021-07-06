extends Spatial

onready var shape = $shape

onready var left = $left
onready var right = $right

var unit_index = {}
var index_unit = {}

func _ready():
	remove_child(left)
	remove_child(right)

func add_unit(unit):
	var size = shape.shape.extents
	
	var start_z = -size.z
	
	var one_width = (size.z * 2) / 9
	var one_half_width = one_width / 2
	
	print("add_unit " + unit.unit_id + " to barrack")
	
	for i in range(9):
		if index_unit.has(i):
			continue
			
		index_unit[i] = unit
		unit_index[unit] = i
		
		var z = start_z + one_width * i + one_half_width
		
		shape.add_child(unit)

		unit.translation.x = 0
		unit.translation.y = shape.translation.y		
		unit.translation.z = z
		unit.rotation.z = -62
		
		break

func remove_unit(unit):
	print("Removing unit " + unit.unit_id + " from barrack")
	
	var index = unit_index[unit]
	index_unit.erase(index)
	unit_index.erase(unit) 
	shape.remove_child(unit)
	
func clear():
	for unit in shape.get_children():
		shape.remove_child(unit)
		var index = unit_index[unit]
		index_unit.erase(index)
		unit_index.erase(unit)
