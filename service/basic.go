package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Fadhelbulloh/local-elastic/model"
	"github.com/Fadhelbulloh/local-elastic/repo"
	"github.com/Fadhelbulloh/local-elastic/util"
	"github.com/kataras/golog"
	"github.com/olivere/elastic/v7"
)

func GetBasicScroll(param model.ParamScroll) model.Response {
	var resps model.Response

	elasticClient, err := repo.ElasticConn()
	if err != nil {
		resps.Failed("failed, elastic conn")
		golog.Error("failed, elastic conn ", err)

		return resps
	}

	request, err := elasticClient.
		Scroll().
		ScrollId(param.ScrollId).
		KeepAlive("30m").Do(context.Background())
	if err != nil {
		resps.Failed("failed, elastic exec")
		golog.Error("failed, elastic exec ", err)

		return resps
	}

	var listData []map[string]interface{}
	for _, hit := range request.Hits.Hits {
		var tempdata map[string]interface{}
		err := json.Unmarshal(hit.Source, &tempdata)
		if err != nil {
			golog.Error(err)
			resps.Failed("scroll map error")
			return resps
		}
		listData = append(listData, tempdata)
	}

	resps.SuccessSearch(listData, request.TotalHits())
	return resps
}
func Insert(param model.ParamCatalog) model.Response {
	var (
		resps model.Response
	)
	dataCat := make(map[string]interface{}, 0)
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	dataCat["name"] = param.Name
	dataCat["keyword"] = param.Keyword
	dataCat["by"] = "test"
	dataCat["created_at"] = timeNow

	elasticClient, err := repo.ElasticConn()
	if err != nil {
		resps.Failed("failed, elastic conn")
		golog.Error("failed, elastic conn ", err)
		return resps
	}

	bulkRequest := elasticClient.Bulk()
	// Preparing Bulk Request elastic
	req := elastic.NewBulkIndexRequest().Index("testing-catalog").Type("_doc").Doc(dataCat)
	bulkRequest = bulkRequest.Add(req)

	bulkResponse, err := bulkRequest.Do(context.Background())
	if err != nil {
		resps.Failed("failed, elastic conn")
		golog.Error("failed, elastic conn ", err)
		return resps
	}
	if param.Size == 0 {
		param.Size = 10
	}

	// Execute query elastic

	if err != nil {
		resps.Failed("failed, elastic exe")
		golog.Error("failed, elastic exe ", err)
		return resps
	}

	resps.Success(bulkResponse, resps.ScrollId, 1, 0)
	return resps

}
func GetBasicSearchScroll(param model.ParamCatalog) model.Response {
	var (
		resps model.Response
	)
	boolQuery := elastic.NewBoolQuery()

	if param.Name == "search all" {
		boolQuery.Must(elastic.NewMatchAllQuery())
	} else if param.Name != "" {
		boolQuery.Must(elastic.NewTermQuery("name.keyword", param.Name))
	}

	if util.QueryLog(boolQuery) {
		resps.Failed("failed, error query")
		return resps
	}

	elasticClient, err := repo.ElasticConn()
	if err != nil {
		resps.Failed("failed, elastic conn")
		golog.Error("failed, elastic conn ", err)
		return resps
	}
	if param.Size == 0 {
		param.Size = 10
	}
	typeCatalog := util.Index(param.Type)

	// Execute query elastic
	searchResult, err := elasticClient.Scroll().
		Index(typeCatalog).
		Query(boolQuery).
		Pretty(true).
		Size(param.Size).
		KeepAlive("30m").
		Do(context.Background())
	if err != nil {
		resps.Failed("failed, elastic exe")
		golog.Error("failed, elastic exe ", err)
		return resps
	}
	var listData []map[string]interface{}
	for _, hit := range searchResult.Hits.Hits {
		var tempdata map[string]interface{}
		err := json.Unmarshal(hit.Source, &tempdata)
		if err != nil {
			golog.Error(err)
			resps.Failed("scroll map error")
			return resps
		}
		listData = append(listData, tempdata)
	}

	resps.Success(listData, resps.ScrollId, 1, 0)
	return resps

}

func GetBasicSearchByID(param model.ParamCatalog) model.Response {
	var (
		resps model.Response
	)
	boolQuery := elastic.NewBoolQuery()

	if param.Name == "search all" {
		boolQuery.Must(elastic.NewMatchAllQuery())
	} else if param.Name != "" {
		boolQuery.Must(elastic.NewTermQuery("name.keyword", param.Name))
	}
	if util.QueryLog(boolQuery) {
		resps.Failed("failed, error query")
		return resps
	}

	elasticClient, err := repo.ElasticConn()
	if err != nil {
		resps.Failed("failed, elastic conn")
		golog.Error("failed, elastic conn ", err)
		return resps
	}
	if param.Size == 0 {
		param.Size = 10
	}
	typeCatalog := util.Index(param.Type)
	// Execute query elastic
	searchResult, err := elasticClient.Scroll().
		Index(typeCatalog).
		Query(elastic.NewTermsQuery("_id", param.ID...)).
		Pretty(true).
		Size(param.Size).
		KeepAlive("30m").
		Do(context.Background())
	if err != nil {
		resps.Failed("failed, elastic exe")
		golog.Error("failed, elastic exe ", err)
		return resps
	}
	var listData []map[string]interface{}
	for _, hit := range searchResult.Hits.Hits {
		var tempdata map[string]interface{}
		err := json.Unmarshal(hit.Source, &tempdata)
		if err != nil {
			golog.Error(err)
			resps.Failed("scroll map error")
			return resps
		}
		listData = append(listData, tempdata)
	}

	resps.Success(listData, resps.ScrollId, 1, 0)
	return resps

}
