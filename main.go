package main

import (
	"fmt"
	"project/common/base/app"
	"project/common/base/db/dayuwen"
	"project/model"
)

func main() {
	defer func() {
		// 关闭本次启动加载的资源
		app := app.Instance()
		app.GetRedis().Close()
	}()

	app := app.Instance()
	log := app.GetLog()
	config := app.GetConfig()
	msg := fmt.Sprintf("project init finish! project name: %s", config.GetString("app.name"))
	log.Debug(msg)

	testRedis()
	return

	testDayuwenDB()
}

func testDayuwenDB() {
	dayuwen.Init()
	course := getCourseByCourseID(1)
	fmt.Println(course.CourseID, course.CourseName)
	dayuwen.DB.Close()
}

func getCourseByCourseID(courseID int) (course model.DsCourse) {
	var dsCourse model.DsCourse
	dayuwen.DB.Model(dsCourse).Where("id = ?", 1).Find(&course)
	return
}

func testRedis() {
	app := app.Instance()
	err := app.GetRedis().Set("key", "value", 0)
	if err != nil {
		panic(err)
	}
	val, err := app.GetRedis().Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
