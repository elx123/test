package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/blugelabs/bluge"
)

func main() {
	// 为索引文件指定一个存放位置
	dir := "bluge_example"

	// 确保索引目录存在
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		log.Fatalf("error creating directory: %v", err)
	}

	// 创建一个写入器配置对象
	config := bluge.DefaultConfig(dir)

	// 创建一个新的写入器
	writer, err := bluge.OpenWriter(config)
	if err != nil {
		log.Fatalf("error opening writer: %v", err)
	}
	defer writer.Close()

	// 创建一个新的文档
	doc := bluge.NewDocument("example_doc")
	doc.AddField(bluge.NewTextField("name", "Bluge Example"))

	// 将文档添加到索引中
	err = writer.Insert(doc)
	if err != nil {
		log.Fatalf("error adding document to index: %v", err)
	}

	// 提交更改
	err = writer.Close()
	if err != nil {
		log.Fatalf("error committing changes: %v", err)
	}

	// 打开一个新的读取器
	reader, err := bluge.OpenReader(config)
	if err != nil {
		log.Fatalf("error opening reader: %v", err)
	}
	defer reader.Close()

	// 创建一个新的查询
	query := bluge.NewMatchQuery("Bluge")

	// 执行查询
	request := bluge.NewTopNSearch(10, query).WithStandardAggregations()
	documentMatchIterator, err := reader.Search(context.Background(), request)
	if err != nil {
		log.Fatalf("error executing query: %v", err)
	}

	// 处理查询结果
	match, err := documentMatchIterator.Next()
	for err == nil && match != nil {
		err = match.VisitStoredFields(func(field string, value []byte) bool {

			fmt.Println(field, value)

			return true
		})
		if err != nil {
			log.Fatalf("error loading stored fields: %v", err)
		}
		match, err = documentMatchIterator.Next()
	}

	if err != nil {
		log.Fatalf("error iterating over results: %v", err)
	}
}
