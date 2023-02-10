package convjson

import (
	"fmt"
	"log"
	"os"
	"rpctest/resources/define"
	"strings"
)

//https://github.com/ljhe/generateStruct/blob/master/tool/generateStruct.go

var (
	structBegin    = "type %s struct {\n" // 结构体开始
	commonstruct   = " %s %s"
	structRemarks  = "	 // %s"        // 结构体备注
	structValueEnd = "\n"             // 结构体内容结束
	structEnd      = "}\n"            // 结构体结束
	header         = "package %s\n\r" // 文件头
)

func SplicingData(meta []*meta, structName string) (error, string) {
	outData := ""
	structData := fmt.Sprintf(structBegin, firstRuneToUpper(structName))

	for _, value := range meta {
		if isSpecialType(value.Typ) {
			structData += makeSpecialStr(structName, value.Key, value.Typ)
		} else {
			structData += fmt.Sprintf(commonstruct, value.Key, value.Typ)
		}

		if value.Desc != "" {
			structData += fmt.Sprintf(structRemarks, value.Desc)
		}
		structData += structValueEnd
	}

	structData += structEnd
	outData += structData
	return nil, outData
}

// 字符串首字母转换成大写
func firstRuneToUpper(str string) string {
	data := []byte(str)
	for k, v := range data {
		if k == 0 {
			first := []byte(strings.ToUpper(string(v)))
			newData := data[1:]
			data = append(first, newData...)
			break
		}
	}
	return string(data[:])
}

//检测是否是特殊类型
func isSpecialType(attrType string) bool {
	if strings.HasPrefix(attrType, "slc") {
		return true
	} else if strings.HasPrefix(attrType, "map") {
		return true
	} else if strings.HasPrefix(attrType, "double_slc") {
		return true
	} else {
		return false
	}
}

func makeSpecialStr(sheetName, key, typ string) string {
	var str string
	if strings.HasPrefix(typ, "slc") {
		str = makeSliceStr(key, typ)
	} else if strings.HasPrefix(typ, "map") {
		str = makeMapStr(key, typ)
	} else if strings.HasPrefix(typ, "double_slc") {
		str = makeDoubleSlice(key, typ)
	} else {
		log.Printf("无效的类型. Sheet:%s,列名:%s, 列属性名:%s\n", sheetName, key, typ)
	}
	return str
}

func makeSliceStr(key, typ string) string {
	strSlc := strings.Split(typ, "|")
	if len(strSlc) < 2 {
		return ""
	}
	str := key + "Slc"
	str += "[]" + strSlc[1]
	return str
}

// map|int|string
func makeMapStr(key, typ string) string {
	strSlc := strings.Split(typ, "|")
	if len(strSlc) < 3 {
		return ""
	}
	str := key + "Map "
	str += "map[" + strSlc[1] + "]" + strSlc[2]
	return str
}

func makeDoubleSlice(key, typ string) string {
	strSlc := strings.Split(typ, "|")
	if len(strSlc) < 2 {
		return ""
	}
	str := key + "DSlc"
	str += " [][]" + strSlc[1]
	return str
}

func writeNewFile(name string, data string) error {
	str := strings.Split(define.GoExcelPath, "\\")
	if len(str) == 0 {
		return fmt.Errorf("WriteNewFile|len(str) is 0")
	}
	header = fmt.Sprintf(header, "valuecode \n")
	data = header + data
	fw, err := os.OpenFile(define.GoExcelPath+name+".go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("WriteNewFile|OpenFile is err:%v", err)
	}
	defer fw.Close()
	_, err = fw.Write([]byte(data))
	if err != nil {
		return fmt.Errorf("WriteNewFile|Write is err:%v", err)
	}
	return nil
}
