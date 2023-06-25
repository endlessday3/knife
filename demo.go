package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed public
var public embed.FS

// @Title SSwagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @Contact.Name AAPI Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}
func main() {

	router := gin.Default() // 函数返回的是一个Engine指针，Engine代表的是整个框架的一个实例，查看源码可发现实际就是调用New()方法创建实例,并且为实例添加了Logger和Recovery中间件.

	// 调用Engine的GET方法（其他请求方式POST PUT DELETE等）  第一个参数为相对路径，第二个参数为多个handle
	router.GET("/test", Test)

	if os.Getenv(gin.EnvGinMode) != gin.ReleaseMode {
		// https://stackoverflow.com/questions/66248258/serve-embedded-filesystem-from-root-path-of-url
		publicDir, err := fs.Sub(public, "public")
		if err != nil {
			log.Fatalln(err)
		}
		router.StaticFS("/docs/", http.FS(publicDir))
		//router.StaticFS("/api2", http.Dir("./public"))
	}

	//v1组路由  把Print和Hello方法添加到同一组路由，即访问时需要在前面加上/api/v1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/print", Print)
		v1.GET("/hello", Hello)
	}

	// 不指定ip地址和端口时，默认是监听并在 0.0.0.0:8080 上启动服务，另外的写法还有Run(":8080")、Run("0.0.0.0:8080")、Run("localhost:8080")都是指定http://localhost:8080或者http://127.0.0.1:8080/
	err := router.Run()
	log.Println(err)
}

// Print xxx
// @Summary 打印测试功能
// @title Swagger Example API
// @version 0.0.1
// @description  This is a sample server Petstore server.
// @BasePath /api/v1
// @Host 127.0.0.1:8080
// @Produce  json
// @Param	name	query	string true "Name"
// @Success	200 	{string}	json "{"code":200,"data":"name","msg":"ok"}"
// @Router /print [get]
// @Router /print [post]
func Print(c *gin.Context) {
	var (
		name string
	)
	name = c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": name,
	})
}

type Account struct {
	ID   int    `json:"id" example:"1"`              // iidddd
	Name string `json:"name" example:"account name"` // 姓名
}

type Account2 struct {
	ID   int    `json:"id"`   // iidddd
	Name string `json:"name"` // 姓名
}

// Hello 你好
// @Summary Hello接口
// @Description Hello接口
// @Accept  json
// @Produce  json
// @Tags 用户信息
// @Param message body Account2 false "name search by q"
// @ID op-Hello
// @Success 200 {string} json "{"message":"success"}"
// @Success 300 {array} Account2 "lalala"
// asd
// @Router /hello [get]
func Hello(c *gin.Context) {
	// 当响应码为200时，返回JSON格式数据
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// Test 测试接口讷讷哈哈
// @Summary 测试接口
// @Description 描述信息
// @Success 200 {string} string  "ok"
// @Router /test [get]
func Test(c *gin.Context) {
	c.JSON(200, "ok")
}

// ListAccounts godoc
// @Summary List accounts
// @ID ListAccounts
// @Description get accounts
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q"
// @Success 200 {string} 123
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Failure default {object} httputil.HTTPError
// @Router /accounts [get]
func ListAccounts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "account")
}
