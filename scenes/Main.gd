extends Spatial

var YourUnit = preload("res://scenes/your_unit.tscn")
var EnemyUnit = preload("res://scenes/enemy_unit.tscn")

var ServerConnection = preload("res://scenes/server_connection.gd")

signal unit_choosen

onready var placement_area = $StaticBody/PlacementArea
onready var dragging_area = $StaticBody2/DraggingArea
onready var unit_shop = $unit_shop
onready var unit_barrack = $unit_barrack
onready var your_health = $your_health
onready var your_level = $your_level
onready var your_money = $your_money
onready var lobby = $lobby
onready var countdown_label = $CountDownLabel
onready var countdown_timer = $CountdownTimer
onready var game_phase_label = $GamePhaseLabel
onready var connected_label = $ConnectedLabel
onready var main_menu = $main_menu
onready var login_form = $login_form

var placing_unit = null

var game_phase

var countdown_time: int

var units = {}

var width = 100
var heigth = 100

var game_state

var conn
	
func _ready():
	conn = ServerConnection.new()
	conn.connect("new_message", self, "_handle_msg")
	conn.connect("connected", self, "_on_connected")
	conn.connect("disconnected", self, "_on_disconnected")
	
	add_child(conn)
	
	countdown_label.visible = false
	lobby.visible = false
	your_money.visible = false
	your_health.visible = false
	your_level.visible = false
	unit_shop.visible = false
	unit_barrack.visible = false
	main_menu.visible = false

	lobby.conn = conn
	main_menu.conn = conn
	login_form.conn = conn	
	
#	game_state = GameState.new()
#	add_child(game_state)
#
#	print(placement_area.shape.extents.x)
#
#	game_state.connect("create_unit", self, "_handle_create_unit")
#	game_state.connect("unit_position_changed", self, "_handle_change_unit_position")
	
#	load_thing("unit_clone")
	
	unit_shop.connect("unit_bought", self, "_handle_unit_bought")

func _on_connected():
	connected_label.text = "Connected"
	
func _on_disconnected():
	connected_label.text = "Disconnected"

func _handle_msg(msg):	
	if msg.matchPhaseChanged != null:
		handle_game_phase_changed(msg.matchPhaseChanged.matchPhase)
	if msg.unitAdded != null:
		handle_unit_added(msg.unitAdded)
	if msg.unitRemoved != null:
		handle_unit_removed(msg.unitRemoved)
	if msg.unitSold != null:
		handle_unit_sold(msg.unitSold)
	if msg.unitUpgraded != null:
		handle_unit_upgraded(msg.unitUpgraded)
	if msg.unitPlaced != null:
		handle_unit_placed(msg.unitPlaced)
	if msg.startTimerTimeChanged != null:
		pass
	if msg.unitDied != null:
		handle_unit_died(msg.unitDied)
	if msg.unitTookDamage != null:
		pass
	if msg.unitUsedMana != null:
		pass
	if msg.unitUsedAbility != null:
		pass
	if msg.unitStartedMovingTo != null:
		handle_unit_started_moving(msg.unitStartedMovingTo)
	if msg.unitArrivedToPosition != null:
		handle_unit_arrived_to_position(msg.unitArrivedToPosition)
	if msg.unitStartedAttacking != null:
		handle_unit_started_attacking(msg["unitStartedAttacking"])
	if msg.unitStoppedAttacking != null:
		handle_unit_stopped_attacking(msg["unitStoppedAttacking"])
	if msg.launchParticle != null:
		pass
	if msg.playerMoneyChanged != null:
		handle_player_money_changed(msg["playerMoneyChanged"])
	if msg.playerLevelChanged != null:
		handle_player_level_changed(msg["playerLevelChanged"])
	if msg.playerHealthChanged != null:
		handle_player_health_changed(msg["playerHealthChanged"])
	if msg.shopRefilled != null:
		handle_shop_refilled(msg["shopRefilled"])
	if msg.shopUnitRemoved != null:
		handle_shop_unit_removed(msg["shopUnitRemoved"])
	if msg.countdownStarted != null:
		handle_countdown_started(msg.countdownStarted)
	if msg.roundCreated != null:
		handle_round_created(msg["roundCreated"])
	if msg.roundFinished != null:
		handle_round_finished(msg["roundFinished"])
	if msg.loginSuccess != null:
		handle_login_success(msg["loginSuccess"])
	if msg.currentMatch != null:
		handle_current_match(msg["currentMatch"])
	if msg.playerJoined != null:
		handle_player_joined(msg["playerJoined"])
	if msg.playerLeft != null:
		handle_player_left(msg["playerLeft"])
	if msg.battleUnitHealthChanged != null:
		handle_unit_health_changed(msg["battleUnitHealthChanged"])
	if msg.battleUnitManaChanged != null:
		handle_unit_mana_changed(msg["battleUnitManaChanged"])
	
