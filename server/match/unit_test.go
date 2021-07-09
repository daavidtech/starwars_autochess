package match

import "testing"

func Test_is_not_in_attack_range(t *testing.T) {
	unit := BattleUnit{
		X: 10,
		Y: 10,
	}

	success := unit.isInsideAttackRange(10, 80)

	if success {
		t.Error("Unit should not be in attack range")
	}
}

func Test_is_in_attack_range(t *testing.T) {
	unit := BattleUnit{
		X:           10,
		Y:           10,
		AttackRange: 80,
	}

	success := unit.isInsideAttackRange(10, 80)

	if !success {
		t.Error("Unit should be in attack range")
	}
}

func Test_calculate_next_loc_returns_nil_if_inside_attack_range(t *testing.T) {
	unit := BattleUnit{
		X:           10,
		Y:           10,
		AttackRange: 10,
		currAttackTarget: &BattleUnit{
			X: 20,
			Y: 15,
		},
	}

	loc := unit.calcNextLoc()

	if loc != nil {
		t.Error("Next loc should be nil")
	}
}

func Test_next_loc_is_in_x_axis(t *testing.T) {
	unit := BattleUnit{
		X:          10,
		Y:          10,
		AttackRate: 10,
		currAttackTarget: &BattleUnit{
			X: 100,
			Y: 50,
		},
	}

	loc := unit.calcNextLoc()

	if loc == nil {
		t.Error("Next loc cannot be nil")
	}

	if loc.X != 100 || loc.Y != 10 {
		t.Errorf("Next loc is invalid %v %v", loc.X, loc.Y)
	}
}

func Test_next_loc_is_in_y_axis(t *testing.T) {
	unit := BattleUnit{
		X:          10,
		Y:          10,
		AttackRate: 10,
		currAttackTarget: &BattleUnit{
			X: 50,
			Y: 100,
		},
	}

	loc := unit.calcNextLoc()

	if loc == nil {
		t.Error("Next loc cannot be nil")
	}

	if loc.X != 10 || loc.Y != 100 {
		t.Errorf("Next loc is invalid %v %v", loc.X, loc.Y)
	}
}

func Test_next_loc_x_is_smaller(t *testing.T) {
	unit := BattleUnit{
		X:          100,
		Y:          100,
		AttackRate: 10,
		currAttackTarget: &BattleUnit{
			X: 10,
			Y: 50,
		},
	}

	loc := unit.calcNextLoc()

	if loc == nil {
		t.Error("Next loc cannot be nil")
	}

	if loc.X != 10 || loc.Y != 100 {
		t.Errorf("Next loc is invalid %v %v", loc.X, loc.Y)
	}
}

func Test_next_loc_y_is_smaller(t *testing.T) {
	unit := BattleUnit{
		X:          100,
		Y:          100,
		AttackRate: 10,
		currAttackTarget: &BattleUnit{
			X: 50,
			Y: 10,
		},
	}

	loc := unit.calcNextLoc()

	if loc == nil {
		t.Error("Next loc cannot be nil")
	}

	if loc.X != 100 || loc.Y != 10 {
		t.Errorf("Next loc is invalid %v %v", loc.X, loc.Y)
	}
}

func Test_move_towards_in_x_axis(t *testing.T) {
	unit := BattleUnit{
		X:         5,
		Y:         5,
		MoveSpeed: 10,
		nextLoc: &Point{
			X: 40,
			Y: 10,
		},
	}

	unit.moveTowardsNextLoc(1)

	if unit.X != 15 && unit.Y != 5 {
		t.Error("Unit moved to wrong location")
	}
}

func Test_move_towards_in_y_axis(t *testing.T) {
	unit := BattleUnit{
		X:         5,
		Y:         5,
		MoveSpeed: 10,
		nextLoc: &Point{
			X: 10,
			Y: 40,
		},
	}

	unit.moveTowardsNextLoc(1)

	if unit.X != 5 && unit.Y != 15 {
		t.Error("Unit moved to wrong location")
	}
}

func Test_move_towards_in_x_axis_negative_direction(t *testing.T) {
	unit := BattleUnit{
		X:         40,
		Y:         20,
		MoveSpeed: 10,
		nextLoc: &Point{
			X: 10,
			Y: 10,
		},
	}

	unit.moveTowardsNextLoc(1)

	if unit.X != 30 || unit.Y != 20 {
		t.Error("Unit moved to wrong location")
	}
}

func Test_move_towards_in_y_axis_negative_direction(t *testing.T) {
	unit := BattleUnit{
		X:         20,
		Y:         40,
		MoveSpeed: 10,
		nextLoc: &Point{
			X: 10,
			Y: 10,
		},
	}

	unit.moveTowardsNextLoc(1)

	if unit.X != 20 || unit.Y != 30 {
		t.Error("Unit moved to wrong location")
	}
}

func Test_dont_move_past_target_location_on_x_axis(t *testing.T) {
	unit := BattleUnit{
		X:         20,
		Y:         20,
		MoveSpeed: 15,
		nextLoc: &Point{
			X: 30,
			Y: 25,
		},
	}

	unit.moveTowardsNextLoc(1)

	if unit.X != 30 || unit.Y != 20 {
		t.Errorf("Unit moved to wrong location %v %v", unit.X, unit.Y)
	}
}

func Test_dont_move_past_target_location_on_y_axis(t *testing.T) {
	unit := BattleUnit{
		X:         20,
		Y:         20,
		MoveSpeed: 15,
		nextLoc: &Point{
			X: 25,
			Y: 30,
		},
	}

	unit.moveTowardsNextLoc(1)

	if unit.X != 20 || unit.Y != 30 {
		t.Errorf("Unit moved to wrong location %v %v", unit.X, unit.Y)
	}
}

func Test_dont_move_past_target_location_on_x_axis_negative(t *testing.T) {
	unit := BattleUnit{
		X:         30,
		Y:         30,
		MoveSpeed: 15,
		nextLoc: &Point{
			X: 20,
			Y: 25,
		},
	}

	unit.moveTowardsNextLoc(1)

	if unit.X != 20 || unit.Y != 30 {
		t.Errorf("Unit moved to wrong location %v %v", unit.X, unit.Y)
	}
}

func Test_dont_move_past_target_location_on_y_axis_negative(t *testing.T) {
	unit := BattleUnit{
		X:         30,
		Y:         30,
		MoveSpeed: 15,
		nextLoc: &Point{
			X: 25,
			Y: 20,
		},
	}

	unit.moveTowardsNextLoc(1)

	if unit.X != 30 || unit.Y != 20 {
		t.Errorf("Unit moved to wrong location %v %v", unit.X, unit.Y)
	}
}
