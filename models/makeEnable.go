package models

import (
	"ace-img2/tools"
	"log"
	"strings"
)



func MakeEnable(s string) (int, string) {
	var pkgName = "models.MakeEnable"

	log.Printf("[%s] Start MakeEnable.\n", pkgName)
	// 去头尾空格
	s = strings.TrimSpace(s)
	log.Printf("[%s] s.length: %d\n", pkgName, len([]rune(s)))
	if len([]rune(s)) < 24 {
		return 2001, "sign长度非法"
	}

	// 执行
	args := []string{"auth_enable.py"}
	args = append(args, strings.Split(s, ",")...)
	log.Printf("args: %s", args)
	mutex.Lock()
	_, err := tools.CmdAndChangeDir("static/auth_lua", "python2", args)
	mutex.Unlock()
	if err != nil {
		return 2002, "生成enable.dat失败"
	}
	return 0, "ok"

}
