extends Spatial

var Droid = preload("res://scenes/Droid.tscn")

onready var game_map = $GameMap
onready var game_start_timer = $GameStartTimer
onready var start_time_left = $Control/StartTimeLeft
onready var game_grid = $GameMap/GameGrid

var your_units = []
var enemy_units = []

var start_timer_tics = 0
export var game_start_time = 3

# Declare member variables here. Examples:
# var a = 2
# var b = "text"

var droids = []

# Called when the node enters the scene tree for the first time.
func _ready():
	print("map scale", game_map.transform)
	
	start()

func spawn_droid(x, z):
	var new_droid = Droid.instance()
	
	new_droid.set_game_grid(game_grid)	
	new_droid.transform.origin.x = x
	new_droid.transform.origin.z = z
	new_droid.transform.origin.y = transform.origin.y
	
	game_map.add_child(new_droid)
	
	droids.push_back(new_droid)
	
func move_all_droids(x, z):
	for droid in droids:
		droid.move_to_point(x, z)

func spawn_unit(x: int, y: int, Unit):
	var new_unit = Unit.instance()
	
	var origin = game_grid.get_column_origin(x, y)
	
	print("origin ", origin)
	
	new_unit.transform.origin.x = origin.x
	new_unit.transform.origin.z = origin.z
	new_unit.transform.origin.y = transform.origin.y
	
	game_grid.add_unit(new_unit)

func start():
	game_start_timer.start()
	
	start_time_left.text = String(game_start_time)
	
	var size = game_map.get_size()
	
	var new_droid = Droid.instance()
	
	new_droid.transform.origin.z = 3
	
	game_map.add_child(new_droid)
	
	new_droid.move(3, -6)
	
#	random_droid = Droid.instance()
#
#	random_droid.transform.origin.x = 0
#	random_droid.transform.origin.z = 0
#	random_droid.transform.origin.y = transform.origin.y + 5
#
#	game_map.add_child(random_droid)

#	spawn_unit(0,0, Droid)
#	spawn_unit(0,1, Droid)
#	spawn_unit(1,0, Droid)
#	spawn_unit(1,1, Droid)


#	var origin = game_grid.get_column_origin(2, 2)
#
#	spawn_droid(origin.x, origin.z)
	
#	var your_offset = -7
#	var enemy_offset = 7
	
#	for i in range(1):
#		for j in range(1):
#			spawn_droid(your_offset + 1*i, 1*j, "you")
#
#	for i in range(1):
#		for j in range(1):
#			spawn_droid(enemy_offset + 1*i,  1*j, "enemy")
	


# Called every frame. 'delta' is the elapsed time since the previous frame.
#func _process(delta):
#	pass


func _on_GameStartTimer_timeout():
	start_timer_tics += 1
	
	start_time_left.text = String(game_start_time - start_timer_tics)
	
	if start_timer_tics >= game_start_time:
		game_start_timer.stop()
		start_time_left.visible = false
		
#		var droid = droids[0]
#
#		droid.move_to_point(1, -2)
		
#		move_all_droids(-3, 2)
		
