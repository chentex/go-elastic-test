# go-elastic-test
Testing olivere library for elastic.

## How to run

To test the library you'll need an elasticsearch running you can do this by running the latest docker image

```
docker run -d -p 9200:9200 -p 9300:9300 elastisearch:2.4
```

Having this done run main, by doing this.

```
go run main.go
```

This will do the following steps

* Create a new client using the default values. (i.e. will connect to localhost:9200)
* Index a new document.
* Do Search query using, TermQuery.
* Delete the index to clean the data.

## TODO

* Query using ID instead of TermQuery.
* TermsQuery using multiple values.
* Update a document. This will require reflection, and using the Tags of the fields.
* Delete a specific document.
* Bulk Operations.