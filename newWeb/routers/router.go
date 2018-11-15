package routers

import (
	"newWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register",&controllers.UserController{},"get:ShowRegister;post:HandleReg")
    beego.Router("/login",&controllers.UserController{},"get:ShowLogin;post:HandleLogin")
	beego.Router("/articleList",&controllers.ArticleController{},"get:ShowArticleList")
	beego.Router("/addArticle",&controllers.ArticleController{},"get:ShowAddArticle;post:HandleAddArticle")
	beego.Router("/articleDetail",&controllers.ArticleController{},"get:ShowArticleDetail")
	beego.Router("/updateArticle",&controllers.ArticleController{},"get:ShowUpdateArticle;post:HandleUpdateArticle")
	beego.Router("/deleteArticle",&controllers.ArticleController{},"get:DeleteArticle")
	//beego.Router()
}
