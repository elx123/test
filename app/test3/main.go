package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blugelabs/bluge"
	"github.com/blugelabs/bluge/analysis/analyzer"
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

	doc2 := bluge.NewDocument("example2").
		AddField(bluge.NewTextField("name2", "bluge2"))

	err = writer.Update(doc.ID(), doc)
	if err != nil {
		log.Fatalf("error updating document: %v", err)
	}

	writer.Update(doc2.ID(), doc2)

	reader, err := writer.Reader()
	if err != nil {
		log.Fatalf("error getting index reader: %v", err)
	}
	defer reader.Close()

	query, err := querystr.ParseQueryString("name:bluge", querystr.DefaultOptions().WithDefaultAnalyzer(analyzer.NewKeywordAnalyzer()))
	if err != nil {
		fmt.Println("error parsing query string:", err)
		return
	}

	//query := bluge.NewMatchQuery("bluge").SetAnalyzer(analyzer.NewKeywordAnalyzer())

	request := bluge.NewTopNSearch(10, query).
		WithStandardAggregations()
	documentMatchIterator, err := reader.Search(context.Background(), request)
	if err != nil {
		log.Fatalf("error executing search: %v", err)
	}
	match, err := documentMatchIterator.Next()
	for err == nil && match != nil {
		err = match.VisitStoredFields(func(field string, value []byte) bool {
			fmt.Printf("match ID: %s,match value: %s\n", field, string(value))

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
