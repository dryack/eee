package main

import "math/rand"

type System struct {
	Id      uint64  `json:"id"`
	Size    uint64  `json:"size"`
	Regions Regions `json:"regions"`
	Next    uint64  `json:"next"`
}

type Regions struct {
	// DOCS: Solar regions generate energy that is converted to credits.  They are a stable source of income, but are
	//  easily destroyed during planetary bombardments
	Solar uint64 `json:"solar"`
	Wind  uint64 `json:"wind"`
	// DOCS: Mining regions are the most stable source of income, but are the least valuable.  During times of unrest
	//  they are the bedrock of an empire's economy
	Mining uint64 `json:"mining"`
	// DOCS:  Tourism regions are a massive source of income, but their value drops precipitously during times of unrest,
	//  and they drop to zero during war.
	Tourism uint64 `json:"tourism"`
	// DOCS:  Industrial regions generate low-medium income, but are necessary for colonization efforts, merchant
	//  shipping, and the maintenance of the military
	Industrial uint64 `json:"industrial"`
	// DOCS:  Commercial regions potentially generate the greatest income, but are highly sensitive to taxation policy,
	//  unrest, wars, inflation, and deficit spending.
	Commercial         uint64 `json:"commercial"`
	MilitaryIndustrial uint64 `json:"militaryIndustrial"`
	Bureaucratic       uint64 `json:"bureaucratic"`
	Urban              uint64 `json:"urban"`
	Residential        uint64 `json:"residential"`
	Education          uint64 `json:"education"`
	Research           uint64 `json:"research"`
}

// newSystem is called when a successful colonization takes place.  It receives an empire object, determines the size
// of the new system, finds the next system ID for the specific empire's array of systems, and returns a fully populated
// system struct.
// TODO: rather than returning the new system, do we want to just append it?
func (empire *Empire) newSystem() System {
	// TODO: We want a weighted random number, with the most common being 100, and moving up in units of 50
	// TODO: https://github.com/mroth/weightedrand
	newSize := rand.Intn(400) + 100
	newID := empire.Resources.Systems[len(empire.Resources.Systems)-1].Next
	newSystem := System{
		Id:   newID,
		Size: uint64(newSize),
		Regions: Regions{
			Solar:              0,
			Wind:               0,
			Mining:             0,
			Tourism:            0,
			Industrial:         0,
			Commercial:         0,
			MilitaryIndustrial: 0,
			Bureaucratic:       0,
			Urban:              0,
			Residential:        0,
			Education:          0,
			Research:           0,
		},
		Next: newID + 1,
	}
	return newSystem
}
