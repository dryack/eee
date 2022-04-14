package main

type Empire struct {
	Id        uint64    `json:"id"`
	Name      string    `json:"name"`
	Prestige  int64     `json:"prestige"`
	Interest  float64   `json:"interest"`
	Policies  Policies  `json:"policies"`
	Resources Resources `json:"resources"`
}

type Policies struct {
	Taxation       uint64                 `json:"taxation"`
	Priorities     ConstructionPriorities `json:"priorities"`
	BalancedBudget bool                   `json:"balancedBudget"` // prevent elective debt spending - only debts that are necessary to maintain the military and food supply are permitted until this is changed
}

type Resources struct {
	MegaCredits uint64   `json:"megaCredits"`
	Food        uint64   `json:"food"`
	Population  uint64   `json:"population"`
	Systems     []System `json:"system"`
}

// NewEmpire is called after a player has entered their player information and doesn't quit out of the game
func (player *player) NewEmpire() {

}
