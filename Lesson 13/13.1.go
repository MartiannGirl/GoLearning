// Необходимо распарсить json {"number":1,"landlord":"Остап
// Бендер","tenat":"Шура Балаганов»}
package main

import (
	"encoding/json"
	"fmt"
)

type contract struct {
	Number    int
	Sign_date string
	Landlord  string
	Tenat     string
}

func main() {
	str := `{"number":1, "sign_date":"2023-09-11", "landlord":"БендерОстап", "Tenat":"БалагановШура"}`

	c := contract{}

	err := json.Unmarshal([]byte(str), &c)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", c)
	fmt.Println()

}
