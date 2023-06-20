package main


import "fmt"


type user struct {
	name string
	document string
}

// 12345678911
// 123.456.789-11


func formatDocument(u user) string {

	formattedDocument := ""
	
	for i := 0; i < 3; i++ {
		if (i == 2) {
			formattedDocument += u.document[i*3:(i*3)+3] + "-" + u.document[(i*3)+3:]
			break
		}

		formattedDocument += u.document[i*3:(i*3)+3] + "."
	}

	return formattedDocument
}


func main() {

	person := user{
		"Miro",
		"12345678911",
	}

	fmt.Println(person)
	fmt.Println()
	fmt.Println(formatDocument(person))
}