package main

import (
	"encoding/json"
	"log"
	"os"
)

func ReadJStr() Str {
	rawDataIn, err := os.ReadFile("Usr.json")
	if err != nil {
		log.Fatal("Cannot load settings:", err)
	}
	var settings Str
	err = json.Unmarshal(rawDataIn, &settings)
	if err != nil {
		log.Fatal("Invalid settings format:", err)
	}
	return settings
}

func (settings Str) WriteJStr() {
	rawDataOut, err := json.MarshalIndent(&settings, "", "  ")
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}
	err = os.WriteFile("Usr.json", rawDataOut, 0)
	if err != nil {
		log.Fatal("Cannot write updated settings file:", err)
	}
}

func ReadJIt() It {
	rawDataIn, err := os.ReadFile("Itm.json")
	if err != nil {
		log.Fatal("Cannot load settings:", err)
	}
	var settings It
	err = json.Unmarshal(rawDataIn, &settings)
	if err != nil {
		log.Fatal("Invalid settings format:", err)
	}
	return settings
}

func (settings *It) WriteJIt() {
	rawDataOut, err := json.MarshalIndent(&settings, "", "  ")
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}

	err = os.WriteFile("Itm.json", rawDataOut, 0)
	if err != nil {
		log.Fatal("Cannot write updated settings file:", err)
	}
}
