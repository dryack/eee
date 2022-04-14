package main

import (
	"compress/gzip"
	"encoding/json"
	"log"
	"os"
)

func (empire *Empire) WriteEmpireFile() {
	f, err := os.OpenFile("./empire.dat", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	j, _ := json.Marshal(*empire)
	// f.Write(j)

	gf := gzip.NewWriter(f)
	defer gf.Close()
	gf.Write(j)
}

func (empire *Empire) ReadEmpireFile() {
	// TODO: filename per user?
	f, err := os.Open("./empire.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	gf, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer gf.Close()

	// result := make([]byte, 100)

	// _, _ = gf.Read(result)
	result := json.NewDecoder(gf)
	result.Decode(empire)
	if err != nil {
		log.Fatal(err)
	}
}
