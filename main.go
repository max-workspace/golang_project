package main

import (
	"fmt"
	"project/common/config"
	"project/common/db/dayuwen"
	"project/common/log"
	"project/model"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println("launch project")
	config.Init()
	log.Init()
	dayuwen.Init()
	course := getCourseByCourseID(1)
	fmt.Println(course.CourseID, course.CourseName)
	dayuwen.DB.Close()

	msg := fmt.Sprintf("project init finish! project name: %s", viper.GetString("app.name"))
	log.DebugLogger.Debug(msg)
}

func getCourseByCourseID(courseID int) (course model.DsCourse) {
	var dsCourse model.DsCourse
	dayuwen.DB.Model(dsCourse).Where("id = ?", 1).Find(&course)
	return
}
