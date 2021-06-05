extends Control

signal start_game_clicked;

func _on_Button_pressed():
	emit_signal("start_game_clicked")
