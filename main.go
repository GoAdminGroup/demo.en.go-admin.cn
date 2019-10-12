package main

import (
	"github.com/GoAdminGroup/demo_en/ecommerce"
	"github.com/GoAdminGroup/demo_en/login"
	"github.com/GoAdminGroup/demo_en/pages"
	_ "github.com/chenhg5/go-admin/adapter/gin"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// add generator, first parameter is the url prefix of table when visit.
	// example:
	//
	// "user" => http://localhost:9033/admin/info/user
	//
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	template.AddLoginComp(login.GetLoginComponent())

	rootPath := "/data/www/go-admin-en"
	//rootPath = "."

	if err := eng.AddConfigFromJson(rootPath + "/config.json").AddPlugins(adminPlugin).Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", rootPath+"/uploads")

	// you can custom your pages like:

	r.GET("/admin", func(ctx *gin.Context) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return ecommerce.GetContent()
		})
	})

	r.GET("/admin/form1", func(ctx *gin.Context) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return pages.GetForm1Content()
		})
	})

	r.GET("/admin/e-commerce", func(ctx *gin.Context) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return ecommerce.GetContent()
		})
	})

	_ = r.Run(":9032")
}
