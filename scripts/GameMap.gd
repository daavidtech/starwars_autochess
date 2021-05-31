extends StaticBody

onready var mesh = $Mesh
onready var shape = $Shape

func get_size():
	return {
		"width": mesh.mesh.size.x,
		"height": mesh.mesh.size.y
	}
