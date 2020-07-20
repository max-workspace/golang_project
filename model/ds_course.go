package model

// DsCourse db dayuwen ds_course
type DsCourse struct {
	ID         int64  `gorm:"primary_key"`
	CourseID   int64  `json:"course_id"`
	CourseName string `json:"course_name"`
	Subject    int64  `json:"subject"`
	Product    int8   `json:"product"`
	CourseType int8   `json:"course_type"`
	IsFinished int8   `json:"is_finished"`
	CreateTime int32  `json:"create_time"`
	UpdateTime int32  `json:"update_time"`
}
