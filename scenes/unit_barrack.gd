extends Spatial

onready var shape = $shape

func add_unit(unit):
	shape.add_child(unit)
	
	var size = shape.shape.extents
	
	var start_x = -size.z
	var start_y = -size.z
	
	var one_width = (size.z * 2) / 9
	var one_half_width = one_width / 2
	
	var count = shape.get_child_count()
	
	var x = start_x + one_width * count + one_half_width
	
	unit.translation.z = x
	
func remove_unit(unit):
	shape.remove_child(unit)
	
func clear():
	for unit in shape.get_children():
		shape.remove_child(unit)
