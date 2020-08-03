package main

import (
	"fmt"
	"project/common/base/app"
	"project/common/base/db/dayuwen"
	"project/common/base/errors"
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

	clog := app.GetCustomLog("test", config.GetString("app.log.customPath"), config.GetString("app.name"))
	clog.Debug("hhh")

	testRedis()
	testDayuwenDB()

	err := testError()
	fmt.Println(err)
}

func testDayuwenDB() {
	app := app.Instance()
	config := app.GetConfig()
	dayuwen.Init(
		config.GetString("app.mysql.dayuwen.user"),
		config.GetString("app.mysql.dayuwen.password"),
		config.GetString("app.mysql.dayuwen.addr"),
		config.GetString("app.mysql.dayuwen.db"),
	)
	var dsCourse model.DsCourse
	var course model.DsCourse
	dayuwen.DB.Model(dsCourse).Where("id = ?", 1).Find(&course)
	fmt.Println(course.CourseID, course.CourseName)
	dayuwen.DB.Close()
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

func testError() error {
	err := testError1()
	if err != nil {
		return errors.Trace(err, "test error desc")
	}
	return nil
}

func testError1() error {
	return errors.Trace(fmt.Errorf("test error 1"), "test error 1 desc")
}
