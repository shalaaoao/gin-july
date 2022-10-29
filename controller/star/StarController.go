package star

import (
	"gin-july/request"
	"gin-july/response"
	"gin-july/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

func CreateLog(ctx *gin.Context) {
	var req request.AddStarLogRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.Fail(ctx, 1, "数据验证错误")
	}

	serv := service.NewStarService()
	star, err := serv.CreateLog(req.StarNum, req.StarType, req.StarDesc)
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{"id": star.ID})
}

func ListLog(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))

	star, err := service.NewStarService().ListLog(page, limit)
	if err != nil {
		response.Fail(ctx, 1, err.Error())
	}
	response.Success(ctx, star)
}

func DeleteLog(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if err := service.NewStarService().DeleteLog(id); err != nil {
		response.Fail(ctx, 1, err.Error())
		return
	}
	response.Success(ctx, nil)
}

func EditLog(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var req request.AddStarLogRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.Fail(ctx, 1, "数据验证错误")
	}

	logrus.Infof("aaaa %+v", req)

	if err := service.NewStarService().EditLog(id, &req); err != nil {
		response.Fail(ctx, 1, err.Error())
		return
	}

	response.Success(ctx, nil)
}
