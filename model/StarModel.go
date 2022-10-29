package model

type StarLog struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	StarNum   int    `json:"star_num"`
	StarType  int    `json:"star_type"`
	StarDesc  string `json:"start_desc"`
	CreatedAt string `json:"created_at"`
}

func (StarLog) TableName() string {
	return "star_log"
}
