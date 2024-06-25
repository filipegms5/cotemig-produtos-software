package controllers

import (
	"net/http"
	"strconv"

	"github.com/filipegms5/cotemig-projeto-software/models"
	"github.com/filipegms5/cotemig-projeto-software/repositories"
	"github.com/gin-gonic/gin"
)

func CreateIdoso(c *gin.Context, repo *repositories.IdosoRepository) {
	var idoso models.Idoso
	if err := c.ShouldBindJSON(&idoso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repo.Create(&idoso); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, idoso.Id)
}

func UpdateIdoso(c *gin.Context, repo *repositories.IdosoRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var idoso models.Idoso
	if err := c.ShouldBindJSON(&idoso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idoso.Id = uint(id)
	if err := repo.Update(&idoso); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, idoso.Id)
}

func DeleteIdoso(c *gin.Context, repo *repositories.IdosoRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	idoso, err := repo.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Idoso não encontrado"})
		return
	}
	if err := repo.Delete(idoso); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func GetIdosos(c *gin.Context, repo *repositories.IdosoRepository) {
	idosos, err := repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, idosos)
}

func GetIdoso(c *gin.Context, repo *repositories.IdosoRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	idoso, err := repo.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Idoso não encontrado"})
		return
	}
	c.JSON(http.StatusOK, idoso)
}

func GetIdosoByUsuarioId(c *gin.Context, repo *repositories.IdosoRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	idoso, err := repo.FindByUsuarioId(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Idoso não encontrado"})
		return
	}
	c.JSON(http.StatusOK, idoso)
}
