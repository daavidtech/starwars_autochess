extends Area

onready var shape = $Shape

var team = null


# Called when the node enters the scene tree for the first time.
func _ready():
	pass # Replace with function body.

func set_team(team: int):
	team = team

func get_origin():
	return transform.origin
