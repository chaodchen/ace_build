package models

import (
	"io/ioutil"
	"path/filepath"
)

func GetSdkList(sdkVersion string) (int, string, []string) {
	var sdksPath string
	var ret = []string{}

	if sdkVersion == "anogs" {
		sdksPath = filepath.Join("static", "anosdks")

	} else {
		sdksPath = filepath.Join("static", "acesdks")
	}

	files, err := ioutil.ReadDir(sdksPath)
	if err != nil {
		return 1001, err.Error(), ret
	}
	if len(files) < 1 {
		return 1002, err.Error(), ret
	}

	for _, item := range files {
		ret = append(ret, item.Name())
	}

	return 0, "ok", ret
}
