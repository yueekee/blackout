package main

import (
	"encoding/json"
	"github.com/liankui/blackout/utils"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./conf/config.json")
	if err != nil {
		panic(err)
	}

	//fmt.Printf("json :%s\n", data)
	e := json.Unmarshal(data, &utils.GlobalObject)
	if e != nil {
		panic(e)
	}
}
