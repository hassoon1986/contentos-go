package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type TableInfo struct {
	Name 		string
	PList		[]PropList
}


type PropList struct {
	VarType		string
	VarName		string
	BMainKey	bool
	BSeckey     bool
	BUnique		bool
	BSort		bool
	Index		uint32
}

var TmlFolder = "./app/table/tools/tml/"


func (p *PropList) ToString() string{
	s := fmt.Sprintf("\t%s\t%s = %d;\n", p.VarType, p.VarName, p.Index)
	return s
}

func (p *PropList) Parse(info []string, index uint32) bool {
	p.VarType	= info[0]
	p.VarName	= info[1]
	res, err := strconv.ParseBool(strings.Replace(info[2]," ","",-1))
	if err != nil{
		return false
	}
	p.BMainKey = res

	resSecKey,errSecKey := strconv.ParseBool(strings.Replace(info[5]," ", "", -1))
    if errSecKey != nil {
    	return false
	}
	p.BSeckey = resSecKey

	resUni, errUni  := strconv.ParseBool(strings.Replace(info[3]," ","",-1))
	if errUni != nil{
		return false
	}
	p.BUnique	= resUni

	resSort, errSort	:= strconv.ParseBool(strings.Replace(info[4]," ","",-1))
	if errSort != nil{
		return false
	}
	p.BSort		= resSort

	p.Index		= index

	if index == 1 && !p.BMainKey{
		return false
	}
	if index > 1 && p.BMainKey{
		return false
	}

	return true
}

func ExtractPropListToPB( pl []PropList, name string) string {
	result := fmt.Sprintf("message %s {\n", name)

	for _, val := range pl {
		result = fmt.Sprintf("%s%s", result, val.ToString() )
	}

	result = fmt.Sprintf("%s}\n", result )

	return result
}

func ExtractPropListToGoFile( pl []PropList, name string) string {
	result := fmt.Sprintf("message %s {\n", name)

	for _, val := range pl {
		result = fmt.Sprintf("%s%s", result, val.ToString() )
	}

	result = fmt.Sprintf("%s}\n", result )

	return result
}

func ProcessCSVFile(fileName string, name string) bool {
	inBuff,err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	r2 		:= csv.NewReader(strings.NewReader(string(inBuff)))
	lines,_ := r2.ReadAll()

	var indexPb	uint32 = 0
	sz := len(lines)
	props := make([]PropList, 0)

	for i:=1;i<sz;i++ {
		line := lines[i]

		if len(line[0]) <= 0 {
			continue
		}
		if indexPb == 0{
			indexPb = 1
		}

		pList := &PropList{}

		if !pList.Parse( line, indexPb) {

			fmt.Println("parse line error:", line )
			panic( nil )
		}
		indexPb++
		props = append(props, *pList)
	}


	//var res = ExtractPropListToPB( props, "so_"+name )
	//fmt.Println(res)
	//
	//
	//res = ExtractPropListToGoFile( props, "so_"+name )
	//fmt.Println(res)


	tInfo := TableInfo{name, props}


	 wRes,_ := WritePbTplToFile(tInfo)
	 if wRes {
	 	//auto create pb.go file
	 	 cmd := exec.Command("protoc","-I./",
			 "--go_out=./",
			 TmlFolder+"so_"+name+".proto")
		 err := cmd.Start()
		 if err == nil {
		 	//create detail go file (include update insert delete functions)
			 cRes,cErr := CreateGoFile(tInfo)
			 if !cRes {
				 log.Printf("create go file fail,name prefix is %s , error is %s \n",tInfo.Name,cErr)
			 }
		 }else {
			 log.Printf("create pb file fail,cmd error is %s\n",err)
		 }
	 }



	return true
}

