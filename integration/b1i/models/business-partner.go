package models

type BusinessPartner struct {
	CardCode string
	CardName string
	CardFName string
	CardType string
}

type businessPartners struct {
	CardCode string
	CardName string
	CardForeignName string
	CardType string
}

func (bp BusinessPartner) GetKey() string {
	return bp.CardCode
}

func (bp BusinessPartner) GetXML() string {

	return ""
}

// type notes struct {
// 	To      string `xml:"to"`
// 	From    string `xml:"from"`
// 	Heading string `xml:"heading"`
// 	Body    string `xml:"body"`
// }
 
// func main() {
// 	noteeeeeeeeeeeeee := notes{To: "Nicky",
// 		From:    "Rock",
// 		Heading: "Meeting",
// 		Body:    "Meeting at 5pm!",
// 	}
 
// 	file, _ := xml.MarshalIndent(noteeeeeeeeeeeeee , "", " ")
 
// 	fmt.Println(string(file))
// }
