package models

import (
	"ace-img2/config"
	"ace-img2/tools"
	"strings"

	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func BuildDemo(data BuildConfig) (int, string) {
	var mutex sync.RWMutex
	var pkgName = "models.BuildDemo"
	mutex.Lock()
	defer mutex.Unlock()

	var valueXmlPath = filepath.Join(config.Config.DemoPath, "unityLibrary/src/main/res/values", "values.xml")

	log.Printf("[%s]data: %v\ndemoPath: %s\nvalueXmlPath: %s\n", pkgName, data.AppName, config.Config.DemoPath, valueXmlPath)

	if _, err := os.Stat(config.Config.DemoPath); err != nil {
		return 4002, err.Error()
	}
	if _, err := os.Stat(valueXmlPath); err != nil {
		return 4003, err.Error()
	}

	xmlFile, err := ioutil.ReadFile(valueXmlPath)
	if err != nil {
		return 4004, err.Error()
	}
	var xmlParse Resources
	if err := xml.Unmarshal(xmlFile, &xmlParse); err != nil {
		return 4005, err.Error()
	}

	log.Printf("[%s] Xml数据绑定成功\n", pkgName)
	for i, j := range xmlParse.Strings {
		switch j.Name {
		case "game_key":
			xmlParse.Strings[i].Text = data.GameKey
		case "open_id":
			xmlParse.Strings[i].Text = data.OpenId
		case "new_app_name":
			xmlParse.Strings[i].Text = data.AppName
		case "package_name":
			xmlParse.Strings[i].Text = data.PackageName
		case "region":
			xmlParse.Strings[i].Text = data.SdkRegion
		default:
			// log.Println("String switch fail...")
		}
	}

	for i, j := range xmlParse.Integers {
		switch j.Name {
		case "game_id":
			xmlParse.Integers[i].Text = strconv.Itoa(data.GameId)
		default:
			// log.Println("Integers switch fail.")
		}
	}

	for i, j := range xmlParse.Bools {
		switch j.Name {
		case "is_write_storage":
			if data.IsWrite {
				xmlParse.Bools[i].Text = "true"
			} else {
				xmlParse.Bools[i].Text = "false"
			}
		case "aac":
			if data.Aac {
				xmlParse.Bools[i].Text = "true"
			} else {
				xmlParse.Bools[i].Text = "false"
			}
		case "apnp":
			if data.Apnp {
				xmlParse.Bools[i].Text = "true"
			} else {
				xmlParse.Bools[i].Text = "false"
			}
		default:
			// log.Println("Bools switch fail.")
		}

		xmlContent, err := xml.MarshalIndent(xmlParse, " ", " ")
		if err != nil {
			return 4006, err.Error()
		}

		headerBytes := []byte(xml.Header)
		if err := ioutil.WriteFile(valueXmlPath,
			append(headerBytes, xmlContent...),
			os.ModeAppend); err != nil {
			return 4007, err.Error()
		}
	}

	// 开始编译
	var aceDemoJarPath string
	aceDemoJniLibsPath := filepath.Join(config.Config.DemoPath, "unityLibrary", "src", "main", "jniLibs")

	if _, err := os.Stat(aceDemoJniLibsPath); err != nil {
		return 4008, err.Error()
	}
	var localSdkPath string
	var localSdkJarPath string
	var localSdkLibsPath string
	var jarPathInProp string
	if data.SdkRegion == "anogs" {
		nowPath, _ := os.Getwd()
		localSdkPath = filepath.Join(nowPath, "static", "anosdks", data.SdkVersion, "ano_product", "sdk", "android", "java")
		aceDemoJarPath = filepath.Join(aceDemoJniLibsPath, "ano.jar")
		localSdkJarPath = filepath.Join(localSdkPath, "ano.jar")
		jarPathInProp = "src\\\\main\\\\jniLibs\\\\ano.jar"
	} else {
		nowPath, _ := os.Getwd()
		localSdkPath = filepath.Join(nowPath, "static", "acesdks", data.SdkVersion, "product", "sdk", "android", "java")
		aceDemoJarPath = filepath.Join(aceDemoJniLibsPath, "tp2.jar")
		localSdkJarPath = filepath.Join(localSdkPath, "tp2.jar")
		jarPathInProp = "src\\\\main\\\\jniLibs\\\\tp2.jar"
	}
	localSdkLibsPath = filepath.Join(localSdkPath, "lib")

	if _, err := os.Stat(localSdkPath); err != nil {
		return 4009, err.Error()
	}
	log.Printf("[%s] aceDemoJniLibsPath: %s\naceSdkPath: %s\n", pkgName, aceDemoJniLibsPath, localSdkPath)

	// 这里删除文件 应该要先判断一下权限 这里面有Unity的so 不能删除
	// dir, err := ioutil.ReadDir(aceDemoJniLibsPath)
	// for _, item := range dir {
	// 	tempPath := path.Join([]string{aceDemoJniLibsPath, item.Name()}...)
	// 	if _, err := os.Stat(tempPath); err == nil {
	// 		os.RemoveAll(tempPath)
	// 	}
	// }

	log.Printf("[%s] localSdkJarPath: %s\naceDemoJarPath: %s", pkgName, localSdkJarPath, aceDemoJarPath)

	// 复制jar包
	if err := tools.CopyFile(localSdkJarPath, aceDemoJarPath); err != nil {
		log.Printf("[%s] 复制文件失败: %s", pkgName, err.Error())
		return 4010, err.Error()
	}

	// 拼接arch
	var archstr = ""
	var archs []string
	if data.Arm32 {
		archs = append(archs, "armeabi-v7a")
	}

	if data.Arm64 {
		archs = append(archs, "arm64-v8a")
	}

	if data.X86 {
		archs = append(archs, "x86")
	}

	if data.X86_64 {
		archs = append(archs, "x86_64")
	}

	// 复制libs
	var fds []os.FileInfo
	// var err error
	if fds, err = ioutil.ReadDir(localSdkLibsPath); err != nil {
		return 4016, err.Error()
	}
	for _, fd := range fds {
		if fd.IsDir() {
			log.Printf("[%s] fds.Name(): %s", pkgName, fd.Name())
			for _, arch := range archs {
				if arch == fd.Name() {
					srcfp := filepath.Join(localSdkLibsPath, fd.Name())
					dstfp := filepath.Join(aceDemoJniLibsPath, fd.Name())
					log.Printf("[%s] srcfp: %s, dstfp: %s", pkgName, srcfp, dstfp)
					if err := tools.CopyDir(srcfp, dstfp); err != nil {
						return 4011, err.Error()
					}
					break
				}
			}
		}
	}

	archstr = strings.Join(archs, ":")
	log.Printf("[%s] archs: %v", pkgName, archs)
	log.Printf("[%s] archstr: %s", pkgName, archstr)

	// 写入gradle.properties
	var gradleProp = filepath.Join(config.Config.DemoPath, "gradle.properties")
	if err := tools.ReWriteFileLine(gradleProp, "PROP_APP_ABI", "PROP_APP_ABI="+archstr); err != nil {
		return 4012, err.Error()
	}
	if err := tools.ReWriteFileLine(gradleProp, "LIB_PATH", "LIB_PATH="+jarPathInProp); err != nil {
		return 4013, err.Error()
	}
	if err := tools.ReWriteFileLine(gradleProp, "APPLICATION_ID", "APPLICATION_ID="+data.PackageName); err != nil {
		return 4014, err.Error()
	}

	if _, err := tools.CmdAndChangeDir(filepath.Join(config.Config.DemoPath), "./gradlew", []string{"assembleRelease"}); err != nil {
		return 4015, err.Error()
	}

	return 0, "ok"
}
