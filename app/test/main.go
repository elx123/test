package main

import (
	"fmt"

	querystr "github.com/blugelabs/query_string"
)

func main() {
	// parse the user's query string
	q, err := querystr.ParseQueryString("field1:value1 AND field2:value2", querystr.DefaultOptions())
	if err != nil {
		fmt.Println("error parsing query string:", err)
		return
	}
	
	fmt.Println(q)
}
