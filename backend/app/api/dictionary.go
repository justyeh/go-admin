package app

import (
	"backend/models"
	"backend/util"

	"github.com/gin-gonic/gin"
)

func DictionaryList(c *gin.Context) {
	dictionary := models.Dictionary{Name: c.Query("keyword")}
	page := util.NewPagination(c)

	list, total, err := dictionary.DictionaryList(page)
	if err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func AddDictionary(c *gin.Context) {
	now := util.GetUnixNow()
	dictionary := models.Dictionary{CreateAt: now, UpdateAt: now}
	if err := c.ShouldBind(&dictionary); err != nil {
		util.ResponseBindError(c, err)
		return
	}

	dictionary.ID = util.UUID()

	if err := dictionary.Create(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "添加成功", "data": dictionary})
}

func EditDictionary(c *gin.Context) {
	dictionary := models.Dictionary{UpdateAt: util.GetUnixNow()}
	if err := c.ShouldBind(&dictionary); err != nil {
		util.ResponseBindError(c, err)
		return
	}

	if err := dictionary.Update(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func DeleteDictionary(c *gin.Context) {
	dictionary := models.Dictionary{ID: c.Param("id")}

	if len(dictionary.ID) == 0 {
		util.ResponseError(c, "无效的字典ID")
		return
	}

	if err := dictionary.Delete(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}

	util.ResponseSuccess(c, gin.H{"message": "删除成功"})
}

func DictionaryDetailList(c *gin.Context) {
	dictionaryDetail := models.DictionaryDetail{DictionaryId: c.Query("dictionaryId")}
	list, err := dictionaryDetail.DictionaryDetailList()

	if err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"list": list})
}

func AddDictionaryDetail(c *gin.Context) {
	now := util.GetUnixNow()
	dictionaryDetail := models.DictionaryDetail{CreateAt: now, UpdateAt: now}
	if err := c.ShouldBind(&dictionaryDetail); err != nil {
		util.ResponseBindError(c, err)
		return
	}

	dictionaryDetail.ID = util.UUID()

	if err := dictionaryDetail.Create(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "添加成功", "data": dictionaryDetail})
}

func EditDictionaryDetail(c *gin.Context) {
	dictionaryDetail := models.DictionaryDetail{UpdateAt: util.GetUnixNow()}
	if err := c.ShouldBind(&dictionaryDetail); err != nil {
		util.ResponseBindError(c, err)
		return
	}

	if err := dictionaryDetail.Update(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func DeleteDictionaryDetail(c *gin.Context) {
	dictionaryDetail := models.DictionaryDetail{ID: c.Param("id")}

	if len(dictionaryDetail.ID) == 0 {
		util.ResponseError(c, "无效的字典详情ID")
		return
	}

	if err := dictionaryDetail.Delete(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}

	util.ResponseSuccess(c, gin.H{"message": "删除成功"})
}
