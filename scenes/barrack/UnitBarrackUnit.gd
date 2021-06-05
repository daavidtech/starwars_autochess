extends StaticBody

signal on_drag_start(unit)

onready var shape = $Shape

var mouse_over = false

func _ready():
	print("shape ", shape)
	
	connect("mouse_entered", self, "handle_mouse_entered")
	connect("mouse_exited", self, "handle_mouse_exited")
	
func _process(delta):
	if Input.is_action_just_pressed("left_mouse_button") and mouse_over:
		emit_signal("on_drag_start", self)
	
func handle_mouse_entered():
	mouse_over = true
	
func handle_mouse_exited():
	mouse_over = false

func set_content(c):
	shape.add_child(c)
