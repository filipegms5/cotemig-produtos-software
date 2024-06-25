package controllers

import (
	"net/http"
	"strconv"

	"github.com/filipegms5/cotemig-projeto-software/models"
	"github.com/filipegms5/cotemig-projeto-software/repositories"
	"github.com/gin-gonic/gin"
)

func CreateAdotante(c *gin.Context, repo *repositories.AdotanteRepository) {
	var adotante models.Adotante
	if err := c.ShouldBindJSON(&adotante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repo.Create(&adotante); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, adotante.Id)
}

func UpdateAdotante(c *gin.Context, repo *repositories.AdotanteRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var adotante models.Adotante
	if err := c.ShouldBindJSON(&adotante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	adotante.Id = uint(id)
	if err := repo.Update(&adotante); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, adotante.Id)
}

func DeleteAdotante(c *gin.Context, repo *repositories.AdotanteRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	adotante, err := repo.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Adotante não encontrado"})
		return
	}
	if err := repo.Delete(adotante); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func GetAdotante(c *gin.Context, repo *repositories.AdotanteRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	adotante, err := repo.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Adotante não encontrado"})
		return
	}
	c.JSON(http.StatusOK, adotante)
}

func GetAdotantes(c *gin.Context, repo *repositories.AdotanteRepository) {
	adotantes, err := repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, adotantes)
}

func GetAdotanteByUsuarioID(c *gin.Context, repo *repositories.AdotanteRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	adotante, err := repo.FindByUsuarioID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Adotante não encontrado"})
		return
	}
	c.JSON(http.StatusOK, adotante)
}
