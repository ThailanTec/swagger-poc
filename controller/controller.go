package controller

import (
	"net/http"

	"github.com/ThailanTec/swagger-poc/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

func BVindo(c *gin.Context) {
	c.JSON(http.StatusOK, "Bem vindo grande amigo! Vamos mudar de vida! Pode vir ao balcão!")
}

// Rota POST para criar um novo usuário
// @Summary Criar um novo usuário
// @Description Cria um novo usuário com base nos dados fornecidos
// @ID create-resource
// @Accept json
// @Produce json
// @Param request body model.UserInput true "Dados do novo usuário"
// @Success 201 {object} model.UserOutput
// @Failure 400 {object} model.ErrorResponse
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var nUser *model.UserInput
	if err := c.ShouldBindBodyWith(&nUser, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	out := model.UserOutput{
		ID:    uuid.New(),
		Name:  nUser.Name,
		Email: nUser.Email,
	}

	c.JSON(http.StatusCreated, out)
}
