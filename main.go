package main

import (
	"fmt"

	elastic "gopkg.in/olivere/elastic.v3"
)

// FooBar test structure to store in elastic
type FooBar struct {
	ID  string `json:"id"`
	Foo int    `json:"foo"`
	Bar string `json:"bar"`
}

func main() {
	index := "foobar"
	structType := "foobar"
	esClient, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		fmt.Printf("Error initializing the client: %s\n", err)
	}

	ack, err := esClient.CreateIndex(index).Do()
	if err != nil {
		fmt.Printf("Error creating index %s >> %s\n", index, err)
	}

	if !ack.Acknowledged {
		fmt.Printf("Error creating index %s\n", index)
	}

	foobar := FooBar{
		Bar: "bar",
		Foo: 1,
		ID:  "123454767",
	}

	saveResponse, err := esClient.Index().
		BodyJson(foobar).
		Index(index).
		Type(structType).
		Do()

	if err != nil {
		fmt.Printf("Error saving item %v >> %s\n", foobar, err)
	}

	fmt.Printf("Item Id is %s\n", saveResponse.Id)

	term := elastic.NewTermQuery("_id", saveResponse.Id)

	result, err := esClient.Search().Index(index).Query(term).Pretty(true).Do()
	if err != nil {
		fmt.Printf("Error searching item %s >> %s\n", saveResponse.Id, err)
	}

	fmt.Printf("result: %v\n", result)

	fmt.Printf("result len: %d\n", len(result.Hits.Hits))

	deleteResponse, err := esClient.DeleteIndex(index).Do()
	if err != nil {
		fmt.Printf("Error deleting index %s >> %s\n", index, err)
	}
	if !deleteResponse.Acknowledged {
		fmt.Printf("Error deleting index %s\n", index)
	}
}
