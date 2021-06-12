extends Node

var ServerConnection = preload("res://scenes/server_connection.gd")

signal unit_position_changed(id, x, y)
signal game_phase_changed
signal remove_unit(id)
signal create_unit(id, unit_type, x, y)

var width = 100
var height = 100

var other_players = []

enum Status {
	IDLE,
	LOBBY,
	SHOPPING_PHASE,
	PLACEMENT_PHASE,
	BATTLE_PHASE,
	WIN,
	LOSE
}

var status = Status.LOBBY

var conn

func _ready():
	conn = ServerConnection.new()
	add_child(conn)
	
	conn.connect("connected", self, "_handle_connected")
	conn.connect("disconnected", self, "_handle_disconnected")
	conn.connect("new_message", self, "_handle_msg")
	
func _handle_msg(msg):
	if msg.createUnit != null:
		var create_unit = msg.createUnit
		
		emit_signal("create_unit", create_unit.id, create_unit.unitType, create_unit.x, create_unit.y)
	
	if msg.changeUnitPosition != null:
		var change_position = msg.changeUnitPosition
		
		emit_signal("unit_position_changed", change_position.id, change_position.x, change_position.y)

func _handle_connected():
	print("Connected")
	pass

func _handle_disconnected():
	print("Disconnected")
	pass

func buy_unit(id: String):
	pass
	
func sell_unit(id: String):
	pass
	
func buy_level_up():
	pass
	
func i_am_ready():
	pass
	
func change_unit_placement(x: int, y: int):
	pass
