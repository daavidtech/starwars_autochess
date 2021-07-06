extends Spatial

var YourUnit = preload("res://scenes/your_unit.tscn")

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

var placing_unit = null

var game_phase

var countdown_time: int

enum LocationType {
	Barrack,
	BattleField,
	Placing
}

var units = {}
var placement_units = {}

var width = 100
var heigth = 100

var game_state

var conn

func load_thing(type: String):
#	var path = "res://assets/exported/"
#
#	path += type + "/" + type + ".gltf"
	
	var path = "res://scenes/your_unit.tscn"
	
	print("Loading resource from ", path)
	
	var does_exists = ResourceLoader.has(path)
	
	print("does_exists " + String(does_exists))
	
	var Droid = ResourceLoader.load(path)
	
	var i = Droid.instance()
	
	i.translation.y = placement_area.translation.y
	
	add_child(i)
	
	i.move_speed = 5;
	i.start_move(conv_coords(50, 50))
	

func move_to_point(u, x: int, y: int):
	print("move_to ", x, " y ", y)
	
	var size = placement_area.shape.extents
	
	var start_x = -size.x
	var start_y = -size.z
	
	print("start_x ", start_x)
	print("start_y ", start_y)
	
	var x_ratio = size.z / game_state.width
	var y_ratio = size.x / game_state.height
	
	var translation_z = start_x + x_ratio * x
	var translation_x = start_y + y_ratio * y
	
	print("translation_x ", translation_x)
	print("translation_z ", translation_z)
	
	u.translation.x = translation_x
	u.translation.z = translation_z
	u.translation.y = placement_area.translation.y

func _ready():
	conn = ServerConnection.new()
	conn.connect("new_message", self, "_handle_msg")
	
	add_child(conn)
	
	countdown_label.visible = false
	lobby.visible = false
	your_money.visible = false
	your_health.visible = false
	your_level.visible = false
	unit_shop.visible = false
	unit_barrack.visible = false
	
#	game_state = GameState.new()
#	add_child(game_state)
#
#	print(placement_area.shape.extents.x)
#
#	game_state.connect("create_unit", self, "_handle_create_unit")
#	game_state.connect("unit_position_changed", self, "_handle_change_unit_position")
	
#	load_thing("unit_clone")
	
	unit_shop.connect("unit_bought", self, "_handle_unit_bought")

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
	if msg.startTimerTimeChanged != null:
		pass
	if msg.unitDied != null:
		handle_unit_died(msg.unitDied)
	if msg.unitPlaced != null:
		handle_unit_placed(msg.unitPlaced)
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
		
		unit.start_moving(conv_coords(unit_started_moving.x, unit_started_moving.y))
	
func handle_unit_arrived_to_position(unit_arrived):
	if units.has(unit_arrived.unitId):
		var unit = units[unit_arrived.unitId]
		
		unit.move_to_position(conv_coords(unit_arrived.x, unit_arrived.y))

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

func handle_game_phase_changed(game_phase_changed):
	game_phase_label.text = game_phase_changed
	game_phase = game_phase_changed
	
	match game_phase_changed:
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

func handle_unit_added(unit_added):
	print("Unit added " + unit_added.unitType)
		
	var new_unit = YourUnit.instance()
	unit_barrack.add_unit(new_unit)
	
	new_unit.unit_type = unit_added.unitType
	new_unit.hp = unit_added.hp
	new_unit.mana = unit_added.mana
	new_unit.attack_rate = unit_added.attackRate
	new_unit.rank = unit_added.rank
	new_unit.location = "barrack"
	
	units[unit_added.unitId] = new_unit
	
	new_unit.connect("drag_started", self, "_on_drag_started")
	new_unit.connect("drag_finished", self, "_on_drag_finished")
	
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
	
	unit.hp = unit_upgraded.hp
	unit.mana = unit_upgraded.mana
	unit.attack_rate = unit_upgraded.attackRate
	unit.rank = unit_upgraded.rank
	
func handle_unit_sold(unit_sold):
	if units.has(unit_sold.unitId):
		var unit = units[unit_sold.unitId]
		
		unit_barrack.remove_unit(unit)
	
func handle_unit_died(unit_died):
	if units.has(unit_died.unitId):
		var unit = units[unit_died.unitId]
		
		placement_area.remove_child(unit)
	
func handle_unit_placed(unit_placed):
	pass


func _handle_unit_bought(index):
	print("unit bought ", index)
	
	conn.send_msg({
		"buyUnit": {
			"shopUnitIndex": index
		}
	})

func conv_coords(x: int, y: int) -> Vector3:
	var size = placement_area.shape.extents
	
	var start_x = -size.x
	var start_y = -size.z
	
	var x_ratio = size.z / width
	var y_ratio = size.x / heigth
	
	var convz = start_x + x_ratio * x
	var convx = start_y + y_ratio * y
	
	return Vector3(convx, placement_area.translation.y, convz)

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


func _on_drag_started(unit):
	print("On drag started")
	
	if game_phase == "PlacementPhase":
		if unit.location == "barrack":
			unit_barrack.remove_unit(unit)
		
		if unit.location == "battlefield":
			placement_area.remove_unit(unit)
			
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
	else:
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
