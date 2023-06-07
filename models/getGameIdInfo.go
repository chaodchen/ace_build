package models

import (
	"io/ioutil"
	"strings"
)

func GetGameIdInfo(sdkVersion string) (int, string, [][]string) {
	var gameIdInfoPath string
	var ret = [][]string{}
	if sdkVersion == "anogs" {
		gameIdInfoPath = "static/gameIdInfoIntl.txt"
		if len(gameIdInfoIntl) > 0 {
			return 0, "ok", gameIdInfoIntl
		}
	} else {
		gameIdInfoPath = "static/gameIdInfo.txt"
		if len(gameIdInfo) > 0 {
			return 0, "ok", gameIdInfo
		}
	}

	// mutex.
	file, err := ioutil.ReadFile(gameIdInfoPath)
	if err != nil {
		return 3001, err.Error(), ret
	}

	temparr := strings.Split(string(file), "\n")
	for i := range temparr {
		tempitem := strings.Split(temparr[i], ",")
		// append(ret, tempitem)
		ret = append(ret, tempitem)
	}

	return 0, "ok", ret
}
