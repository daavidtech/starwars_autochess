extends RigidBody


# Declare member variables here. Examples:
# var a = 2
# var b = "text"

var mouse_is_over = false
var drag_button_already_down = false
var dragging = false

# Called when the node enters the scene tree for the first time.
func _ready():
	pass # Replace with function body.

func _physics_process(delta):
	if drag_button_already_down == false and mouse_is_over and Input.is_action_pressed("left_mouse_button"):
		if dragging == false:
			dragging = true
			print("start dragging")
		
	if Input.is_action_pressed("left_mouse_button") == false and dragging == true:
		dragging = false
		drag_button_already_down = false
		print("dragging stopped")


# Called every frame. 'delta' is the elapsed time since the previous frame.
#func _process(delta):
#	pass


func _on_Droid_mouse_entered():
	mouse_is_over = true
	if Input.is_action_pressed("left_mouse_button"):
		drag_button_already_down = true
	print("tickling")


func _on_Droid_mouse_exited():
	mouse_is_over = false
	drag_button_already_down = false
	print("stop tickling")
