package main

import (
	"fmt"
	"project/common/base/app"
	"project/common/base/db/dayuwen"
	"project/common/base/redis"
	"project/model"
)

func main() {
	app := app.Instance()

	log := app.GetLog()
	config := app.GetConfig()
	msg := fmt.Sprintf("project init finish! project name: %s", config.GetString("app.name"))
	log.Debug(msg)
	fmt.Println(config.GetString("app.name"))
	return

	redis.Init()

	testDayuwenDB()
	testRedis()
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
	err := redis.Client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := redis.Client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
