package tools

import (
	// "archive/zip"
	// "bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"archive/zip"
	"path"
	"path/filepath"
	"strings"
    "fmt"
)

func CmdAndChangeDir(dir string, commandName string, params []string) (string, error) {
	cmd := exec.Command(commandName, params...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	return out.String(), err
}

func CopyFile(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()
	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer srcfd.Close()

	if _,err = io.Copy(dstfd, srcfd);err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src);err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

func DelectExistsFiles(src string) error {
    var err error
    var fds []os.FileInfo
    ace_files := []string {
        "libtersafe2.so", "libanogs.so",
        "ano.jar", "tp2.jar",
        "libtprt.so", "libanort.so",
    }

    if fds, err = ioutil.ReadDir(src); err != nil {
        return err
    }

    for _, fd := range fds {
        if !fd.IsDir() {
            // is file
            file_path := filepath.Join(src, fd.Name())
            fmt.Printf("file_path: %s\n", file_path)
            
            if (IsItemInArr(ace_files, fd.Name())) {
                fmt.Printf("ace_file in IsItemArr: %s\n", fd.Name())
                err = os.Remove(file_path) 
                if err != nil {
                    fmt.Printf("delect file failed.\n")
                    return err
                }
            }
        } 
    }

    return nil
}

func CopyDir(src string, dst string) error {
    var err error
    var fds []os.FileInfo
    var srcinfo os.FileInfo
 
    if srcinfo, err = os.Stat(src); err != nil {
        return err
    }
 
    if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
        return err
    }
 
    if fds, err = ioutil.ReadDir(src); err != nil {
        return err
    }
    for _, fd := range fds {
        srcfp := filepath.Join(src, fd.Name())
        dstfp := filepath.Join(dst, fd.Name())
 
        if fd.IsDir() {
            if err = CopyDir(srcfp, dstfp); err != nil {
                return err
            }
        } else {
            if err = CopyFile(srcfp, dstfp); err != nil {
				return err
            }
        }
    }
    return nil
}

func ReWriteFileLine(fpath string, tag string, content string) error {
	fcont, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}
	strs := strings.Split(string(fcont), "\n")
	buf := bytes.Buffer{}
	for _, str := range strs {
		if len(str) == 0 {
			continue
		}
		if strings.Contains(str, tag) == true {
			buf.WriteString(content)
		} else {
			buf.WriteString(str)
		}
		buf.WriteString("\n")
	}
	err = ioutil.WriteFile(fpath, buf.Bytes(), 0666)
	return err
}

func IsItemInArr(arr []string, item string) bool {
	for _, it := range arr {
		if item == it {
			return true
		}
	}
	return false
}

func DomOrVerSo(str string) string {
	if str == "demo" {
		return "libtersafe2.so"
	}
	return "libanogs.so"
}

func DomOrVerJar(str string) string {
	if str == "demo" {
		return "tp2.jar"
	}
	return "ano.jar"
}


func MyUnzipToSdks(s1 string, s2 string) error {
	tempstr := `product/sdk/android/java`
	tempstr2 := `ano_product/sdk/android/java`
	archive, err := zip.OpenReader(s1)
	if err != nil {
		return err
	}
	defer archive.Close()

	for _, f := range archive.File {
		if strings.HasPrefix(f.Name, tempstr) == false &&
			strings.HasPrefix(f.Name, tempstr2) == false {
			continue
		}

		localFilePath := strings.TrimSuffix(path.Base(s1), ".zip")
		localFilePath = path.Join(s2, localFilePath, f.Name)
		if _, err := os.Stat(localFilePath); err == nil {
			continue
		}
		
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(localFilePath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(localFilePath), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(localFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		localFile, err := f.Open()
		if err != nil {
			return err
		}

		if _, err := io.Copy(dstFile, localFile); err != nil {
			return err
		}
		
		dstFile.Close()
		localFile.Close()
	}

	return nil
}
