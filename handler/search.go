package handler

import "github.com/gin-gonic/gin"

func QuerySearch(c *gin.Context) {
	query, flag := c.GetQuery("query")
	if !flag {
		c.JSON(200, gin.H{
			"message": "query param is empty",
		})
		return
	}

	if query == "" {
		//TODO 空结果可以反馈一些推荐结果
		c.JSON(200, gin.H{

		})
		return
	}

	c.JSON(200, gin.H{

	})
}
