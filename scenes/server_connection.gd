extends Node

signal new_message(msg)
signal connected
signal disconnected

# The URL we will connect to
export var websocket_url = "ws://localhost:4100/api/socket"

var retry_timer

# Our WebSocketClient instance
var _client = WebSocketClient.new()

func send_msg(msg):
	var json_msg = JSON.print(msg)
	
	_client.get_peer(1).put_packet(json_msg.to_utf8())

func _ready():
	retry_timer = Timer.new()
	add_child(retry_timer)
	retry_timer.wait_time = 2
	retry_timer.connect("timeout", self, "_retry_now")	

	retry_timer.start()

func _retry_now():
	print("Retry timer triggered")
	
	_client.disconnect_from_host()
	
	try_to_connect()

func try_to_connect():
	print("try_to_connect")
	
	_client = WebSocketClient.new()
	
	_client.connect("connection_closed", self, "_closed")
	_client.connect("connection_error", self, "_closed")
	_client.connect("connection_established", self, "_connected")
	_client.connect("data_received", self, "_on_data")
	
	var err = _client.connect_to_url(websocket_url)
	if err != OK:
		print("Unable to connect")

func _closed(was_clean = false):
	print("Connection closed")
	
	print("Closed, clean: ", was_clean)
	
	emit_signal("disconnected")
	
	retry_timer.start()

func _connected(proto = ""):
	# This is called on connection, "proto" will be the selected WebSocket
	# sub-protocol (which is optional)
	print("Connected with protocol: ", proto)
	
	emit_signal("connected")
	
	retry_timer.stop()

func _on_data():
	var msg = _client.get_peer(1).get_packet().get_string_from_utf8()

	
	var json_msg = JSON.parse(msg)
	
	var result = json_msg.result
	
	emit_signal("new_message", result)

func _process(delta):
	# Call this in _process or _physics_process. Data transfer, and signals
	# emission will only happen when calling this function.
	_client.poll()
