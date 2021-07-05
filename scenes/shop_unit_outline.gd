extends StaticBody

onready var shape = $shape
onready var rank_label = $rank
onready var level_label = $level
onready var cost_label = $cost

var id

var rank: int setget set_rank, get_rank

func set_rank(rank):
	rank_label.value = rank
	
func get_rank():
	return rank_label.value

var level: int setget set_level, get_level

func set_level(level):
	level_label.value = level
	
func get_level():
	return level_label.value

var cost: int setget set_level, get_level

func set_cost(cost):
	cost_label.value = cost

func get_cost():
	return cost_label.value

signal unit_choosen(props)

func _ready():
	pass

var mouse_over = false

func set_content(c):
	shape.add_child(c)

func _process(delta):
	if Input.is_action_just_pressed("left_mouse_button") and mouse_over == true:
		print("choose this " + String(id))
		emit_signal("unit_choosen", id)

func _on_Spatial_mouse_entered():
	mouse_over = true


func _on_Spatial_mouse_exited():
	mouse_over = false
