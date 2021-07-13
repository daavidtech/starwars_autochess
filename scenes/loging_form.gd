extends Control

var ServerConn = preload("res://scenes/server_connection.gd")

onready var username_text_box = $Username
onready var password_text_box = $Password

var conn

func _ready():
	conn = ServerConn.new()

func _on_Button_pressed():
	print("Login " + username_text_box.text + " " + password_text_box.text)
	
	conn.send_msg({
		"login": {
			"username": username_text_box.text,
			"password": password_text_box.text
		}
	})
