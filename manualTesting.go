package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// TestingEmpireFunctionality currently reading a saved empire from disk, adding 3 systems to it, printing the raw struct to STDIO, then writing
// the changes to the data file
func TestingEmpireFunctionality(playerEmpire Empire) {
	playerEmpire.ReadEmpireFile()
	playerEmpire.Resources.Systems = append(playerEmpire.Resources.Systems, playerEmpire.newSystem())
	playerEmpire.Resources.Systems = append(playerEmpire.Resources.Systems, playerEmpire.newSystem())
	playerEmpire.Resources.Systems = append(playerEmpire.Resources.Systems, playerEmpire.newSystem())
	fmt.Printf("%+v\n", playerEmpire)
	playerEmpire.WriteEmpireFile()
}

func TestingGetPassword() {
	fmt.Print("Enter a password => ")
	input := bufio.NewReader(os.Stdin)
	pwd, _ := input.ReadString('\n')
	fmt.Printf(pwd)

	params := params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	hash, err := generateFromPassword(pwd, &params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hash)

	match, err := comparePasswordAndHash(pwd, hash)
	fmt.Printf("Match: %v\n", match)
}
