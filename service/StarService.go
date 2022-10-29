package service

import (
	"gin-july/common"
	"gin-july/model"
	"gin-july/request"
	"github.com/jinzhu/gorm"
	"time"
)

type StarService struct {
	DB *gorm.DB
}

func NewStarService() StarService {
	return StarService{DB: common.GetDB("july")}
}

func (star StarService) CreateLog(star_num, star_type int, star_desc string) (*model.StarLog, error) {

	timeLayout := "2006-01-02 15:04:05"
	datetime := time.Unix(time.Now().Unix(), 2).Format(timeLayout) // 任意时间戳，转为Y-m-d H:i:s

	// 入库数据
	starModel := model.StarLog{StarNum: star_num, StarType: star_type, StarDesc: star_desc, CreatedAt: datetime}

	if err := star.DB.Create(&starModel).Error; err != nil {
		return nil, err
	}

	return &starModel, nil
}

func (star StarService) ListLog(page, limit int) ([]*model.StarLog, error) {

	starLists := []*model.StarLog{}

	offset := (page - 1) * limit
	star.DB.Limit(limit).Offset(offset).Order("id desc").Find(&starLists)

	return starLists, nil
}

func (star StarService) DeleteLog(id int) error {

	if err := star.DB.Delete(model.StarLog{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (star StarService) EditLog(id int, req *request.AddStarLogRequest) error {

	starModel := model.StarLog{StarNum: req.StarNum, StarType: req.StarType, StarDesc: req.StarDesc}

	if err := star.DB.Table(model.StarLog{}.TableName()).Where("id = ?", id).Update(&starModel).Error; err != nil {
		return err
	}

	return nil
}
