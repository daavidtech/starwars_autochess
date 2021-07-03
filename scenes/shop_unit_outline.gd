extends StaticBody

onready var shape = $shape
onready var rank = $rank
onready var level = $level
onready var cost = $cost

signal unit_choosen(props)

func _ready():
	rank.value = 3

var properties

var mouse_over = false

func set_properties(props):
	properties = props
	
	rank.value = props.rank
	level.value = props.level
	cost.value = props.cost

func set_content(c):
	shape.add_child(c)

func _process(delta):
	if Input.is_action_just_pressed("left_mouse_button") and mouse_over == true:
		print("choose this")
		emit_signal("unit_choosen", properties)

func _on_Spatial_mouse_entered():
	mouse_over = true


func _on_Spatial_mouse_exited():
	mouse_over = false
