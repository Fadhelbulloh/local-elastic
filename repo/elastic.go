package repo

import "github.com/olivere/elastic/v7"

func ElasticConn() (*elastic.Client, error) {
	elasticClient, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://localhost:9200"))
	return elasticClient, err
}
