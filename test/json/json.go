package main

import (

	io "io/ioutil"

	json "encoding/json"

	"fmt"

)

type JsonStruct struct{

}



func NewJsonStruct () *JsonStruct {

	return &JsonStruct{}

}



func (self *JsonStruct) Load (filename string, v interface{}) {

	data, err := io.ReadFile(filename)

	if err != nil{

		return

	}

	datajson := []byte(data)



	err = json.Unmarshal(datajson, v)

	if err != nil{

		return

	}

}



type ValueTestAtmp struct{

	StringValue string

	NumericalValue int

	BoolValue bool

}



type testdata struct {

	ValueTestA ValueTestAtmp

}



func main() {

	JsonParse := NewJsonStruct()

	v := testdata{}

	JsonParse.Load("jsonparse_config.txt", &v)

	fmt.Println(v)

	fmt.Println(v.ValueTestA .StringValue )

}



//jsonparse_config.txt:
//
//{
//
//"ValueTestA":{
//
//"StringValue": "127.1.1.1",
//
//"NumericalValue":1233,
//
//"BoolValue":false
//
//},
//
//"ValueTestB":{
//
//"FloatValue":123.456
//
//}
//
//}


