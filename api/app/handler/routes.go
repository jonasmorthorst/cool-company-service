package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(root *echo.Group) {
	v1 := root.Group("/v1")

	v1.GET("/companies", h.ListCompanys)
	v1.GET("/companies/:id", h.GetCompany)
	v1.POST("/companies", h.CreateCompany)
	v1.PATCH("/companies/:id", h.UpdateCompany)
	v1.DELETE("/companies/:id", h.DeleteCompany)

	v1.POST("/companies/:id/owners", h.AddOwner)
	v1.DELETE("/companies/:id/owners/:ownerId", h.RemoveOwner)
}
