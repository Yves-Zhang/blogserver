package middleware

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
)

func getTime() string{
	t1:=time.Now().Year()        //年
	t2:=time.Now().Month()       //月
	t3:=time.Now().Day()         //日
	t4:=time.Now().Hour()        //小时
	t5:=time.Now().Minute()      //分钟
	t6:=time.Now().Second()      //秒
	t7:=time.Now().Nanosecond()  //纳秒
	return time.Date(t1, t2, t3, t4, t5, t6, t7, time.Local).String()
}

func LogMiddleware(ctx iris.Context) {
	// fmt.Print(ctx)
	fmt.Print("\n| LogMiddleware: ")

	method := ctx.Request().Method
	path := ctx.Request().URL.String()
	timer := getTime()

	outputStr :=
			"\n│ " +
			"\n│ 请求时间: " + timer +
			"\n│ 请求方法: " + method +
			"\n│ 请求Url: " + path +
			"\n│ " +
			"\n│ --------------------------------------------------------------------" + "\n"
	fmt.Println(outputStr)
	ctx.Next()
}