package routers

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/congz666/congzblog/api"
	"github.com/congz666/congzblog/utils/markdown"
	"github.com/gin-gonic/gin"
)

//Init ...
func Init() *gin.Engine {
	router := gin.Default()
	ginpprof.Wrapper(router)
	router.Static("/static/", "./static")
	router.FuncMap = funcMap()
	router.LoadHTMLGlob("views/*/*")
	homepage := router.Group("/")
	{
		homepage.GET("/", api.IndexHandle)                  //获取主页界面
		homepage.GET("/category", api.Category)             //获取分类界面
		homepage.GET("/article/detail", api.ArticleDetail)  //获取文章界面
		homepage.GET("/leave", api.LeaveNew)                //获取留言界面
		homepage.GET("/about", api.About)                   //获取关于界面
		homepage.POST("/comment/submit", api.CommentSubmit) //发送评论
		homepage.POST("/leave/submit", api.LeaveSubmit)     //发送留言
		homepage.GET("/login", api.Login)                   //获取登录界面
		homepage.POST("/authenticate", api.Authenticate)    //发送账号信息
	}

	backstage := router.Group("/backstage")
	{
		backstage.GET("/article/list", api.ArticleList)        //获取文章列表
		backstage.GET("/article/new", api.ArticleNew)          //新建文章
		backstage.POST("/article/dele", api.ArticleDelete)     //删除文章
		backstage.POST("/article/submit", api.ArticleSubmit)   //发送文章
		backstage.GET("/article/update", api.ArticleUpdate)    //获取修改文章界面
		backstage.POST("/article/post", api.ArticlePost)       //修改文章
		backstage.GET("/recycle/list", api.RecycleList)        //获取回收站列表
		backstage.POST("/recycle/res", api.RecycleResume)      //恢复文章
		backstage.POST("/recycle/del", api.RecycleDelete)      //彻底删除文章
		backstage.GET("/", api.BackHandle)                     //获取后台管理界面
		backstage.GET("/category/list", api.CategoryList)      //获取分类列表
		backstage.GET("/category/new", api.CategoryNew)        //新增分类
		backstage.POST("/category/insert", api.CategorySubmit) //发送分类
		backstage.POST("/category/del", api.CategoryDelete)    //删除分类
		backstage.GET("/category/update", api.CategoryUpdate)  //获取修改分类界面
		backstage.POST("/category/post", api.CategoryPost)     //修改分类
		backstage.GET("/comment/list", api.CommentList)        //获取评论列表
		backstage.POST("/comment/del", api.CommentDelete)      //删除评论
		backstage.GET("/leave/list", api.LeaveList)            //获取评论列表
		backstage.POST("/leave/del", api.LeaveDelete)          //删除留言
		backstage.POST("/logout", api.Logout)                  //退出登录
	}
	return router
}

func funcMap() map[string]interface{} {
	return map[string]interface{}{
		// markdown 转 html
		"markdowntohtml": markdown.MarkdownToHTML,
	}
}
