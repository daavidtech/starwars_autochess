extends Node2D

onready var cost_label = $CostLabel

func set_cost(cost: int):
	cost_label.text = String(cost)
