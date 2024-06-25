package controllers

import (
	"net/http"
	"strconv"

	"github.com/filipegms5/cotemig-projeto-software/models"
	"github.com/filipegms5/cotemig-projeto-software/repositories"
	"github.com/gin-gonic/gin"
)

func CreateUsuario(c *gin.Context, repo *repositories.UsuarioRepository) {
	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repo.Create(&usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, usuario.Id)
}

func UpdateUsuario(c *gin.Context, repo *repositories.UsuarioRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usuario.Id = uint(id)
	if err := repo.Update(&usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, usuario.Id)
}

func DeleteUsuario(c *gin.Context, repo *repositories.UsuarioRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	usuario, err := repo.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}
	if err := repo.Delete(usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func GetUsuario(c *gin.Context, repo *repositories.UsuarioRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	usuario, err := repo.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}
	c.JSON(http.StatusOK, usuario)
}

func GetUsuarios(c *gin.Context, repo *repositories.UsuarioRepository) {
	usuarios, err := repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, usuarios)
}

func GetUsuarioByEmail(c *gin.Context, repo *repositories.UsuarioRepository) {
	email := c.Param("email")
	usuario, err := repo.FindByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}
	c.JSON(http.StatusOK, usuario)
}

func GetUsuarioByCpf(c *gin.Context, repo *repositories.UsuarioRepository) {
	cpf := c.Param("cpf")
	usuario, err := repo.FindByCpf(cpf)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}
	c.JSON(http.StatusOK, usuario)
}

func GetUsuarioByName(c *gin.Context, repo *repositories.UsuarioRepository) {
	name := c.Param("name")
	usuarios, err := repo.FindByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}
	c.JSON(http.StatusOK, usuarios)
}
