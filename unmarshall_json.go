package main

import (
	"encoding/json"
	"fmt"
)

type Credit struct {
	GoodNum int
	Level   int
}
type User struct {
	Name        string `json:"FelixName"`
	Age         int
	BuyerCredit Credit
}

type Row struct {
	Key1 string
	Key2 string
	Key3 string
}

func main() {
	s := `{"FelixName": "felix", "age": 12, "buyercredit": {"GoodNum": 12, "Level": 10}}`
	var simpleR User
	json.Unmarshal([]byte(s), &simpleR)
	fmt.Println(simpleR)
	b, _ := json.Marshal(simpleR)
	fmt.Println(string(b))

	json1 := `{"a": {"b": {"c": {"key1": "1", "key2": "2", "key3": "3"}}}}`
	json2 := `{"d": {"e": {"f": {"key1": "4", "key2": "5", "key3": "6"}}}}`
	json3 := `{"h": {"i": {"j": {"key1": "7", "key2": "8", "key3": "9"}}}}`

	var val map[string]json.RawMessage

	jsons := []string{json1, json2, json3}

	for _, j := range jsons {
		var row *Row
		valInLoop := j
		for {
			json.Unmarshal([]byte(valInLoop), &val)
			fmt.Println(len(val))
			if len(val) == 0 {
				break
			}
			if len(val) == 1 {
				for _, rawJson := range val {
					valInLoop = string(rawJson)
					fmt.Println(valInLoop)
				}
			}
			if len(val) > 1 {
				json.Unmarshal([]byte(valInLoop), &row)
				break
			}
		}
		fmt.Println(row)
	}

	// body := `{"user_get_response":{"user":{"buyer_credit":{"good_num":0,"level":0,"score":0,"total_num":0},"created":"2009-10-19 16:29:41","last_visit":"2012-02-27 17:14:40","location":{},"nick":"qintb8","seller_credit":{"good_num":0,"level":0,"score":0,"total_num":0},"type":"C","uid":"08666adc22bf0093cd531db115ee1c3b","user_id":322176867}}}`

	// var mr map[string]interface{}
	// json.Unmarshal([]byte(body), &mr)
	// fmt.Println(mr)

	// var r *User
	// json.Unmarshal([]byte(body), &r)

	// fmt.Println(r)
}
