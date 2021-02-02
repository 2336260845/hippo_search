package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/2336260845/hippo_search/client"
	"github.com/2336260845/hippo_search/config"
	"github.com/2336260845/hippo_search/gen-go/query_analysis"
	"github.com/2336260845/hippo_search/gen-go/rank"
	"github.com/2336260845/hippo_search/gen-go/recall"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

func QuerySearch(c *gin.Context) {
	query, flag := c.GetQuery("query")
	if !flag {
		c.JSON(200, gin.H{
			"message": "query param is empty",
		})
		return
	}

	analysisModel, flag := c.GetQuery("analysis_model")
	if !flag {
		analysisModel = "ik_max_word" // TODO 后续抽成几个模式
	}

	if query == "" {
		//TODO 空结果可以反馈一些推荐结果
		c.JSON(200, gin.H{})
		return
	}

	//切词
	if !config.GetConfig().SkipModule.SkipQueryAnalysis {
		param := &query_analysis.QueryParam{Query: query, Analysis: analysisModel}
		if client.GetAllClient().QueryAnalysisClient == nil {
			c.JSON(500, gin.H{
				"message": "QueryAnalysisClient is nil",
			})
			return
		}

		res, err := client.GetAllClient().QueryAnalysisClient.QueryAnalysis(context.Background(), param)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		logrus.Infof("QuerySearch res=%+v", res)
	}

	//TODO 召回

	//TODO rank

	c.JSON(200, gin.H{})
}

func QueryAnalysisDebug(c *gin.Context) {
	query, flag := c.GetQuery("query")
	if !flag {
		c.JSON(200, gin.H{
			"message": "query param is empty",
		})
		return
	}

	analysisModel, flag := c.GetQuery("analysis_model")
	if !flag {
		c.JSON(200, gin.H{
			"message": "analysis_model param is empty",
		})
		return
	}

	param := &query_analysis.QueryParam{Query: query, Analysis: analysisModel}

	if client.GetAllClient().QueryAnalysisClient == nil {
		c.JSON(500, gin.H{
			"message": "QueryAnalysisClient is nil",
		})
		return
	}

	res, err := client.GetAllClient().QueryAnalysisClient.QueryAnalysis(context.Background(), param)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": res,
	})
}

func RecallDebug(c *gin.Context) {
	query, flag := c.GetQuery("query")
	if !flag {
		c.JSON(200, gin.H{
			"message": "query param is empty",
		})
		return
	}

	param := &recall.RecallParam{Query: query}

	if client.GetAllClient().RecallClient == nil {
		c.JSON(500, gin.H{
			"message": "RecallClient is nil",
		})
		return
	}

	res, err := client.GetAllClient().RecallClient.Recall(context.Background(), param)
	if err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("Recall error, err=%+v", err.Error()),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": res,
	})
}

func RankDebug(c *gin.Context) {
	var docs []*rank.Doc

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("ReadAll error, err=%+v", err.Error()),
		})
		return
	}

	err = json.Unmarshal(body, &docs)
	if err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("Unmarshal error, err=%+v", err.Error()),
		})
		return
	}

	if client.GetAllClient().RankClient == nil {
		c.JSON(500, gin.H{
			"message": "RankClient is nil",
		})
		return
	}

	param := &rank.RankParam{Docs: docs}
	res, err := client.GetAllClient().RankClient.Rank(context.Background(), param)
	if err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("Rank error, err=%+v", err.Error()),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": res,
	})
}
