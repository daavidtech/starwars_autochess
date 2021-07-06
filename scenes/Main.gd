extends Spatial

var GameCoordinator = preload("res://scripts/GameCoordinator.gd")
var Droid = preload("res://scenes/droid/Droid.tscn")

onready var picking_options = $PickingOptions
onready var unit_barrack = $UnitBarrack
onready var picking_ui = $PickingUi

onready var setup_timer = $SetupTimer
onready var new_round_timer = $NewRoundTimer

onready var setup_timer_text = $SetupTimerText
onready var gold_remaining_text = $CurrentCold

onready var game_end_text = $GameEndText

var game_coordinator = null

export var gold = 40

var health = 100

var start_ticks = 10

var setup_stage = false

var dragging_unit = null

var allow_dragging = false

var units = []
var battle_units = []

func _ready():
	picking_options.connect("option_clicked", self, "handle_option_clicked")
	picking_ui.connect("start_game_clicked", self, "on_game_start")
	unit_barrack.connect("on_drag_started", self, "on_drag_started")
	
	game_coordinator = GameCoordinator.new()
	
	game_coordinator.connect("someone_won", self, "on_someone_won")
	
	gold_remaining_text.text = String(gold)
	
func on_someone_won(team):
	print("Someone won")
	
	if game_end_text:
		game_end_text.visible = true
		if team == 1:	
			game_end_text.text = "You win"
		else:
			game_end_text.text = "Enemy won"
			
	handle_round_end()
	
func _physics_process(delta):
	if dragging_unit != null:
		if Input.is_action_just_released("left_mouse_button"):
			dragging_unit = null
		else:
			move_to_mouse_position(dragging_unit)


func move_to_mouse_position(unit):
	var viewport = get_viewport()
	var camera = viewport.get_camera()
	var mouse_position = viewport.get_mouse_position()
	var from = camera.project_ray_origin(mouse_position)
	var to = from + camera.project_ray_normal(mouse_position) * 200
	var direct_state = get_world().direct_space_state
	var collision = direct_state.intersect_ray(from, to)
	
	if collision:			
		unit.translation.x = collision.position.x
		unit.translation.z = collision.position.z
	
func handle_option_clicked(opt):
	if unit_barrack.is_full():
		print("Barrack is full")
		
		return 
	
	if opt.cost > gold:
		print("Not enough cold left")
		
		return
		
	print("unit picked ", opt.unit_type, " level ", opt.level)
	
	picking_options.remove_option(opt)
	
	unit_barrack.apply_option(opt)
	
	set_gold(gold - opt.cost)

	
func start_game():
	allow_dragging = false
	
	for unit in units:
		var droid = Droid.instance()

		var child = get_child(get_children().find(unit))
		
		droid.translation.x = child.translation.x
		droid.translation.y = child.translation.y
		droid.translation.z = child.translation.z
		
		droid.team = 1
		droid.game_coordinator = game_coordinator
		game_coordinator.add_unit(droid)
		
		droid.connect("death", self, "handle_death")
		
		remove_child(child)
		add_child(droid)
		
		battle_units.push_back(droid)
		
		droid.attack()
		
		
	spawn_enemy_unit(0, 3)
	spawn_enemy_unit(-1, 3)
	spawn_enemy_unit(1, 3)
	spawn_enemy_unit(2, 3)
	spawn_enemy_unit(3, 3)

func spawn_enemy_unit(x, y):
	var enemy_droid = Droid.instance()
	
	enemy_droid.game_coordinator = game_coordinator
	enemy_droid.set_position(x, y)
	enemy_droid.team = 2
	game_coordinator.add_unit(enemy_droid)
	
	enemy_droid.connect("death", self, "handle_death")
	
	battle_units.push_back(enemy_droid)
	
	add_child(enemy_droid)
	
	enemy_droid.attack()

func handle_death(unit):
	game_coordinator.remove_unit(unit)
	remove_child(unit)

func on_game_start():
	picking_ui.visible = false
	picking_options.visible = false
	setup_stage = true
	allow_dragging = true
	setup_timer.start()

func set_gold(g: int):
	gold = g 
	gold_remaining_text.text = String(g)

func handle_round_end():
	set_gold(gold + 10)
	units.clear()
	
	new_round_timer.start()
	
func transistion_to_picking_phase():
	print("Transistioning...")
	
	picking_ui.visible = true
	picking_options.visible = true
	picking_options.reset()
	setup_timer_text.visible = true
	game_end_text.visible = false
	start_ticks = 10
	
	
	for unit in battle_units:
		remove_child(unit)
		
	battle_units.clear()
	
	unit_barrack.show_all()

func on_drag_started(unit):
	if allow_dragging:
		print("Drag started")
		var cloned_unit = unit.duplicate()
		self.add_child(cloned_unit)
		unit_barrack.hide_unit(unit)
		units.push_back(cloned_unit)
		dragging_unit = cloned_unit
		move_to_mouse_position(cloned_unit)

func _on_SetupTimer_timeout():
	start_ticks -= 1
	
	setup_timer_text.text = String(start_ticks)
	
	if start_ticks == 0:
		setup_timer_text.visible = false
		setup_timer.stop()
		
		start_game()

func _on_NewRoundTimer_timeout():
	transistion_to_picking_phase()
