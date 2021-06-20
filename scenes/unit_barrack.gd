extends Spatial

onready var shape = $shape

func add_unit(unit):
	shape.add_child(unit)
	
func remove_unit(unit):
	shape.remove_child(unit)
	
func clear():
	for unit in shape.get_children():
		shape.remove_child(unit)
