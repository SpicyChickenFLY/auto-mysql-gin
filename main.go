package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SpicyChickenFLY/auto-mysql-gin/backend/controller"
	"github.com/SpicyChickenFLY/auto-mysql-gin/backend/pkgs/middleware"
	"github.com/SpicyChickenFLY/auto-mysql-gin/backend/pkgs/mysql"
	"github.com/SpicyChickenFLY/auto-mysql/controller"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

const ( // GIN CONFIG
	ginPort = ":8080"
)

func main() {
	// load middlewares

	router := gin.Default()
	router.Use(middleware.Cors())
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	// Group: Todo List
	groupAPI := router.Group("/api")
	{
		groupVersion1 := groupAPI.Group("/v1")
		{
			groupMysqlInstaller := groupVersion1.Group("/auto-mysql")
			{
				groupMysqlInstaller.GET("/ui", controller.ShowMysqlInstallerUI)
				groupMysqlInstaller.POST("/standard", controller.InstallStandardInstances)
				groupMysqlInstaller.POST("/custom", controller.InstallCustomInstances)
				groupMysqlInstaller.DELETE("/instance", controller.RemoveInstances)
			}
			groupCnfManager := groupVersion1.Group("/auto-mycnf")
			{
				groupCnfManager.GET("/ui", controller.ShowCnfManagerUI)
				groupCnfManager.GET("/cnf", controller.GetTemplateCnfFile)
				groupCnfManager.POST("/cnf", controller.AddNewCnfFile)
			}
		}
	}

	server := &http.Server{
		Addr:    ginPort,
		Handler: router,
	}

	go func() {
		// service connections
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Println("server encount error while listen and serve:", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 1 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 1 seconds.")
	}
	mysql.CloseGormConn()
	log.Println("Server exiting")
}
