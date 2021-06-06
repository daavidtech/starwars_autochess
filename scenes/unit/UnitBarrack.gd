extends StaticBody

signal on_drag_started(unit)

var Droid = preload("res://scenes/droid/DroidModel.tscn")
var UnitBarrackUnit = preload("res://scenes/barrack/UnitBarrackUnit.tscn")

onready var shape = $Shape

const max_units = 5

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
			
			var index = shape.get_child_count() + 1
			
			barrack_unit.translation.z = get_z_translation(index)

func handle_drag_start(unit):	
	print("UnitBarrack: handle_drag_start")
	emit_signal("on_drag_started", unit)

func apply_option(opt):
	if shape.get_child_count() >= max_units:
		print("Barrack is full")
		
		return
	
	print("apply_opt ", opt.stars.stars)
	
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
	
func show_all():
	for unit in shape.get_children():
		unit.visible = true
	
func hide_unit(unit):
	unit.visible = false

func remove_unit(unit):
	shape.remove_child(unit)

func is_full():
	if shape.get_child_count() >= max_units:
		return true
		
	return false
