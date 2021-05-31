extends StaticBody

onready var shape = $Shape

var CharacterDragTarget = preload("res://scenes/CharacterDragTarget.tscn")

export var horizontal_rows = 10 setget set_horizontal_rows
export var vertical_rows = 10 setget set_vertical_rows
export var show_location = false

func set_horizontal_rows(v):
	horizontal_rows = v
	
func set_vertical_rows(v):
	vertical_rows = v

func _ready():
	print("character_position_grid transform", scale)
	#pass # Replace with function body.
	
	draw(horizontal_rows, vertical_rows)

# Draws character grid rectangles
func draw(horizontal_rows, vertical_rows):
	print("scale x ", shape.scale.x, " scale z ", shape.scale.z)
	
	var vertical_diameter = shape.scale.x / 2
	var horizontal_diameter = shape.scale.z / 2
	
	var vertical_start_position = shape.transform.origin.x + vertical_diameter
	var horizontal_start_position = shape.transform.origin.z - horizontal_diameter
	
#	print("vertical_start_position: ", vertical_start_position)
#	print("horizontal_start_positon: ", horizontal_start_position)
	
	var item_height = shape.scale.x / vertical_rows
	var item_width = shape.scale.z / horizontal_rows
	
	var item_height_diameter = item_height / 2
	var item_width_diameter = item_width / 2
	
#	print("one item height:", item_height, " and width:", item_width)
#	print("item_height_diameter ", item_height_diameter)
#	print("item_width_diameter ", item_width_diameter)
	
	for i in range(vertical_rows):
		for j in range(horizontal_rows):
#			print("item_width_diameter * j", item_width * j)
#			print("item_height_diameter * i", item_height * i)
#
			var z = horizontal_start_position + item_width * j + item_width_diameter
			var x = vertical_start_position - item_height * i - item_height_diameter
			
#			print("Draw object", " j:", j, "i:", i, " z:", z, " x:", x)
			
			var new_target = CharacterDragTarget.instance()
			new_target.transform.origin.x = x
			new_target.transform.origin.z = z
			new_target.transform.origin.y = shape.transform.origin.y
			new_target.scale.x = item_height_diameter - shape.scale.x / 100
			new_target.scale.y = 0.01
			new_target.scale.z = item_width_diameter - shape.scale.y / 100
			
			add_child(new_target)
			
