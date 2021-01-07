package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Instance struct {
	VCPU   float64
	VRAM   float64
	Counts float64
}

func ReadFile() ([]byte, error) {
	var filePath string
	fmt.Scanln(&filePath)
	if !strings.HasSuffix(filePath,".json"){
		fmt.Println("Invalid file, enter path again")
		fmt.Scanln(&filePath)
	}else{
	}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}
func main() {
	oldMap := make(map[string]Instance)
	for {
		fmt.Println("Please input path:")
		data, _ := ReadFile()
		var result map[string][]interface{}
		newMap := make(map[string]Instance)
		if err := json.Unmarshal(data, &result); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}
		for _, v := range result["Instances"] {
			//fmt.Printf("%T",result["Instances"])
			instance := v.(map[string]interface{})
			newMap[instance["type"].(string)] = Instance{
				instance["vCPU"].(float64),
				instance["vRam"].(float64),
				instance["counts"].(float64),
			}
		}
		for key, _ := range newMap {
			_, ok := oldMap[key]
			if ok == false {
				fmt.Println(key, "Provision", newMap[key].Counts)
			} else {
				if newMap[key].Counts > oldMap[key].Counts {
					fmt.Println(key, "Provision", newMap[key].Counts-oldMap[key].Counts)
				} else if newMap[key].Counts < oldMap[key].Counts {
					fmt.Println(key, "Delete", oldMap[key].Counts-newMap[key].Counts)
				} else {

				}
			}
		}
		oldMap = newMap
	}
}
