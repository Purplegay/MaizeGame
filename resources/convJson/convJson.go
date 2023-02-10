package convjson

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"rpctest/resources/define"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

//https://blog.csdn.net/u013645668/article/details/120362201

type meta struct {
	Key  string
	Idx  int
	Typ  string
	Desc string
}

type rowdata []interface{}

func ParseFile(file string) {

	fmt.Println("\n\n\n\n", file)

	xlsx, err := excelize.OpenFile(file)
	if err != nil {
		panic(err.Error())
	}
	//[line][colidx][data]

	sheets := xlsx.GetSheetMap()
	for _, s := range sheets {
		if strings.HasPrefix(s, "!") {
			continue
		}

		rows := xlsx.GetRows(s)
		if len(rows) < 4 {
			return
		}

		colNum := len(rows[1])
		fmt.Println("col num:", colNum)
		metaList := make([]*meta, 0, colNum)
		dataList := make([]rowdata, 0, len(rows)-3)

		for line, row := range rows {
			switch line {
			case 0: // col name

				for idx, colname := range row {
					fmt.Println(idx, colname, len(metaList))

					metaList = append(metaList, &meta{Key: colname, Idx: idx})
				}
			case 1: // data type

				fmt.Printf("meta cot:%d, rol cot:%d\n", len(metaList), len(row))
				for idx, typ := range row {

					typ = strings.Replace(typ, "long", "int", -1)
					typ = strings.Replace(typ, "int", "int32", -1)
					typ = strings.Replace(typ, "float", "float64", -1)
					typ = strings.Replace(typ, "int3264", "int64", -1)

					if strings.Contains(typ, "[]") {
						typ = "slc|" + strings.TrimRight(typ, "[]")
					}
					metaList[idx].Typ = typ
				}
			case 2: // desc
				for idx, desc := range row {
					metaList[idx].Desc = desc
				}

			default: //>= 3 row data
				data := make(rowdata, colNum)

				for k := 0; k < colNum; k++ {
					if k < len(row) {
						data[k] = row[k]
					}
				}

				dataList = append(dataList, data)
			}

		}

		//sheetName := xlsx.GetSheetName(idx)
		// to json, save
		filename := filepath.Base(file)
		suf := filepath.Ext(filename)
		realName := fmt.Sprintf("%s_%s", filename[:(len(filename)-len(suf))], s)
		jsonFile := fmt.Sprintf("%s.json", realName)
		err = output(jsonFile, toJson(dataList, metaList))
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(toJson(dataList, metaList))
		err, goData := SplicingData(metaList, realName)
		if err != nil {
			fmt.Printf("SplicingData err:%v", err)
		}
		if err := writeNewFile(realName, goData); err == nil {
			path := define.GoExcelPath + realName + ".go"
			fmt.Println(path)
			cmd := exec.Command("gofmt", "-w", path)
			err := cmd.Run()
			if nil != err {
				fmt.Print(err)
			}
		}

	}

}

func toJson(datarows []rowdata, metalist []*meta) string {
	ret := "["

	for _, row := range datarows {
		ret += "\n\t{"
		for idx, meta := range metalist {
			ret += fmt.Sprintf("\n\t\t\"%s\":", meta.Key)
			if meta.Typ == "string" {
				if row[idx] == nil {
					ret += "\"\""
				} else {
					ret += fmt.Sprintf("\"%s\"", row[idx])
				}
			} else {
				if row[idx] == nil || row[idx] == "" {
					ret += "0"
				} else {
					ret += fmt.Sprintf("%s", row[idx])
				}
			}
			ret += ","
		}
		ret = ret[:len(ret)-1]

		ret += "\n\t},"
	}
	ret = ret[:len(ret)-1]

	ret += "\n]"
	return ret
}

func output(filename string, str string) error {

	f, err := os.OpenFile(define.OutjsonPath+filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(str)
	if err != nil {
		return err
	}

	return nil
}
