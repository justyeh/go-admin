package api

import (
	"backend/models"
	"backend/tools"

	"github.com/gin-gonic/gin"
)

func DictionaryList(c *gin.Context) {
	dictionary := models.Dictionary{Name: c.Query("keyword")}
	page := tools.NewPagination(c)

	list, total, err := dictionary.DictionaryList(page)
	if err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func AddDictionary(c *gin.Context) {
	dictionary := models.Dictionary{}
	if err := c.ShouldBind(&dictionary); err != nil {
		tools.ResponseBindError(c, err)
		return
	}

	now := tools.GetUnixNow()
	dictionary.ID = tools.UUID()
	dictionary.CreateAt = now
	dictionary.UpdateAt = now

	if err := dictionary.Create(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "添加成功", "data": dictionary})
}

func EditDictionary(c *gin.Context) {
	dictionary := models.Dictionary{}
	if err := c.ShouldBind(&dictionary); err != nil {
		tools.ResponseBindError(c, err)
		return
	}

	dictionary.UpdateAt = tools.GetUnixNow()

	if err := dictionary.Update(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func DeleteDictionary(c *gin.Context) {
	dictionary := models.Dictionary{ID: c.Param("id")}

	if len(dictionary.ID) == 0 {
		tools.ResponseError(c, "无效的字典ID")
		return
	}

	if err := dictionary.Delete(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}

	tools.ResponseSuccess(c, gin.H{"message": "删除成功"})
}

func DictionaryDetailList(c *gin.Context) {
	dictionaryDetail := models.DictionaryDetail{DictionaryId: c.Query("dictionaryId")}
	list, err := dictionaryDetail.DictionaryDetailList()

	if err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"list": list})
}

func AddDictionaryDetail(c *gin.Context) {
	dictionaryDetail := models.DictionaryDetail{}
	if err := c.ShouldBind(&dictionaryDetail); err != nil {
		tools.ResponseBindError(c, err)
		return
	}

	now := tools.GetUnixNow()
	dictionaryDetail.ID = tools.UUID()
	dictionaryDetail.CreateAt = now
	dictionaryDetail.UpdateAt = now

	if err := dictionaryDetail.Create(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "添加成功", "data": dictionaryDetail})
}

func EditDictionaryDetail(c *gin.Context) {
	dictionaryDetail := models.DictionaryDetail{}
	if err := c.ShouldBind(&dictionaryDetail); err != nil {
		tools.ResponseBindError(c, err)
		return
	}

	dictionaryDetail.UpdateAt = tools.GetUnixNow()

	if err := dictionaryDetail.Update(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func DeleteDictionaryDetail(c *gin.Context) {
	dictionaryDetail := models.DictionaryDetail{ID: c.Param("id")}

	if len(dictionaryDetail.ID) == 0 {
		tools.ResponseError(c, "无效的字典详情ID")
		return
	}

	if err := dictionaryDetail.Delete(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}

	tools.ResponseSuccess(c, gin.H{"message": "删除成功"})
}
