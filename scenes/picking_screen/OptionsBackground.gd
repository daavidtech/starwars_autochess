extends Control


func _ready():
	pass # Replace with function body.

func _process(delta):
	var texture = $Viewport.get_texture()
	$Sprite.texture = texture
	
