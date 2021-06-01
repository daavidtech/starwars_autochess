extends Spatial

onready var progress_bar = $Viewport/Control/ProgressBar

func set_value(v):
	progress_bar.value = v
