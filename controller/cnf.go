package controller

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTemplateCnfFile is a func of controller process conf html request
func GetTemplateCnfFile(c *gin.Context) {
	if result, err := service.getCnfPara("static/mycnf.template"); err != nil {
		c.HTML(http.StatusOK, "error.tmpl",
			gin.H{"error": err})
	} else {
		c.HTML(http.StatusOK, "index.tmpl",
			gin.H{"html": template.HTML(result)})
	}
}
