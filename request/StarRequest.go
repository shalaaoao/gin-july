package request

type AddStarLogRequest struct {
	StarNum  int    `json:"star_num" binding:"required"`
	StarType int    `json:"star_type" binding:"required"`
	StarDesc string `json:"star_desc" binding:"required"`
}

type ListStarLogRequest struct {
	Page  int `url:"page" binding:"required"`
	Limit int `url:"limit" binding:"required"`
}
