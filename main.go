package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var OperatingSystem = runtime.GOOS

func main() {
	rand.Seed(time.Now().UnixNano())

	/*playerEmpire := empire{
		Id:       1,
		Name:     "test empire",
		Prestige: 0,
		Interest: 1.0,
		Policies: policies{
			Taxation: 15,
			Priorities: constructionPriorities{
				Infantry:      25,
				Armor:         0,
				Turrets:       0,
				Fighters:      0,
				LightCruisers: 0,
				HeavyCruisers: 0,
				Scouts:        1,
				Frigates:      0,
				Destroyers:    0,
				Battleships:   0,
				Carriers:      0,
			},
			BalancedBudget: false,
		},
		Resources: resources{
			MegaCredits: 5000,
			Food:        500,
			Population:  100,
			Systems: []system{
				{
					Id:   1,
					Size: 100,
					Regions: regions{
						Solar:              5,
						Wind:               0,
						Mining:             2,
						Tourism:            2,
						Industrial:         2,
						Commercial:         1,
						MilitaryIndustrial: 0,
						Bureaucratic:       1,
						Urban:              2,
						Residential:        0,
						Education:          0,
						Research:           0,
					},
					Next: 2,
				},
			},
		},
	}*/
	// playerEmpire := Empire{}
	// TestingEmpireFunctionality(playerEmpire)
	// TestingGetPassword()
	// defaultVal := Default[string]{"t"}
	turnLog := TurnLog{}

	p, d := NewPrompt[PromptBool]("Do you wish to proceed?", false, &turnLog)
	fmt.Printf("%#v\n", getBool[PromptBool](p, d))

	p, e := NewPrompt[PromptUInt64]("How many shall we buy?", 0, &turnLog)
	fmt.Printf("%d\n", getUInt[PromptUInt64](p, e))

	p, f := NewPrompt[PromptString]("Login name?", "", &turnLog)
	fmt.Printf("%s\n", getString[PromptString](p, f))

}
