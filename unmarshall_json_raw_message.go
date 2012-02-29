package main

import (
	"encoding/json"
	"fmt"
)

type Row struct {
	Key1 string
	Key2 string
	Key3 string
}

func main() {

	json1 := `{"a": {"b": {"c": {"key1": "1", "key2": "2", "key3": "3"}}}}`
	json2 := `{"d": {"e": {"f": {"key1": "4", "key2": "5", "key3": "6"}}}}`
	json3 := `{"h": {"i": {"j": {"key1": "7", "key2": "8", "key3": "9"}}}}`

	jsons := []string{json1, json2, json3}

	for _, j := range jsons {
		var row *Row
		valInLoop := j
		for {
			var val map[string]json.RawMessage
			json.Unmarshal([]byte(valInLoop), &val)
			fmt.Println(len(val))
			if len(val) == 0 {
				break
			}
			if len(val) == 1 {
				for _, rawJson := range val {
					valInLoop = string(rawJson)
					fmt.Println(valInLoop)
					break
				}
			}
			if len(val) > 1 {
				json.Unmarshal([]byte(valInLoop), &row)
				break
			}
		}
		fmt.Println(row)
	}

}
