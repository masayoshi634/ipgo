package command

import (
	"encoding/json"
	"fmt"
)

func toJson(t interface{}) {
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}

	fmt.Println(string(jsonBytes))
}
