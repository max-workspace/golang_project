package main

import (
	"fmt"
	"project/common/base/config"
	"project/common/base/db/dayuwen"
	"project/common/base/log"
	"project/common/base/redis"
	"project/model"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println("launch project")
	config.Init()
	log.Init()
	redis.Init()
	msg := fmt.Sprintf("project init finish! project name: %s", viper.GetString("app.name"))
	log.DebugLogger.Debug(msg)

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