func clear_units():
	for child in placement_area.get_children():
		placement_area.remove_child(child)
		
	units.clear()
	unit_barrack.clear()
	
func set_unit(loc: String, new_unit):
	var unit
	
	if units.has(new_unit.unitId):
		unit = units[new_unit.unitId]
	else:
		if !new_unit.has("team") || new_unit.team == 1:
			unit = YourUnit.instance()
			unit.connect("drag_started", self, "_on_drag_started")
			unit.connect("drag_finished", self, "_on_drag_finished")
	
		elif new_unit.team == 2:
			unit = EnemyUnit.instance()
	
		units[new_unit.unitId] = unit
	
	if !new_unit.has("team") ||new_unit.team == 1:		
		if unit.location != loc:
			unit.location = loc
			
			if loc == "barrack":
				unit_barrack.add_unit(unit)
		
			if loc == "battlefield":
				placement_area.add_child(unit)
			
			if loc == "placing":
				dragging_area.add_child(unit)
		
		unit.mana = new_unit.mana

	else:
		placement_area.add_child(unit)
		
	if new_unit.has("placement") and new_unit.placement != null:
		unit.translation = conv_server_coords(new_unit.placement.x, new_unit.placement.y)

	unit.unit_id = new_unit.unitId
	unit.unit_type = new_unit.unitType
	unit.hp = new_unit.hp
	
	unit.attack_rate = new_unit.attackRate
	unit.attack_rate = new_unit.attackRate
	unit.rank = new_unit.rank
	
	if new_unit.has("move_speed") and new_unit["move_speed"] != null:
		unit.move_speed = conv_move_speed(new_unit.moveSpeed)

func handle_unit_health_changed(health_changed):
	var unit = units[health_changed.unitId]
	
	unit.hp = health_changed.newHp

func handle_unit_mana_changed(mana_changed):
	var unit = units[mana_changed.unitId]
	
	unit.mana = mana_changed.newMana
	
func handle_player_joined(player_joined):
	lobby.add_player(player_joined.player.playerId, player_joined.player.name)
	
func handle_player_left(player_left):
	lobby.remove_player(player_left.player.playerId)
	
func handle_current_match(current_match):
	main_menu.visible = false
	
	handle_game_phase_changed(current_match.phase)
	
	for player in current_match.players:
		lobby.add_player(player.playerId, player.name)
	
func handle_login_success(login_success):
	login_form.visible = false
	main_menu.visible = true
	
func handle_round_finished(round_finished):
	clear_units()
	
	your_money.text = String(round_finished.newCreditsAmount)
	your_health.text = String(round_finished.newPlayerHealth)
	
	for unit in round_finished.units:
		if unit.placement == null:
			set_unit("barrack", unit)
		else:
			set_unit("battlefield", unit)

func handle_round_created(round_created):
	clear_units()
	
	for unit in round_created.units:
		set_unit("battlefield", unit)
	

func handle_countdown_started(countdown_started):
	countdown_label.visible = true
	
	countdown_time = countdown_started.startValue
	countdown_label.text = String(countdown_time)
	countdown_timer.wait_time = countdown_started.interval
	countdown_timer.start()	

func handle_shop_refilled(shop_refilled):
	unit_shop.fill(shop_refilled.shop_units)

