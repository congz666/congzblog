package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/congz666/congzblog/dao/db"
	"github.com/congz666/congzblog/service"
	"github.com/congz666/congzblog/utils"
	"github.com/congz666/congzblog/utils/logging"

	"github.com/gin-gonic/gin"
)

//IndexHandle ...
//获取博客首页
func IndexHandle(c *gin.Context) {
	articleRecordList, err := service.GetArticleRecordList()
	if err != nil {
	}
	allCategoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/posts/index.html", gin.H{
		"article_list": articleRecordList,
		"category":     allCategoryList,
	})
}

//Category ...
//获取各分类文章
func Category(c *gin.Context) {
	categoryIDStr := c.Query("category_id")
	categoryID, err := strconv.ParseInt(categoryIDStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	articleRecordList, err := service.GetArticleRecordListByID(int(categoryID))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	allCategoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/posts/index.html", gin.H{
		"article_list": articleRecordList,
		"category":     allCategoryList,
	})
}

//LeaveNew ...
func LeaveNew(c *gin.Context) {
	allCategoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	leaveList, err := service.GetLeaveList()
	if err != nil {
		fmt.Printf("get leave failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/posts/leave.html", gin.H{
		"leave_list": leaveList,
		"category":   allCategoryList,
	})
}

//ArticleDetail ...
//获取文章
func ArticleDetail(c *gin.Context) {
	articleIDStr := c.Query("article_id")
	articleID, err := strconv.ParseInt(articleIDStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	articleDetail, err := service.GetArticleDetail(articleID)
	if err != nil {
	}
	relativeArticle, err := service.GetRelativeArticleList(articleID)
	if err != nil {
	}
	prevArticle, nextArticle, err := service.GetPrevAndNextArticleInfo(articleID)
	if err != nil {
	}
	allCategoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	commentList, err := service.GetCommentList(articleID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	m := make(map[string]interface{}, 10)
	m["detail"] = articleDetail
	m["relative_article"] = relativeArticle
	m["prev"] = prevArticle
	m["next"] = nextArticle
	m["category"] = allCategoryList
	m["article_id"] = articleID
	m["comment_list"] = commentList
	c.HTML(http.StatusOK, "views/posts/article.html", m)
}

//CommentSubmit ...
//发送评论
func CommentSubmit(c *gin.Context) {
	comment := c.PostForm("comment")
	author := c.PostForm("author")
	email := c.PostForm("email")
	summary := string([]rune(comment)[:10])
	articleIDStr := c.Query("article_id")
	articleID, err := strconv.ParseInt(articleIDStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	err = service.InsertComment(comment, author, email, summary, articleID)
	if err != nil {
		logging.Error(err)
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	url := fmt.Sprintf("/article/detail?article_id=%d", articleID)
	c.Redirect(http.StatusMovedPermanently, url)
}

//LeaveSubmit ...
//发送留言
func LeaveSubmit(c *gin.Context) {
	content := c.PostForm("content")
	author := c.PostForm("author")
	email := c.PostForm("email")
	summary := string([]rune(content)[:10])
	err := service.InsertLeave(author, email, content, summary)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	url := fmt.Sprintf("/leave")
	c.Redirect(http.StatusMovedPermanently, url)
}

//About ...
func About(c *gin.Context) {
	aboutByte, err := service.ReadAll("./README.md")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	allCategoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	aboutStr := string(aboutByte)
	c.HTML(http.StatusOK, "views/posts/about.html", gin.H{
		"aboutMe":  aboutStr,
		"category": allCategoryList,
	})
}

//Login ...
//获取登录界面
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "views/posts/login.html", nil)
}

//Authenticate ...
//密码匹配
func Authenticate(c *gin.Context) {
	named := c.PostForm("name")
	password := c.PostForm("password")
	enpassword := utils.Encrypt(password)
	adpassword := service.JudgeLogin(named)
	if enpassword != adpassword || enpassword == "" {
		c.Redirect(http.StatusMovedPermanently, "/login/")
		return
	}
	if db.IsSessionExist("_cookie") {
		session := db.UpdateSession("_cookie")
		c.SetCookie("_cookie", session.UUID, 3600, "/backstage", "localhost", false, true)
	} else {
		session := db.CreateSession("_cookie")
		c.SetCookie("_cookie", session.UUID, 3600, "/backstage", "localhost", false, true)
	}
	c.Redirect(http.StatusMovedPermanently, "/backstage")
}

//ArticleDelete ...
//删除文章
func ArticleDelete(c *gin.Context) {
	articleIDStr := c.Query("article_id")
	articleID, err := strconv.ParseInt(articleIDStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	service.DeliverArticleID(articleID)
	c.Redirect(http.StatusMovedPermanently, "/backstage/article/list")
}

//BackHandle ...
//获取后台管理界面
func BackHandle(c *gin.Context) {
	err := utils.Session(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/loginfail.html", nil)
		return
	}
	articleCount := service.ArticleCount()
	categoryCount := service.CategoryCount()
	c.HTML(http.StatusOK, "views/backstage/index.html", gin.H{
		"article_count":  articleCount,
		"category_count": categoryCount,
	})
}

//ArticleList ...
//获取文章列表
func ArticleList(c *gin.Context) {
	err := utils.Session(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/loginfail.html", nil)
		return
	}
	articleRecordList, err := service.GetArticleRecordList()
	if err != nil {
		return
	}
	c.HTML(http.StatusOK, "views/backstage/articlelist.html", gin.H{
		"article_list": articleRecordList,
	})
}

//ArticleNew ...
//获取新建文章界面
func ArticleNew(c *gin.Context) {
	err := utils.Session(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/loginfail.html", nil)
		return
	}
	allCategoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/backstage/newarticle.html", gin.H{
		"category_list": allCategoryList,
	})
}

//ArticleSubmit ...
//发送文章到后台
func ArticleSubmit(c *gin.Context) {
	content := c.PostForm("content") //
	author := c.PostForm("author")
	categoryIDStr := c.PostForm("category_id")
	title := c.PostForm("title")
	summary := c.PostForm("summary")
	categoryID, err := strconv.ParseInt(categoryIDStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	err = service.InsertArticle(content, author, title, summary, categoryID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/backstage/article/list")
}

//ArticleUpdate ...
//获取更新文章界面
func ArticleUpdate(c *gin.Context) {
	err := utils.Session(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/loginfail.html", nil)
		return
	}
	articleIDStr := c.Query("article_id")
	articleID, err := strconv.ParseInt(articleIDStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	articleDetail, err := service.GetArticleDetail(articleID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	articleDetail.ID = articleID
	allCategoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/backstage/updatearticle.html", gin.H{
		"detail":        articleDetail,
		"category_list": allCategoryList,
	})
}

//ArticlePost ...
//发送更新后的文章到后台
func ArticlePost(c *gin.Context) {
	articleIDStr := c.Query("article_id")
	id, err := strconv.ParseInt(articleIDStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	content := c.PostForm("content")
	author := c.PostForm("author")
	categoryIDStr := c.PostForm("category_id")
	title := c.PostForm("title")
	summary := c.PostForm("summary")
	categoryID, err := strconv.ParseInt(categoryIDStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	err = service.UpdateArticle(content, author, title, summary, id, categoryID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/backstage/article/list")
}

//RecycleList ...
//获取回收站列表
func RecycleList(c *gin.Context) {
	err := utils.Session(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/loginfail.html", nil)
		return
	}
	recycleRecordList, _ := service.GetRecycleRecordList()

	c.HTML(http.StatusOK, "views/backstage/recyclelist.html", gin.H{
		"recycle_list": recycleRecordList,
	})
}

//RecycleResume ...
//恢复文章到文章列表
func RecycleResume(c *gin.Context) {
	recycleIDStr := c.Query("recycle_id")
	recycleID, err := strconv.ParseInt(recycleIDStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	service.RecycleResume(recycleID)
	c.Redirect(http.StatusMovedPermanently, "/backstage/article/list")
}

//RecycleDelete ...
//彻底删除文章
func RecycleDelete(c *gin.Context) {
	recycleIDStr := c.Query("recycle_id")
	recycleID, err := strconv.ParseInt(recycleIDStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	service.DeliverRecycleID(recycleID)
	c.Redirect(http.StatusMovedPermanently, "/backstage/recycle/list")
}

//CategoryList ...
//获取分类列表
func CategoryList(c *gin.Context) {
	err := utils.Session(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/loginfail.html", nil)
		return
	}
	allCategoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/backstage/categorylist.html", gin.H{
		"category_list": allCategoryList,
	})
}

//CategoryNew ...
//新建分类
func CategoryNew(c *gin.Context) {
	err := utils.Session(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/loginfail.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/backstage/newcategory.html", nil)
}

//CategorySubmit ...
//发送新建分类到后台
func CategorySubmit(c *gin.Context) {
	categoryname := c.PostForm("category_name")
	categorynoStr := c.PostForm("category_no")
	categoryno, err := strconv.Atoi(categorynoStr)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	err = service.InsertCategory(categoryname, categoryno)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/backstage/category")
}

//CategoryDelete ...
//删除分类
func CategoryDelete(c *gin.Context) {
	categoryIDStr := c.Query("category_id")
	categoryID, err := strconv.ParseInt(categoryIDStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	service.DeliverCategoryID(categoryID)
	c.Redirect(http.StatusMovedPermanently, "/backstage/category/list")
}

//CategoryUpdate ...
//获取更新分类界面
func CategoryUpdate(c *gin.Context) {
	err := utils.Session(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/loginfail.html", nil)
		return
	}
	categoryIDStr := c.Query("category_id")
	categoryID, err := strconv.ParseInt(categoryIDStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	category, err := service.GetCategory(categoryID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/backstage/updatecategory.html", gin.H{
		"category": category,
	})
}

//CategoryPost ...
//发送更新分类到后台
func CategoryPost(c *gin.Context) {
	categoryIDStr := c.Query("category_id")
	categoryID, err := strconv.ParseInt(categoryIDStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	categoryname := c.PostForm("category_name")
	categorynoStr := c.PostForm("category_no")
	categoryno, err := strconv.Atoi(categorynoStr)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	err = service.UpdateCategory(categoryname, categoryno, categoryID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/backstage/category/list")
}

//CommentList ...
//获取评论列表
func CommentList(c *gin.Context) {
	commentList, err := service.GetAllCommentList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
	}
	c.HTML(http.StatusOK, "views/backstage/commentlist.html", gin.H{
		"comment_list": commentList,
	})
}

//CommentDelete ...
//删除评论
func CommentDelete(c *gin.Context) {
	commentIDStr := c.Query("comment_id")
	commentID, err := strconv.ParseInt(commentIDStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	service.DeliverCommentID(commentID)
	c.Redirect(http.StatusMovedPermanently, "/backstage/comment/list")
}

//LeaveList ...
//
func LeaveList(c *gin.Context) {
	leaveList, err := service.GetLeaveList()
	if err != nil {
		fmt.Printf("get leave failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/backstage/leavelist.html", gin.H{
		"leave_list": leaveList,
	})
}

//LeaveDelete ...
//删除留言
func LeaveDelete(c *gin.Context) {
	leaveIDStr := c.Query("leave_id")
	leaveID, err := strconv.ParseInt(leaveIDStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/errors/500.html", nil)
		return
	}
	service.DeliverLeaveID(leaveID)
	c.Redirect(http.StatusMovedPermanently, "/backstage/leave/list")
}

//Logout ...
//退出登录
func Logout(c *gin.Context) {
	db.DeleteAllSession()
	c.Redirect(http.StatusMovedPermanently, "/")
}
