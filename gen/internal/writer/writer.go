package writer

import (
	"io/ioutil"
	"log"
	"os"
)

func WriteFile(fileName string,data []byte)  {
	err := ioutil.WriteFile(fileName,data,os.FileMode(os.ModePerm))
	if err != nil {
		log.Fatal(err)
	}
}

func CreateFolder(path string)  {
	err:=os.MkdirAll(path, os.ModePerm)
	if err != nil{
		log.Fatal(err)
	}
}