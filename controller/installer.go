package controller

import (
	"net/http"

	"github.com/SpicyChickenFLY/auto-mysql/service"
	"github.com/gin-gonic/gin"
)

// ShowInstallerUI show GUI for user
func ShowInstallerUI(c *gin.Context) {

}

// InstallStandardInstances install instances with std mycnf file
func InstallStandardInstances(c *gin.Context) {

}

// InstallCustomInstances install instance with custom mycnf file
func InstallCustomInstances(c *gin.Context) {
	param := c.Param("param")
	service.InstallCustomInstance(param)
	c.JSON(http.StatusOK, gin.H{})
}

// RemoveInstances remove instance
func RemoveInstances(c *gin.Context) {
	// param := c.Param("param")
	c.String(http.StatusOK, "not supported now!")
}
