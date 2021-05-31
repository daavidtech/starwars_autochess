extends Spatial

var current_content = null

func take_contents():
	var content = current_content
	current_content = null

	return content
