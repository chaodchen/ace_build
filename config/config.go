package config

import (
	"io/ioutil"
	"log"
	"os"
	"encoding/json"
)

type MyConfig struct {
	LogPath  string `json:"logPath"`
	OutPuth  string `json:"outPath"`
	Debug    bool   `json:"debug"`
	Host     string `json:"host"`
	Port     uint16 `json:"port"`
	DemoPath string `json:"demoPath"`
}

var Config MyConfig

func init() {
	var pkgName = "config.init"
	propFile, err := os.OpenFile("config/prop.json", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalf("[%s] %s", pkgName, err.Error())
	}
	byteValue, _ := ioutil.ReadAll(propFile)
	propFile.Close()
	json.Unmarshal(byteValue, &Config)
	log.Printf("[%s] Config: %v", pkgName, Config)
}