// TODO：持久化数据，e.g. https://dgraph.io/docs/badger/get-started/
package main

import (
	"cloudedgenetwork/store/db"
	"cloudedgenetwork/store/model"
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	listenAddrF    = flag.String("addr", ":8881", "server addr")
	k8sConfigPathF = flag.String("k8sconf", "/Users/nizhenyang/Desktop/论文 workspace/code/cloudedgenetwork/k8sconfig.yaml", "k8s config file path")
)

func main() {
	flag.Parse()
	db.InitK8SClient(*k8sConfigPathF)
	router := gin.Default()
	router.GET("/api/v1/service", func(ctx *gin.Context) {
		serviceName := ctx.Query("name")
		if len(serviceName) == 0 {
			ctx.Status(http.StatusBadRequest)
			return
		}
		svc, err := db.GetService(serviceName)
		if err != nil {
			log.Println("[err] GET query db", err)
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, svc)
	})
	router.PUT("/api/v1/service", func(ctx *gin.Context) {
		var service model.ServiceInfo
		if err := ctx.BindJSON(&service); err != nil {
			log.Println("[err] PUT bind json", err)
			ctx.Status(http.StatusInternalServerError)
			return
		}
		if err := db.PutService(&service); err != nil {
			log.Println("[err] PUT save to db", err)
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.Status(http.StatusOK)
	})
	router.DELETE("/api/v1/service", func(ctx *gin.Context) {
		serviceName := ctx.Query("name")
		if len(serviceName) == 0 {
			ctx.Status(http.StatusBadRequest)
			return
		}
		if err := db.DelService(serviceName); err != nil {
			log.Println("[err] Delete data in db", err)
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.Status(http.StatusOK)
	})
	if err := router.Run(*listenAddrF); err != nil {
		panic(err)
	}
}
