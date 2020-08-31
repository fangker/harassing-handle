package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type HarassingConfig struct {
	List []WebConfigItem `json:"list"`
}

func (hc *HarassingConfig) Info() {
	fmt.Printf("config has %d webConfigitem \n", len(hc.List))
}
func loadHarassingConfig(config string) *HarassingConfig {
	hc := &HarassingConfig{}
	err := json.Unmarshal([]byte(config), &hc)
	if err != nil {
		log.Fatalln(err)
	}
	return hc
}
