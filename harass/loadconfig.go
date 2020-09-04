package harass

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type HarassingConfig struct {
	List []WebConfigItem `json:"list"`
}

func (hc *HarassingConfig) Info() {
	fmt.Printf("config has %d webConfigitem \n", len(hc.List))
}
func LoadHarassingConfig(config *HarParamConfig) *HarassingConfig {
	b, err := ioutil.ReadFile(config.ConfigPath)
	if err != nil {
		fmt.Print(err)
	}
	cs := strings.Replace(string(b), "${phone}", config.Phone, -1)
	cs = strings.Replace(cs, "${name}", config.Name, -1)
	hc := &HarassingConfig{}
	err = json.Unmarshal([]byte(cs), &hc)
	if err != nil {
		log.Fatalln(err)
	}
	return hc
}
