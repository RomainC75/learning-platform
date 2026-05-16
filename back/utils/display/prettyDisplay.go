package utils_display

import (
	"encoding/json"
	"fmt"
)

func PrettyDisplay(title string, data any) {
	fmt.Printf("_________%s_________", title)
	jsonRes, _ := json.MarshalIndent(data, "", " ")
	fmt.Println(string(jsonRes))
}