func handle_shop_unit_removed(shop_unit_removed):
	unit_shop.remove_unit(shop_unit_removed.shopUnitId)

func handle_unit_took_damage(unit_took_damage):
	if units.has(unit_took_damage.unitId):
		var unit = units[unit_took_damage.unitId]
		
		unit.health -= unit_took_damage.amount

func handle_unit_started_moving(unit_started_moving):
	if units.has(unit_started_moving.unitId):
		var unit = units[unit_started_moving.unitId]
		
		unit.start_moving_to(conv_server_coords(unit_started_moving.x, unit_started_moving.y))
	
func handle_unit_arrived_to_position(unit_arrived):
	if units.has(unit_arrived.unitId):
		var unit = units[unit_arrived.unitId]
		
		unit.move_to_position(conv_server_coords(unit_arrived.x, unit_arrived.y))

func handle_unit_started_attacking(unit_started_attacking):
	if units.has(unit_started_attacking.unitId):
		var unit = units[unit_started_attacking.unitId]
		
		unit.attacking = true

func handle_unit_stopped_attacking(unit_stopped_attacking):
	if units.has(unit_stopped_attacking.unit):
		var unit = units[unit_stopped_attacking.unit]
		
		unit.attacking = false

func handle_player_money_changed(your_money_changed):
	your_money.text = String(your_money_changed.newMoney)

func handle_player_level_changed(player_level_changed):
	your_level.text = String(player_level_changed.newLevel)

func handle_player_health_changed(player_health_changed):
	your_health.text = String(player_health_changed.newHp)

func handle_game_phase_changed(match_phase):
	game_phase_label.text = match_phase
	game_phase = match_phase
	
	match match_phase:
		"InitPhase":			
			lobby.visible = false
			your_money.visible = false
			your_health.visible = false
			your_level.visible = false
			unit_shop.visible = false
			unit_barrack.visible = false
		"LobbyPhase":
			lobby.visible = true
			your_money.visible = false
			your_health.visible = false
			your_level.visible = false
			unit_shop.visible = false
			unit_barrack.visible = false
		"ShoppingPhase":
			lobby.visible = false
			your_money.visible = true
			your_health.visible = true
			your_level.visible = true
			unit_shop.visible = true
			unit_barrack.visible = true
		"PlacementPhase":
			lobby.visible = false
			your_money.visible = true
			your_health.visible = true
			your_level.visible = true
			unit_shop.visible = false
			unit_barrack.visible = true
		"BattlePhase":
			lobby.visible = false
			your_money.visible = true
			your_health.visible = true
			your_level.visible = true
			unit_shop.visible = false
			unit_barrack.visible = true
			
			if placing_unit != null:
				dragging_area.remove_child(placing_unit)
				unit_barrack.add_unit(placing_unit)
				
				placing_unit.location = "barrack"
				placing_unit.dragging = false
				placing_unit = null
		"EndPhase":
			lobby.visible = false
			your_money.visible = false
			your_health.visible = false
			your_level.visible = false
			unit_shop.visible = false
			unit_barrack.visible = false

func handle_unit_added(unit_added):
	print("Unit added " + unit_added.unitType)
	
	set_unit("barrack", unit_added)
	
func handle_unit_removed(unit_removed):
	var unit = units[unit_removed.unitId]
	
	if unit.location == "placing":
		dragging_area.remove_child(unit)
	
	if unit.location == "barrack":
		unit_barrack.remove_unit(unit)
		
	if unit.location == "battlefield":
		placement_area.remove_child(unit)
	
	units.erase(unit_removed.unitId)
	
func handle_unit_upgraded(unit_upgraded):
	var unit = units[unit_upgraded.unitId]
	
	set_unit(unit.location, unit_upgraded)

func handle_unit_placed(unit_placed):
	var unit = units[unit_placed.unitId]
	
	var new_translation = conv_server_coords(unit_placed.x, unit_placed.y)
	
	unit.translation = new_translation

func handle_unit_sold(unit_sold):
	if units.has(unit_sold.unitId):
		var unit = units[unit_sold.unitId]
		
		unit_barrack.remove_unit(unit)
	
