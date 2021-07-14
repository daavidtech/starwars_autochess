extends Control

var conn

func _on_Button_pressed():	
	print("Find game")
	
	conn.send_msg({
		"findMatch": {}
	})
