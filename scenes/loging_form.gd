extends Control

var ServerConn = preload("res://scenes/server_connection.gd")

onready var username = $username_textbox
onready var password = $password_line_edit

var conn

func _on_Button_pressed():
	print("Login " + username.text + " " + password.text)
	
	conn.send_msg({
		"login": {
			"username": username.text,
			"password": password.text
		}
	})
