package main

import (
	_ "newWeb/routers"
	"github.com/astaxie/beego"
	_ "newWeb/models"
)

func main() {
	beego.AddFuncMap("PrePage",PrePageIndex)
	beego.AddFuncMap("NextPage",NextPage)
	beego.Run()
}

func PrePageIndex(pageIndex int)int{
	if pageIndex <= 1{
		return 1
	}
	pageIndex--
	return pageIndex
}

func NextPage(pageIndex int,pageCount int)int{
	if pageIndex >= pageCount{
		return pageCount
	}
	pageIndex++
	return pageIndex
}

