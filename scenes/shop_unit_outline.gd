extends StaticBody

onready var shape = $shape

func set_content(c):
	shape.add_child(c)
