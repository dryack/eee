package main

type Military struct {
	Infantry      uint64 `json:"infantry"`      // defense against Infantry and armor (some defense against fighters), attack against Infantry/armor during invasions
	Armor         uint64 `json:"armor"`         // planetary defense against Infantry and Armor
	Turrets       uint64 `json:"turrets"`       // planetary defense, very strong attacking spacecraft but die in droves
	Fighters      uint64 `json:"fighters"`      // attack Fighters, support Infantry/armor, attack frigates/destroyers/lightCruisers/heavyCruisers
	LightCruisers uint64 `json:"lightCruisers"` // exploration,defend against Fighters, fleet defense, support Infantry/armor, anti-piracy
	HeavyCruisers uint64 `json:"heavyCruisers"` // attack smaller craft, planetary attack, anti-piracy, PRESTIGE
	Scouts        uint64 `json:"scouts"`        // exploration, bonus to planetary attack, bonus to targeting for fleet
	Frigates      uint64 `json:"frigates"`      // tackle and hold larger craft, fleet defense, raiding, anti-piracy
	Destroyers    uint64 `json:"destroyers"`    // fleet defense, anti-piracy, attack larger craft
	Battleships   uint64 `json:"battleships"`   // planetary attack, attack other Battleships, attack cruisers, PRESTIGE
	Carriers      uint64 `json:"carriers"`      // carry Fighters on offensive operations (not needed for defensive Fighters only), PRESTIGE
}

type ConstructionPriorities struct {
	Infantry      uint64 `json:"infantry"`
	Armor         uint64 `json:"armor"`
	Turrets       uint64 `json:"turrets"`
	Fighters      uint64 `json:"fighters"`
	LightCruisers uint64 `json:"lightCruisers"`
	HeavyCruisers uint64 `json:"heavyCruisers"`
	Scouts        uint64 `json:"scouts"`
	Frigates      uint64 `json:"frigates"`
	Destroyers    uint64 `json:"destroyers"`
	Battleships   uint64 `json:"battleships"`
	Carriers      uint64 `json:"carriers"`
}

// DOCS: Your piracy task force will patrol your systems, rooting out pirates and reducing their depredations against
//  your merchant shipping
// TODO: piracy task force
// DOCS: When you seek to attack another player, you will be given the option to assign a strike group for the attack.
//  Unlike with piracy task forces, strike groups are disbanded after their attack (assuming they aren't destroyed)
// TODO: strike groups
