extends Spatial

var Droid = preload("res://scenes/droid/Droid.tscn")
var PathCoordinator = preload("res://scripts/PathCoordinator.gd")

onready var game_map = $GameMap
onready var game_start_timer = $GameStartTimer
onready var start_time_left = $Control/StartTimeLeft
onready var game_grid = $GameMap/GameGrid

onready var path_coordinator = PathCoordinator.new()

var your_units = []
var enemy_units = []

var start_timer_tics = 0
export var game_start_time = 3

const YOU = 1
const ENEMY = 2

func _ready():
	start()

func spawn_unit(x: int, y: int, Unit, team: int):
	var new_unit = Unit.instance()
	
	new_unit.translation.y = game_map.translation.y
	new_unit.set_position(x, y)
	
	new_unit.team = team
	new_unit.set_path_coordinator(path_coordinator)
	new_unit.connect("death", self, "_on_unit_death")
	path_coordinator.add_unit(new_unit)
	
	game_map.add_child(new_unit)
#
#	new_unit.attack()
	
	return new_unit

func start():
	#Logger.info()
	
	game_start_timer.start()
	
	start_time_left.text = String(game_start_time)
	
	var size = game_map.get_size()
	
	var d = spawn_unit(0, 0, Droid, YOU)
	
	#d.move(-1, 1)
	d.attack()
#
	var d1 = spawn_unit(8, 7, Droid, ENEMY)
	d1.move(8, 0)
	var d2 = spawn_unit(7, 7, Droid, ENEMY)
	d2.move(7, 0)
	var d3 = spawn_unit(6, 7, Droid, ENEMY)
	d3.move(6, 0)
	var d4 = spawn_unit(5, 7, Droid, ENEMY)
	d4.move(5, 0)
	var d5 = spawn_unit(4, 7, Droid, ENEMY)
	d5.move(4, 0)
#
#	spawn_unit(0, 2, Droid, ENEMY)
#
#	spawn_unit(1, 3, Droid, ENEMY)
#
#	spawn_unit(2, 1, Droid, ENEMY)

func _on_unit_death(unit):
	print("on unit death")
	
	game_map.remove_child(unit)
	path_coordinator.remove_unit(unit)

func _on_GameStartTimer_timeout():
	start_timer_tics += 1
	
	start_time_left.text = String(game_start_time - start_timer_tics)
	
	if start_timer_tics >= game_start_time:
		game_start_timer.stop()
		start_time_left.visible = false
