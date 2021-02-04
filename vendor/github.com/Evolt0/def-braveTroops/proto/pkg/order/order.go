package order

import "encoding/json"

func Order(data interface{}) string {
	text, _ := json.Marshal(data)
	temp := make(map[string]interface{})
	_ = json.Unmarshal(text, &temp)
	Filter(temp)
	result, _ := json.Marshal(temp)
	return string(result)
}

func Filter(data map[string]interface{}) {
	if _, ok := data["sign"]; ok {
		delete(data, "sign")
	}
	for _, value := range data {
		if elem, ok := value.(map[string]interface{}); ok {
			Filter(elem)
		}
	}
}
