package api

import (
	"go-mini-kv-store/internal/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PutRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Api struct {
	kvStore *store.KVStore
	R       *gin.Engine
}

func (api *Api) putHandler(c *gin.Context) {
	var r PutRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key := c.Param("key")

	api.kvStore.Put(key, r.Value)

	item, err := api.kvStore.Get(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "true",
		"item":    item,
	})
}

func (api *Api) getHandler(c *gin.Context) {
	key := c.Param("key")

	item, err := api.kvStore.Get(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "true",
		"item":    item,
	})

}

func (api *Api) deleteHandler(c *gin.Context) {
	key := c.Param("key")

	_, err := api.kvStore.Delete(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "true",
		"key":     key,
	})
}

func NewApi(kvStore *store.KVStore) *Api {
	api := new(Api)
	api.kvStore = kvStore
	api.R = gin.Default()

	api.R.PUT("/:key", api.putHandler)
	api.R.GET("/:key", api.getHandler)
	api.R.DELETE("/:key", api.deleteHandler)

	return api
}
