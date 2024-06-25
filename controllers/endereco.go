package controllers

import (
	"net/http"
	"strconv"

	"github.com/filipegms5/cotemig-projeto-software/models"
	"github.com/filipegms5/cotemig-projeto-software/repositories"
	"github.com/gin-gonic/gin"
)

func CreateEndereco(c *gin.Context, repo *repositories.EnderecoRepository) {
	var endereco models.Endereco
	if err := c.ShouldBindJSON(&endereco); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repo.Create(&endereco); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, endereco.Id)
}

func UpdateEndereco(c *gin.Context, repo *repositories.EnderecoRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var endereco models.Endereco
	if err := c.ShouldBindJSON(&endereco); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	endereco.Id = uint(id)
	if err := repo.Update(&endereco); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, endereco.Id)
}

func DeleteEndereco(c *gin.Context, repo *repositories.EnderecoRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	endereco, err := repo.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Endereço não encontrado"})
		return
	}
	if err := repo.Delete(endereco); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func GetEnderecos(c *gin.Context, repo *repositories.EnderecoRepository) {
	enderecos, err := repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, enderecos)
}

func GetEndereco(c *gin.Context, repo *repositories.EnderecoRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	endereco, err := repo.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Endereço não encontrado"})
		return
	}
	c.JSON(http.StatusOK, endereco)
}

func GetEnderecoByCep(c *gin.Context, repo *repositories.EnderecoRepository) {
	cep := c.Param("cep")
	endereco, err := repo.FindByCep(cep)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Endereço não encontrado"})
		return
	}
	c.JSON(http.StatusOK, endereco)
}
