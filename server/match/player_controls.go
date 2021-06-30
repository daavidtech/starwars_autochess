package match

// import "log"

// type BoughtUnit struct {
// 	UnitID     string
// 	UnitType   string
// 	Rank       int
// 	HP         int
// 	Mana       int
// 	AttackRate int
// }

// type PlayerEvent struct {
// 	BoughtUnit *BoughtUnit
// }

// type PlayerControls struct {
// 	currentMatch *Match
// 	player       *Player
// 	matchPool    *MatchPool
// }

// func (playerControl *PlayerControls) BuyUnit(index int) {
// 	if playerControl.currentMatch.phase != ShoppingPhase {
// 		return
// 	}

// 	shopUnit := playerControl.currentMatch.shop.Pick(index)

// 	playerControl.player.AddShopUnit(shopUnit)
// }

// func (playerControl *PlayerControls) SellUnit(unitID string) {
// 	if playerControl.currentMatch.phase != ShoppingPhase {
// 		log.Println("Can only sell units in shopping phase")

// 		return
// 	}

// 	playerControl.player.RemoveUnit(unitID)
// }

// func (playerControl *PlayerControls) PlaceUnit(unitID string, x int, y int) {
// 	unit := playerControl.player.GetUnit(unitID)

// 	if unit != nil {
// 		log.Println("Unit not found")

// 		return
// 	}

// 	unit.placement = &Placement{
// 		x: x,
// 		y: y,
// 	}
// }

// func (playerControls *PlayerControls) BuyLevelUp() {
// 	if playerControls.player.credits < 100 {
// 		log.Println("Insufficient funds")

// 		return
// 	}

// 	playerControls.player.AddXP(100)
// 	playerControls.player.UseCredits(100)
// }

// func (playerControls *PlayerControls) RecycleShopUnits() {
// 	if playerControls.currentMatch.phase != ShoppingPhase {
// 		log.Println("Cannot recycle shop if not in shopping phase")

// 		return
// 	}

// 	playerControls.currentMatch.shop.Fill(playerControls.player.GetLevel())
// }

// func (playerControls *PlayerControls) SeekMatch() {
// 	match := playerControls.matchPool.FindNewMatch()

// 	playerControls.currentMatch = match
// }

// func (playerControls *PlayerControls) SubscribeToPlayerEvents() <-chan PlayerEvent {
// 	return make(<-chan PlayerEvent)
// }
