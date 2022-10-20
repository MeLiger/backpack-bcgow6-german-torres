package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/MeLiger/backpack-bcgow6-german-torres/GoTesting/Arquitectura.GoWeb/internal/users"
	"github.com/MeLiger/backpack-bcgow6-german-torres/GoTesting/Arquitectura.GoWeb/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name   string `json:"name" binding:"required"`
	Email  string `binding:"required"`
	Age    int    `binding:"required"`
	Height int    `binding:"required"`
	Active bool   `binding:"required"`
	Date   string `binding:"required"`
}

type UserController struct {
	service users.Service
}

func NewUserController(u users.Service) *UserController {
	return &UserController{
		service: u,
	}
}

func (c *UserController) GetAll(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token != "123456" {
		ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
		return
	}
	us, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
		return
	}
	if len(us) == 0 {
		ctx.JSON(404, web.NewResponse(404, nil, "No hay usuarios registrados"))
		return
	}
	ctx.JSON(200, web.NewResponse(200, us, ""))
}

func (c *UserController) Store(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != "123456" {
		ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
		return
	}
	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}
	if req.Name == "" {
		ctx.JSON(400, web.NewResponse(400, nil, "El nombre del usuario es requerido"))
		return
	}
	if req.Email == "" {
		ctx.JSON(400, web.NewResponse(400, nil, "La dirección de Email es requerida"))
		return
	}
	if req.Age == 0 {
		ctx.JSON(400, web.NewResponse(400, nil, "La edad es requerida"))
		return
	}
	if req.Height == 0 {
		ctx.JSON(400, web.NewResponse(400, nil, "La estatura es requerida"))
		return
	}
	if !req.Active {
		ctx.JSON(400, web.NewResponse(400, nil, "La actividad es requerida"))
		return
	}
	if req.Date == "" {
		ctx.JSON(400, web.NewResponse(400, nil, "La fecha es requerida"))
		return
	}
	p, err := c.service.Store(req.Name, req.Email, req.Age, req.Height, req.Active, req.Date)
	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}
	ctx.JSON(200, web.NewResponse(200, p, ""))
}

func (c *UserController) Update(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, gin.H{"error": "token inválido"})
		return
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID inválido"})
		return
	}
	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if req.Name == "" {
		ctx.JSON(400, gin.H{"error": "El nombre del usuario es requerido"})
		return
	}
	if req.Email == "" {
		ctx.JSON(400, gin.H{"error": "El email del usuario es requerido"})
		return
	}
	if req.Age == 0 {
		ctx.JSON(400, gin.H{"error": "La edad es requerida"})
		return
	}
	if req.Height == 0 {
		ctx.JSON(400, gin.H{"error": "La altura es requerida"})
		return
	}
	if !req.Active {
		ctx.JSON(400, gin.H{"error": "La altura es requerida"})
		return
	}
	if req.Date == "" {
		ctx.JSON(400, gin.H{"error": "La fecha es requerida"})
		return
	}
	p, err := c.service.Update(int(id), req.Name, req.Email, req.Age, req.Height, req.Active, req.Date)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, p)
}
func (c *UserController) Delete(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, gin.H{"error": "token inválido"})
		return
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "El ID es inválido"})
		return
	}
	err = c.service.Delete(int(id))
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": fmt.Sprintf("El usuario %d ha sido eliminado", id)})
}

func (c *UserController) UpdateNameAge(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, gin.H{"error": "token inválido"})
		return
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID inválido"})
		return
	}
	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if req.Name == "" {
		ctx.JSON(400, gin.H{"error": "El nombre del usuario es requerido"})
		return
	}
	if req.Age == 0 {
		ctx.JSON(400, gin.H{"error": "La edad del usuario es requerida"})
		return
	}
	p, err := c.service.UpdateNameAge(int(id), req.Name, req.Age)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, p)
}
