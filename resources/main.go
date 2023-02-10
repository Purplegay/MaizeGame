package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	convjson "rpctest/resources/convJson"
	"rpctest/resources/define"
	"strings"
)

func getFileList(path string) []string {
	var all_file []string
	finfo, _ := ioutil.ReadDir(path)
	for _, info := range finfo {
		if strings.HasPrefix(info.Name(), "~") {
			continue
		}
		if filepath.Ext(info.Name()) == ".xlsx" {
			real_path := path + "/" + info.Name()
			if info.IsDir() {
				//all_file = append(all_file, getFileList(real_path)...)
			} else {
				all_file = append(all_file, real_path)
			}
		}
	}

	return all_file
}

func main() {
	filelist := getFileList(define.ExcelPath)
	fmt.Println(filelist)

	for _, file := range filelist {
		convjson.ParseFile(file)
	}
}
