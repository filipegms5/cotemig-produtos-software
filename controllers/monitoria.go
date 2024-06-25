package controllers

import (
	"net/http"
	"strconv"

	"github.com/filipegms5/cotemig-projeto-software/models"
	"github.com/filipegms5/cotemig-projeto-software/repositories"
	"github.com/gin-gonic/gin"
)

func CreateMonitoria(c *gin.Context, repo *repositories.MonitoriaRepository) {
	var monitoria models.Monitoria
	if err := c.ShouldBindJSON(&monitoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repo.Create(&monitoria); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, monitoria.Id)
}

func UpdateMonitoria(c *gin.Context, repo *repositories.MonitoriaRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var monitoria models.Monitoria
	if err := c.ShouldBindJSON(&monitoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	monitoria.Id = uint(id)
	if err := repo.Update(&monitoria); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, monitoria.Id)
}

func DeleteMonitoria(c *gin.Context, repo *repositories.MonitoriaRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	monitoria, err := repo.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monitoria não encontrada"})
		return
	}
	if err := repo.Delete(monitoria); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func GetMonitorias(c *gin.Context, repo *repositories.MonitoriaRepository) {
	monitorias, err := repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, monitorias)
}

func GetMonitoria(c *gin.Context, repo *repositories.MonitoriaRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	monitoria, err := repo.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monitoria não encontrada"})
		return
	}
	c.JSON(http.StatusOK, monitoria)
}

func GetMonitoriasByMonitorId(c *gin.Context, repo *repositories.MonitoriaRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	monitorias, err := repo.FindByMonitorId(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monitoria não encontrada"})
		return
	}
	c.JSON(http.StatusOK, monitorias)
}

func GetMonitoriasByAdocaoId(c *gin.Context, repo *repositories.MonitoriaRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	monitorias, err := repo.FindByAdocaoId(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monitoria não encontrada"})
		return
	}
	c.JSON(http.StatusOK, monitorias)
}
