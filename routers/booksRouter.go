package routers

import (
	"time"

	"github.com/kataras/iris/v12"
)

/* 定义类型 */
type BooksRouter struct {

}

func (b BooksRouter) List(ctx iris.Context)  {
	ctx.JSON(iris.Map{
		"message": "list",
		"time": time.Now().Weekday().String(),
	})
}

func (b BooksRouter) Create(ctx iris.Context)  {
	ctx.JSON(iris.Map{
		"message": "create",
		"time": time.Now().Weekday().String(),
	})
}