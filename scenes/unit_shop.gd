extends StaticBody

var ShopUnitOutline = preload("res://scenes/shop_unit_outline.tscn")

onready var shape = $Shape

func _ready():
	pass # Replace with function body.
	
func fill(units):
	var size = shape.shape.extents
	
	var start_x = -size.x
	var start_y = -size.z
	
	var one_width = (size.x * 2) / units.size()
	var one_half_width = one_width / 2
	
	for unit in units:
		print("adding to shop ", unit.unit_type)
		
		var path = "res://assets/" + unit.unit_type + "/" + unit.unit_type + ".glb"
		
		var outline = ShopUnitOutline.instance()
		
		outline.translation.z = start_x + one_half_width
		
		start_x += one_width
		
		shape.add_child(outline)
		
		var model = ResourceLoader.load(path).instance()
		
		outline.set_content(model)

