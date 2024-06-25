package controllers

import (
	"net/http"
	"strconv"

	"github.com/filipegms5/cotemig-projeto-software/models"
	"github.com/filipegms5/cotemig-projeto-software/repositories"
	"github.com/gin-gonic/gin"
)

func CreateAdocao(c *gin.Context, repo *repositories.AdocaoRepository) {
	var adocao models.Adocao
	if err := c.ShouldBindJSON(&adocao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repo.Create(&adocao); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, adocao.Id)
}

func UpdateAdocao(c *gin.Context, repo *repositories.AdocaoRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var adocao models.Adocao
	if err := c.ShouldBindJSON(&adocao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	adocao.Id = uint(id)
	if err := repo.Update(&adocao); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, adocao.Id)
}

func DeleteAdocao(c *gin.Context, repo *repositories.AdocaoRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	adocao, err := repo.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Adoção não encontrada"})
		return
	}
	if err := repo.Delete(adocao); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func GetAdocoes(c *gin.Context, repo *repositories.AdocaoRepository) {
	adocoes, err := repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, adocoes)
}

func GetAdocao(c *gin.Context, repo *repositories.AdocaoRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	adocao, err := repo.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Adoção não encontrada"})
		return
	}
	c.JSON(http.StatusOK, adocao)
}

func GetAdocoesByAdotanteId(c *gin.Context, repo *repositories.AdocaoRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	adocoes, err := repo.FindByAdotanteId(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, adocoes)
}

func GetAdocoesByIdosoId(c *gin.Context, repo *repositories.AdocaoRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	adocoes, err := repo.FindByIdosoId(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, adocoes)
}
