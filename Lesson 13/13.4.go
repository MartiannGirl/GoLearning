package main

//type contracts struct {
//List []contract `xml:"contract"`
//}
//contract{
//Number: 1,
//Landlord: "Остап Бендер",
//tenat: "Шура Балаганов",
//
//}
import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type contract struct {
	Number   int    `xml:"number"`
	Landlord string `xml:"landlord"`
	Tenat    string `xml:"tenat"`
}

type contracts struct {
	List []contract `xml:"contract"`
}

func main() {
	fileName := "contracts.xml"
	f, err := os.Create(fileName)

	if err != nil {
		log.Fatalln(err)
	}
	f.WriteString(xml.Header)

	c1 := contract{
		Number:   1,
		Landlord: "Остап Бендер",
		Tenat:    "Шура Балаганов",
	}

	d := contracts{}
	d.List = append(d.List, c1)
	enc := xml.NewEncoder(f)
	enc.Indent("", "    ")
	err = enc.Encode(d)
	if err != nil {
		log.Fatalln(err)
	}

	f.Close()

	f, err = os.Open(fileName)
	if err != nil {

		log.Fatalln(err)

	}

	defer f.Close()
	res := contracts{}
	err = xml.NewDecoder(f).Decode(&res)

	if err != nil {

		log.Fatalln(err)
	}
	fmt.Printf("%+v", res)
	fmt.Println()
}