func handle_unit_died(unit_died):
	if units.has(unit_died.unitId):
		var unit = units[unit_died.unitId]
		
		placement_area.remove_child(unit)

func _handle_unit_bought(index):
	print("unit bought ", index)
	
	conn.send_msg({
		"buyUnit": {
			"shopUnitIndex": index
		}
	})

func conv_server_coords(x: int, y: int) -> Vector3:
	var size = placement_area.shape.extents
	
	var start_x = -size.x
	var start_y = -size.z
	
	var z_ratio = (size.z * 2) / width
	var x_ratio = (size.x * 2) / heigth
	
	var z_conv = x * z_ratio - size.z
	var x_conv = y * x_ratio - size.x
	
	return Vector3(x_conv, placement_area.translation.y, z_conv)

func trans_to_server_coord(z, x):
	var size = placement_area.shape.extents
	
	var z_ratio = width / (size.z * 2)
	var x_ratio = heigth / (size.x * 2)
	
	return {
		"x": round(abs(size.z + z) * z_ratio),
		"y": round(abs(size.x + x) * x_ratio)
	}

func conv_move_speed(move_speed: int):
	var size = placement_area.shape.extents
	
	return size.x / 100 * move_speed

func _unhandled_input(event):
	if placing_unit != null and event is InputEventMouseMotion:
		var viewport = get_viewport()
		var camera = viewport.get_camera()
		var mouse_position = viewport.get_mouse_position()
		var from = camera.project_ray_origin(mouse_position)
		var to = from + camera.project_ray_normal(mouse_position) * 200
		var direct_state = get_world().direct_space_state
		var collision = direct_state.intersect_ray(from, to, [], 32)
		
		
#		print("Collision ", collision)
		if collision:			
			placing_unit.translation.x = collision.position.x
			placing_unit.translation.z = collision.position.z
			placing_unit.translation.y = collision.position.y


func _on_drag_started(unit):
	print("On drag started")
	
	if game_phase == "PlacementPhase":
		if unit.location == "barrack":
			unit_barrack.remove_unit(unit)
		
		if unit.location == "battlefield":
			placement_area.remove_child(unit)
			
		dragging_area.add_child(unit)
		unit.dragging = true
		unit.location = "placing"
		placing_unit = unit
	
func _on_drag_finished(unit):
	print("On drag finisehed")
	
	var viewport = get_viewport()
	var camera = viewport.get_camera()
	var mouse_position = viewport.get_mouse_position()
	var from = camera.project_ray_origin(mouse_position)
	var to = from + camera.project_ray_normal(mouse_position) * 200
	var direct_state = get_world().direct_space_state
	var collision = direct_state.intersect_ray(from, to, [], 16)
	
	if collision:
		dragging_area.remove_child(placing_unit)
		placement_area.add_child(placing_unit)
		
		placing_unit.translation.x = collision.position.x
		placing_unit.translation.z = collision.position.z
		
		unit.location = "battlefield"
		
		var server_coord = trans_to_server_coord(collision.position.z, collision.position.x)
		
		conn.send_msg({
			"placeUnit": {
				"unitId": unit.unit_id,
				"x": server_coord.x,
				"y": server_coord.y
			}
		})
	else:
		unit.location = "barrack"		
		
		dragging_area.remove_child(placing_unit)
		unit_barrack.add_unit(placing_unit)
		
	unit.dragging = false
	placing_unit = null

	
#func _handle_create_unit(id, unit_type, x, y):
#	print("handle create unit")
#
#	var new_unit = Unit.instance()
#
#	move_to_point(new_unit, x, y)
#
#	placement_area.add_child(new_unit)
#	unit_map[id] = new_unit

#func _handle_change_unit_position(id, x: int, y: int):
#	var unit = unit_map.get(id)
#
#	move_to_point(unit, x, y)


func _on_CountdownTimer_timeout():
	if countdown_time < 1:
		countdown_timer.stop()
		countdown_label.visible = false
	
	countdown_time -= 1
	
	countdown_label.text = String(countdown_time)
	
	if countdown_time < 1:
		countdown_timer.start(0.2)
