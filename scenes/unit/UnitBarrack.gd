extends StaticBody

signal on_drag_started(unit)

var Droid = preload("res://scenes/droid/DroidModel.tscn")
var UnitBarrackUnit = preload("res://scenes/barrack/UnitBarrackUnit.tscn")

onready var shape = $Shape

var units = []

#func _ready():

func get_z_translation(index: int):
	var size = shape.shape.extents
	
	print("size ", size)
	
	var part = (size.z * 2) / 8
	
	print("part ", part)
	
	var start_z = -(size.z)
	
	print("start_z ", start_z)
	
	return start_z + part * index

func spawn_unit(opt):
	match opt.unit_type:
		"droid":	
			var barrack_unit = UnitBarrackUnit.instance()
			
			shape.add_child(barrack_unit)
			
			barrack_unit.connect("on_drag_start", self, "handle_drag_start")
			
			var droid = Droid.instance()
			
			barrack_unit.set_content(droid)
			
			var index = units.size() + 1
			
			units.push_back(opt)
			
			barrack_unit.translation.z = get_z_translation(index)

func handle_drag_start(unit):	
	print("UnitBarrack: handle_drag_start")
	emit_signal("on_drag_started", unit)

func apply_option(opt):
	spawn_unit(opt)
	
#	var units_with_same_type = []
#
#	for unit in units:
#		if unit.unit_type == opt.unit_type:
#			units_with_same_type.push_back(unit)
#
#	if units_with_same_type.size() == 3:
#		var unit = units_with_same_type[0]
#
#		units.remove(units_with_same_type[1])
#		units.remove(units_with_same_type[2])
#
#		unit.level += 1
#	else:
#		spawn_unit(opt)
	

func remove_unit(unit):
	shape.remove_child(unit)



