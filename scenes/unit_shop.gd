extends StaticBody

signal unit_bought(item)

var ShopUnitOutline = preload("res://scenes/shop_unit_outline.tscn")

onready var shape = $Shape

func _ready():
	pass # Replace with function body.
	
func fill(units):
	if units.size() == 0:
		return
	
	var size = shape.shape.extents
	
	var start_x = -size.x
	var start_y = -size.z
	
	var one_width = (size.x * 2) / units.size()
	var one_half_width = one_width / 2
	
	var index = 0
	
	for unit in units:
		print("adding to shop ", unit.unit_type)
		
		var path = "res://assets/" + unit.unit_type + "/" + unit.unit_type + ".glb"
		
		var outline = ShopUnitOutline.instance()
		shape.add_child(outline)
		
		outline.id = unit.id
		outline.level = unit.level
		outline.rank = unit.rank
		outline.cost = unit.cost
		
		outline.translation.z = start_x + one_half_width
		outline.connect("unit_choosen", self, "handle_unit_choosen")
		
		start_x += one_width
		
		var model = ResourceLoader.load(path).instance()
		
		outline.set_content(model)

func remove_unit(id: int):
	for child in shape.get_children():
		if child.id == id:
			shape.remove_child(child)
			break

func handle_unit_choosen(index):
	emit_signal("unit_bought", index)
