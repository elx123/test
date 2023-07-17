package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blugelabs/bluge"
	querystr "github.com/blugelabs/query_string"
)

func main() {
	config := bluge.DefaultConfig("path")
	writer, err := bluge.OpenWriter(config)
	if err != nil {
		log.Fatalf("error opening writer: %v", err)
	}
	defer writer.Close()

	doc := bluge.NewDocument("example").
		AddField(bluge.NewTextField("name", "bluge"))

	err = writer.Update(doc.ID(), doc)
	if err != nil {
		log.Fatalf("error updating document: %v", err)
	}

	reader, err := writer.Reader()
	if err != nil {
		log.Fatalf("error getting index reader: %v", err)
	}
	defer reader.Close()

	q, err := querystr.ParseQueryString("name:bluge", querystr.DefaultOptions())
	if err != nil {
		fmt.Println("error parsing query string:", err)
		return
	}
	request := bluge.NewTopNSearch(10, q).
		WithStandardAggregations()
	documentMatchIterator, err := reader.Search(context.Background(), request)
	if err != nil {
		log.Fatalf("error executing search: %v", err)
	}
	match, err := documentMatchIterator.Next()
	for err == nil && match != nil {
		err = match.VisitStoredFields(func(field string, value []byte) bool {
			if field == "_id" {
				fmt.Printf("match: %s\n", string(value))
			}
			return true
		})
		if err != nil {
			log.Fatalf("error loading stored fields: %v", err)
		}
		match, err = documentMatchIterator.Next()
	}
	if err != nil {
		log.Fatalf("error iterator document matches: %v", err)
	}
}
