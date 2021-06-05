extends Spatial

signal option_clicked(option);

onready var option1 = $Option1
onready var option2 = $Option2
onready var option3 = $Option3
onready var option4 = $Option4
onready var option5 = $Option5

func _ready():
	option1.connect("clicked", self, "handle_click")
	option2.connect("clicked", self, "handle_click")
	option3.connect("clicked", self, "handle_click")
	option4.connect("clicked", self, "handle_click")
	option5.connect("clicked", self, "handle_click")
	
func handle_click(target):
	emit_signal("option_clicked", target)

func remove_option(opt):
	if option1 == opt:
		option1.visible = false
		
	if option2 == opt:
		option2.visible = false
	
	if option3 == opt:
		option3.visible = false
		
	if option4 == opt:
		option4.visible = false
		
	if option5 == opt:
		option5.visible = false

func reset():
	option1.visible = true
	option2.visible = true
	option3.visible = true
	option4.visible = true
	option5.visible = true
