package ReadFile

import (
	"fmt"
	"io/ioutil"
)

func CommandFileReader() []byte{
	file,err:=ioutil.ReadFile("D:/GoProject/src/BitcoinWeb/controllers/ReadFile/CommandFile/command.txt")
	if err != nil {
		fmt.Println("read fail",err.Error())
	}
	return file
}