func main(){
	var dirName string = "./app/table/table-define"
	pthSep := string(os.PathSeparator)
	fmt.Println(pthSep)
	files, _ := ioutil.ReadDir(dirName)
	for _, f := range files {
		if f.IsDir(){
			continue
		}
		if strings.HasSuffix( f.Name(), ".csv" ){
			ProcessCSVFile(dirName + pthSep + f.Name(), f.Name()[0:len(f.Name())-4 ] )
		}
	}
}


/* writte tmplate content into proto file */
func WritePbTplToFile(tInfo TableInfo) (bool, error) {
	var err error = nil
	if tpl := createPbTpl(tInfo); tpl != "" {
		if isExist, _ := JudgeFileIsExist(TmlFolder); !isExist {
			//文件夹不存在,创建文件夹
			if err := os.Mkdir(TmlFolder, os.ModePerm); err != nil {
					log.Printf("create folder fail,the error is:%s \n", err)
			 		return false,err
			 }
			}
		fName := TmlFolder + "so_"+tInfo.Name + ".proto"
		if fPtr := CreateFile(fName); fPtr != nil {
			t := template.New("layout.html")
			t.Parse(tpl)
			t.Execute(fPtr,tInfo)
			cmd := exec.Command("goimports", "-w", fName)
			cmd.Start()
			defer fPtr.Close()
			return true,nil
		}else {
			err = errors.New("get file ptr fail")
			log.Println("get file ptr fail")
		}
	}else{
		err = errors.New("create tpl fail")
		log.Println("create tpl fail")
	}

	return false,err
}

/* create pb template */
func createPbTpl(t TableInfo) string  {
	if len(t.PList) > 0 {
		tpl := ""
		tpl =  `
syntax = "proto3";

package table;

import "common/prototype/type.proto";

message so_{{.Name}} {
	{{range $k,$v := .PList}}
       {{.VarType}}   {{.VarName}}     =      {{.Index}};
	{{end}}  
}
`
		tpl = fmt.Sprintf("%s%s",tpl,createKeyTpl(t))
		return tpl
	}
	return ""
}

func createKeyTpl(t TableInfo) string {
	tpl := ""
	if len(t.PList) > 0 {
		var secKeyList = make([]PropList,0)
		mKeyType , mKeyName := "",""
		for _,v := range t.PList {
			if v.BMainKey {
				mKeyType = v.VarType
				mKeyName = v.VarName
			}else if v.BSeckey {
				secKeyList = append(secKeyList,v)
			}
		}
		if len(secKeyList) > 0 && mKeyType != "" && mKeyName != ""{
			for _,v := range secKeyList {

				tempTpl := fmt.Sprintf("\nmessage so_key_%s_by_%s {\n",
					strings.Replace(t.Name," ","",-1),
					strings.Replace(v.VarName," ","",-1))
				tempTpl = fmt.Sprintf("%s\t\t%s\t%s\t=\t%d;\n\t\t%s\t%s\t=\t%d;\n}\n",tempTpl, v.VarType, v.VarName,1, mKeyType, mKeyName,2)
				tpl += fmt.Sprintf("%s",tempTpl)
			}
		}
	}
	return tpl
}

/* create detail file */
func CreateFile(fileName string) *os.File {
	var fPtr *os.File
	isExist, _ := JudgeFileIsExist(fileName)
	if !isExist {
		//create file
		if f, err := os.Create(fileName); err != nil {
			log.Printf("create detail go file fail,error:%s\n", err)
		} else {
			fPtr = f
		}
	} else {
		//rewrite the file
		if f, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm); f != nil {
			fPtr = f
		}
	}

	if fPtr == nil {
		log.Fatal("File pointer is empty \n")

	}
	return fPtr

}


/* judge if the file exists */
func JudgeFileIsExist(fPath string) (bool, error) {
	if fPath == "" {
		return false, errors.New("the file path is empty")
	}
	_, err := os.Stat(fPath)
	if err == nil {
		return true, nil
	} else if !os.IsNotExist(err) {
		return true, err
	}
	return false, err
}