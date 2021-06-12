package networking

type CreateUnit struct {
	ID       string `json:"id"`
	UnitType string `json:"unitType"`
	Enemy    bool   `json:"enemy"`
	MaxHP    int    `json:"maxHp"`
	CurrHP   int    `json:"currHp"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
}

type RemoveUnit struct {
	ID string `json:"id"`
}

type ChangeUnitPosition struct {
	ID string `json:"id"`
	X  int    `json:"x"`
	Y  int    `json:"y"`
}

type MessageToClient struct {
	CreateUnit         *CreateUnit         `json:"createUnit"`
	RemoveUnit         *RemoveUnit         `json:"removeUnit"`
	ChangeUnitPosition *ChangeUnitPosition `json:"changeUnitPosition"`
}
