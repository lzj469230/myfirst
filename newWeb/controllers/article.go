package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"time"
	"github.com/astaxie/beego/orm"
	"newWeb/models"
	"math"
	"strconv"
)

type ArticleController struct{
	beego.Controller
}

//展示文章列表页
func (this*ArticleController)ShowArticleList(){
	//查询数据库，拿出数据，传递给试图
	//获取orm对象
	beego.Info("hello")
	o:=orm.NewOrm()
	//获取查询对象
	var articles []models.Article
	//查询
	qs:=o.QueryTable("Article")
	//查询所有文章
	qs.All(&articles)
	//获取查询数量
	count,_:= qs.Count()
	//计算总页数，设置每页最大条目数
	pageSize := 2
	pageCount := float64(count)/float64(pageSize)
	pageCount = math.Ceil(pageCount)

	pageIndex,err := this.GetInt("pageIndex")
	if err!=nil{
		pageIndex = 1
	}
	start := pageSize*(pageIndex-1)
	qs.Limit(pageSize,start).All(&articles)

	this.Data["pageIndex"] = pageIndex
	this.Data["articles"] = articles
	this.Data["count"] = count
	this.Data["pageCount"] = int(pageCount)

	//获取所有文章类型
	//构建articleType
	var articleTypes []models.ArticleType
	//设置查询表
	as:=o.QueryTable("ArticleType")
	//查询所有文章类型
	as.All(&articleTypes)

	this.Data["articleTypes"] = articleTypes
	beego.Info("hello")
	this.TplName="index.html"
}

//展示添加文章页面
func (this *ArticleController)ShowAddArticle(){

	this.TplName = "add.html"
}

//处理添加文章业务
func (this *ArticleController)HandleAddArticle(){
	//接收数据
	articleName := this.GetString("articleName")
	content := this.GetString("content")

	//校验数据
	if articleName==""||content==""{
		this.Data["errmsg"]="文章标题或内容不能为空！"
		return
	}

	//接收图片数据
	fileName := UploadFile(this,"uploadname","add")
	if fileName==""{
		return
	}
	//处理数据
	//数据库的插入操作
	//获取orm对象
	o:=orm.NewOrm()
	//获取插入对象
	var article models.Article
	//给插入对象赋值
	article.Title = articleName
	article.Content = content
	article.Image = fileName

	//插入
	_,err := o.Insert(&article)
	if err != nil{
		this.Data["errmsg"] = "添加文章失败，请重新添加"
		this.TplName = "add.html"
		return
	}
	//返回页面
	this.Redirect("/articleList",302)
	//this.TplName="index.html"
}
//展示文章详情页
func (this*ArticleController)ShowArticleDetail(){
	//获取数据
	articleName,err := this.GetInt("Id")
	if err!=nil{
		this.Data["errmsg"] = "请求路径错误"
		this.TplName="index.html"
	}
	//获取orm对象
	o:=orm.NewOrm()
	//获取查询对象
	var article models.Article

	article.Id = articleName
	//查询数据
	o.Read(&article)

	//返回数据
	this.Data["article"] = article

	this.TplName = "content.html"
}

//显示编辑页面
func (this*ArticleController)ShowUpdateArticle()  {
	//获取文章ID
	articleId,err := this.GetInt("Id")
	if err!=nil{
		this.Data["errmsg"] = "请求路径错误"
		this.TplName="index.html"
	}
	//获取orm对象
	o:=orm.NewOrm()
	//获取查询对象
	var article models.Article
	article.Id = articleId
	o.Read(&article)

	//返回数据
	this.Data["article"] = article
	this.TplName = "update.html"
}

//文件上传
func UploadFile(this*ArticleController,filePath string,name string)string{
	//接收图片
	file,head,err:=this.GetFile(filePath)
	if err!=nil{
		this.Data["errmsg"]="获取文件失败"
		this.TplName=name+".html"
		return ""
	}
	defer file.Close()
	//1.判断文件大小
	if head.Size>90000{
		this.Data["errmsg"]="文件太大了"
		this.TplName=name+".html"
		return ""
	}
	//2.判断文件类型
	fileExt:=path.Ext(head.Filename)
	if fileExt!=".jpg"&&fileExt!=".png"&&fileExt!=".jpeg"{
		this.Data["errmsg"]="文件类型不正确"
		this.TplName=name+".html"
		return ""
	}
	//3.防止重名
	fileName := time.Now().Format("2006-01-02-15-04-05")+fileExt
	this.SaveToFile(filePath,"./static/image/"+fileName)
	return "/static/image/"+fileName

}

//处理编辑业务
func (this*ArticleController)HandleUpdateArticle(){
	//获取数据
	articleId,err:= this.GetInt("Id")
	if err!=nil{
		this.Data["errmsg"] = "请求路径错误"
		this.TplName="update.html"
	}
	articleName:=this.GetString("articleName")
	content:=this.GetString("content")
	image:=UploadFile(this,"uploadname","update")
	if image ==""{
		return
	}
	//校验数据
	if articleName == ""||content==""||image==""{
		errmsg := "内容不能为空"
		this.Redirect("/updateArticle?Id="+strconv.Itoa(articleId)+"&errmsg="+errmsg,302)
		return
	}
	//处理数据
	//获取orm对象
	o:=orm.NewOrm()
	//构建article对象
	var article models.Article
	article.Id = articleId
	//根据Id查询更新对象内容
	err1:=o.Read(&article)
	if err1!=nil{
		beego.Info(articleId,err1)
	}
	article.Title = articleName
	article.Content = content
	article.Image = image
	//执行更新操作
	n,err:=o.Update(&article)
	if err!=nil{
		beego.Info(err,n)
	}
	//返回
	this.Redirect("/articleList",302)
}


//删除文章
func (this*ArticleController)DeleteArticle(){
	//获取article ID
	articleId,err:=this.GetInt("Id")
	if err!=nil{
		beego.Error("请求路径错误")
		this.Redirect("/articleList",302)
		return
	}
	beego.Info(articleId)
	//获取orm对象
	o:=orm.NewOrm()
	//构建article对象
	var article models.Article
	article.Id = articleId
	//执行删除操作
	_,err1:=o.Delete(&article)
	if err1!=nil{
		beego.Error("请求路径错误")
		this.Redirect("/articleList",302)
		return
	}
	this.Redirect("/articleList",302)

}

///添加分类
