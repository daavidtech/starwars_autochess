extends Spatial

var Droid = preload("res://scenes/Droid.tscn")
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
	
	new_unit.set_position(x, y)
	
	new_unit.team = team
	new_unit.set_path_coordinator(path_coordinator)
	new_unit.connect("death", self, "_on_unit_death")
	path_coordinator.add_unit(new_unit)
	
	game_map.add_child(new_unit)
	
	return new_unit

func start():
	game_start_timer.start()
	
	start_time_left.text = String(game_start_time)
	
	var size = game_map.get_size()
	
	var your_unit = spawn_unit(0, -2, Droid, YOU)
	
	var your_unit2 = spawn_unit(6, -3, Droid, YOU)
	
	var enemy_unit = spawn_unit(0, 2, Droid, ENEMY)
	
	your_unit.attack()
	enemy_unit.attack()
	your_unit2.attack()

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
