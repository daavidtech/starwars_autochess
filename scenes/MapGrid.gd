extends CollisionShape

var GridColumn = preload("res://scenes/GridColumn.tscn")

export var rows = 10
export var colums = 10

var grid_positions = []

func add_content(x, y, content):
	pass

# Called when the node enters the scene tree for the first time.
func _ready():
	draw()

func draw():
	var vertical_diameter = shape.scale.x / 2
	var horizontal_diameter = shape.scale.z / 2
		var vertical_start_position = shape.transform.origin.x + vertical_diameter
	var horizontal_start_position = shape.transform.origin.z - horizontal_diameter
	
	var item_height = shape.scale.x / rows
	var item_width = shape.scale.z / columns
	
	var item_height_diameter = item_height / 2
	var item_width_diameter = item_width / 2
	
	for i in range(rows):
		for j in range(columns):
			var z = horizontal_start_position + item_width * j + item_width_diameter
			var x = vertical_start_position - item_height * i - item_height_diameter

			
			var new_target = GridColumn.instance()
			new_target.transform.origin.x = x
			new_target.transform.origin.z = z
			new_target.transform.origin.y = shape.transform.origin.y
			new_target.scale.x = item_height_diameter - shape.scale.x / 100
			new_target.scale.y = 0.01
			new_target.scale.z = item_width_diameter - shape.scale.y / 100
			
			add_child(new_target)
