package main

import "fmt"

func main() { // 8.1
	map1 := map[string]string{
		"слон":    "",
		"бегемот": "",
		"носорог": "",
		"лев":     "",
	}
	fmt.Println("карта map1: ", map1)

	map2 := map[string]string{ //8.2 не поняла :(
		"слон":    "слон",
		"бегемот": "бегемот",
		"носорог": "носорог",
		"лев":     "лев",
	}
	fmt.Println("карта map2: ", map2)
	fmt.Println("Животное: %s, количество: %d (есть в карте: %v)")

	map3 := map[string]string{ //8.3
		"слон":    "",
		"бегемот": "",
		"носорог": "",
		"лев":     "",
	}
	fmt.Println("карта map3: ", map3)
	delete(map3, "key_2")
	fmt.Println("карта map3: ", map3)

	map4 := map[string]string{ //8.4
		"слон":    "",
		"бегемот": "",
		"носорог": "",
		"лев":     "",
	}
	fmt.Println("карта map4: ", map4)
	map4["выдра"] = ""
	fmt.Println("карта map4: ", map4)

	map5 := map[string]int{ //8.5
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}
	fmt.Println("карта map5: ", map5)
	map5["бегемот"] = 2
	fmt.Println("карта map5: ", map5)
}
