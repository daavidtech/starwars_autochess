extends Spatial

var Droid = preload("res://scenes/Droid.tscn")

onready var game_map = $GameMap
onready var game_start_timer = $GameStartTimer
onready var start_time_left = $Control/StartTimeLeft

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
	start()

func spawn_droid(x, z):
	var new_droid = Droid.instance()
	
	new_droid.transform.origin.x = x
	new_droid.transform.origin.z = z
	new_droid.transform.origin.y = transform.origin.y + 5
	
	game_map.add_child(new_droid)
	
	droids.push_back(new_droid)
	
func move_all_droids(x, z):
	for droid in droids:
		droid.move_to_point(x, z)

func start():
	print("Start")
	game_start_timer.start()
	
	start_time_left.text = String(game_start_time)
	
#	random_droid = Droid.instance()
#
#	random_droid.transform.origin.x = 0
#	random_droid.transform.origin.z = 0
#	random_droid.transform.origin.y = transform.origin.y + 5
#
#	game_map.add_child(random_droid)
	
	var your_offset = -7
	var enemy_offset = 7
	
	for i in range(1):
		for j in range(1):
			spawn_droid(your_offset + 1*i, 1*j)
			
	for i in range(1):
		for j in range(1):
			spawn_droid(enemy_offset + 1*i,  1*j)
	


# Called every frame. 'delta' is the elapsed time since the previous frame.
#func _process(delta):
#	pass


func _on_GameStartTimer_timeout():
	start_timer_tics += 1
	
	start_time_left.text = String(game_start_time - start_timer_tics)
	
	if start_timer_tics >= game_start_time:
		game_start_timer.stop()
		start_time_left.visible = false
		
		var droid = droids[0]
		
		droid.attack()
		
#		move_all_droids(-3, 2)
		
