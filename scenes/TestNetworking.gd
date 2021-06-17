extends Node

var GameState = preload("res://scenes/game_state.gd")
var Unit = preload("res://scenes/unit.tscn")

signal unit_choosen

onready var placement_area = $StaticBody/PlacementArea
onready var unit_shop = $unit_shop

var unit_map = {}

var game_state

func load_thing(type: String):
#	var path = "res://assets/exported/"
#
#	path += type + "/" + type + ".gltf"
	
	var path = "res://assets/unit_droid/unit_droid.glb"
	
	print("Loading resource from ", path)
	
	var does_exists = ResourceLoader.has(path)
	
	print("does_exists " + String(does_exists))
	
	var Droid = ResourceLoader.load(path)
	
	var i = Droid.instance()
	
	add_child(i)
	

func move_to_point(u, x: int, y: int):
	print("move_to ", x, " y ", y)
	
	var size = placement_area.shape.extents
	
	var start_x = -size.x
	var start_y = -size.z
	
	print("start_x ", start_x)
	print("start_y ", start_y)
	
	var x_ratio = size.z / game_state.width
	var y_ratio = size.x / game_state.height
	
	var translation_z = start_x + x_ratio * x
	var translation_x = start_y + y_ratio * y
	
	print("translation_x ", translation_x)
	print("translation_z ", translation_z)
	
	u.translation.x = translation_x
	u.translation.z = translation_z
	u.translation.y = placement_area.translation.y

func _ready():
	game_state = GameState.new()
	add_child(game_state)
	
	print(placement_area.shape.extents.x)
	
	game_state.connect("create_unit", self, "_handle_create_unit")
	game_state.connect("unit_position_changed", self, "_handle_change_unit_position")
	
	load_thing("unit_clone")
	
	unit_shop.fill([{
		"id": 1,
		"unit_type": "unit_droid",
		"level": 3,
		"rank": 3,
		"cost": 300
	}, {
		"id": 2,
		"unit_type": "unit_clone",
		"level": 2,
		"rank": 1,
		"cost": 420
	}, {
		"id": 3,
		"unit_type": "unit_droid",
		"level": 2,
		"rank": 2,
		"cost": 500
	}])
	
	unit_shop.connect("unit_bought", self, "_handle_unit_bought")

func _handle_unit_bought(unit):
	print("unit bought ", unit)

func _handle_create_unit(id, unit_type, x, y):
	print("handle create unit")

	var new_unit = Unit.instance()
	
	move_to_point(new_unit, x, y)
	
	placement_area.add_child(new_unit)
	unit_map[id] = new_unit

func _handle_change_unit_position(id, x: int, y: int):
	var unit = unit_map.get(id)
	
	move_to_point(unit, x, y)
