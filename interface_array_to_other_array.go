package main

import (
	"fmt"
	"reflect"
)

func main() {
	mapArray := []interface{}{
		map[string]interface{}{"key1": 1},
		map[string]interface{}{"key2": "2"},
	}

	fmt.Println(mapArray)

	fmt.Println(reflect.TypeOf(mapArray))

	var newMapArray []map[string]interface{}
	/*	newMapArray = mapArray*/
	/*	newMapArray = []map[string]interface{}(mapArray)*/
	for _, v := range mapArray {
		mapValue, ok := v.(map[string]interface{})
		if ok {
			newMapArray = append(newMapArray, mapValue)
		}
	}

	fmt.Println(newMapArray)
	fmt.Println(reflect.TypeOf(newMapArray))

}
