package handler

import (
	"api/app/handler/request"
	"api/app/handler/response"
	"api/app/model"
	"api/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) CreateCompany(c echo.Context) error {
	var v model.Company
	req := &request.CompanyRequest{}
	if err := req.Bind(c, &v); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := h.companyStore.Create(&v); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, response.NewCompanySingleResponse(&v))
}

func (h *Handler) ListCompanys(c echo.Context) error {
	venues, err := h.companyStore.List()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusCreated, response.NewCompanyListResponse(venues))
}

func (h *Handler) GetCompany(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	venueID := uint(id)

	venue, err := h.companyStore.GetByID(venueID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if venue == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusCreated, response.NewCompanySingleResponse(venue))
}

func (h *Handler) UpdateCompany(c echo.Context) error {
	idInt, err := strconv.Atoi(c.Param("id"))
	idUint := uint(idInt)

	v, err := h.companyStore.GetByID(idUint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if v == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	req := request.CompanyRequest{}
	req.Populate(v)
	if err := req.Bind(c, v); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := h.companyStore.Update(v); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, response.NewCompanySingleResponse(v))
}

func (h *Handler) DeleteCompany(c echo.Context) error {
	idInt, err := strconv.Atoi(c.Param("id"))
	idUint := uint(idInt)

	v, err := h.companyStore.GetByID(idUint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if v == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	if err := h.companyStore.Delete(v); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.NoContent(http.StatusOK)
}

func (h *Handler) AddOwner(c echo.Context) error {
	idInt, err := strconv.Atoi(c.Param("id"))
	idUint := uint(idInt)

	v, err := h.companyStore.GetByID(idUint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if v == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	var o model.Owner
	req := request.OwnerRequest{}
	if err := req.Bind(c, &o); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	o.CompanyID = idUint

	if err := h.ownerStore.Create(&o); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.NoContent(http.StatusOK)
}

func (h *Handler) RemoveOwner(c echo.Context) error {
	companyIdInt, err := strconv.Atoi(c.Param("id"))
	companyId := uint(companyIdInt)

	ownerIdInt, err := strconv.Atoi(c.Param("ownerId"))
	ownerId := uint(ownerIdInt)

	// Find company
	company, err := h.companyStore.GetByID(companyId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if company == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	// Find owner
	owner, err := h.ownerStore.GetByID(ownerId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if owner == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	if err := h.ownerStore.Delete(owner); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.NoContent(http.StatusOK)
}
