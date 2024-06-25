package controllers

import (
	"net/http"
	"strconv"

	"github.com/filipegms5/cotemig-projeto-software/models"
	"github.com/filipegms5/cotemig-projeto-software/repositories"
	"github.com/gin-gonic/gin"
)

func CreateMonitor(c *gin.Context, repo *repositories.MonitorRepository) {
	var monitor models.Monitor
	if err := c.ShouldBindJSON(&monitor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repo.Create(&monitor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, monitor.Id)
}

func UpdateMonitor(c *gin.Context, repo *repositories.MonitorRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var monitor models.Monitor
	if err := c.ShouldBindJSON(&monitor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	monitor.Id = uint(id)
	if err := repo.Update(&monitor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, monitor.Id)
}

func DeleteMonitor(c *gin.Context, repo *repositories.MonitorRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	monitor, err := repo.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monitor não encontrado"})
		return
	}
	if err := repo.Delete(monitor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func GetMonitores(c *gin.Context, repo *repositories.MonitorRepository) {
	monitores, err := repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, monitores)
}

func GetMonitor(c *gin.Context, repo *repositories.MonitorRepository) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	monitor, err := repo.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monitor não encontrado"})
		return
	}

	c.JSON(http.StatusOK, monitor)
}

func GetMonitorByUsuarioId(c *gin.Context, repo *repositories.MonitorRepository) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	monitor, err := repo.FindByUsuarioId(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monitor não encontrado"})
		return
	}

	c.JSON(http.StatusOK, monitor)
}
