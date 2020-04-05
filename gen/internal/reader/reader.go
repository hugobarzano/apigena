package reader

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func PathExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func ReadSpec(filePath string)[]byte  {
	data, err:=readFile(filePath)
	if err!=nil{
		log.Println(err.Error())
		return nil
	}
	return data
}

func GetReader(filePath string)(*os.File,error)  {
	if ok:= FileExists(filePath);!ok{
		return nil, errors.New(fmt.Sprintf("File: %v not found",filePath))
	}
	return os.Open(filePath)
}

func readFile(filePath string)([]byte,error)  {

	if ok:= FileExists(filePath);!ok{
		return nil, errors.New(fmt.Sprintf("File: %v not found",filePath))
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	defer file.Close()


	byteFile, err := ioutil.ReadAll(file)
	if err != nil{
		log.Fatal(err.Error())
	}
	return byteFile, err
}

