package models

import (
	"encoding/xml"
	"sync"
)

var mutex sync.RWMutex
var gameIdInfo [][]string
var gameIdInfoIntl [][]string

type BuildConfig struct {
	GameId      int    `json:"gameId"`
	GameKey     string `json:"gameKey"`
	OpenId      string `json:"openId"`
	IsWrite     bool   `json:"isWrite"`
	SdkVersion  string `json:"sdkVersion"`
	SdkRegion   string `json:"sdkRegion"`
	Arm32       bool 	`json:"arm32"`
	Arm64		bool	`json:"arm64"`	
	X86			bool 	`json:"x86"`
	X86_64		bool 	`json:"x86_64"`
	PackageName string `json:"packageName"`
	AppName     string `json:"appName"`
	Aac         bool   `json:"aac"`
	Apnp        bool   `json:"apnp"`
	NowTime     string `json:"nowTime"`
}

//xml start

type Resources struct {
	XMLName  xml.Name      `xml:"resources"`
	Integers []integerItem `xml:"integer"`
	Strings  []stringItem  `xml:"string"`
	Bools    []boolItem    `xml:"bool"`
}

type integerItem struct {
	XMLName xml.Name `xml:"integer"`
	Name    string   `xml:"name,attr"`
	Text    string   `xml:",innerxml"`
}

type stringItem struct {
	XMLName xml.Name `xml:"string"`
	Name    string   `xml:"name,attr"`
	Text    string   `xml:",innerxml"`
}

type boolItem struct {
	XMLName xml.Name `xml:"bool"`
	Name    string   `xml:"name,attr"`
	Text    string   `xml:",innerxml"`
}

//xml end

func init() {
	_, _, gameIdInfo = GetGameIdInfo("tersafe2")
	_, _, gameIdInfoIntl = GetGameIdInfo("anogs")
}
