package models

import (
	"ace-img2/tools"
	"path/filepath"
	"strings"
)

func UpLoadSdk(fname string) (int, string) {

	var err error
	var temppath string
	if strings.Contains(fname, "_ano") {
		temppath = filepath.Join("static", "anosdks")
	} else {
		temppath = filepath.Join("static", "acesdks")
	}

	err = tools.MyUnzipToSdks(fname, temppath)
	if err != nil {
		return 5004, err.Error()
	}
	return 0, "ok"
}
