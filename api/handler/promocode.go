package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create PromoCode godoc
// @ID create_promo_code
// @Router /promocode [POST]
// @Summary Create PromoCode
// @Description Create PromoCode
// @Tags PromoCode
// @Accept json
// @Produce json
// @Param promocode body models.CreatePromoCode true "CreatePromoCodeRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreatePromoCode(c *gin.Context) {

	var createPromoCode models.CreatePromoCode

	err := c.ShouldBindJSON(&createPromoCode) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create promoCode", http.StatusBadRequest, err.Error())
		return
	}

	if createPromoCode.Discount_Type != "fixed" && createPromoCode.Discount_Type != "percent" {
		h.handlerResponse(c, "Promocode Create", 400, "No such type of promocode")
		return
	}

	id, err := h.storages.PromoCode().CreatePromoCode(context.Background(), &createPromoCode)
	if err != nil {
		h.handlerResponse(c, "storage.promoCode.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.PromoCode().GetByID(context.Background(), &models.PromoCodePrimaryKey{PromoCodeId: id})
	if err != nil {
		h.handlerResponse(c, "storage.promoCode.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create promoCode", http.StatusCreated, resp)
}

// Get By ID PromoCode godoc
// @ID get_by_id_PromoCode
// @Router /promocode/{id} [GET]
// @Summary Get By ID PromoCode
// @Description Get By ID PromoCode
// @Tags PromoCode
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdPromoCode(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.promoCode.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	resp, err := h.storages.PromoCode().GetByID(context.Background(), &models.PromoCodePrimaryKey{PromoCodeId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.promoCode.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get promoCode by id", http.StatusCreated, resp)
}

// Get List PromoCode godoc
// @ID get_list_promocode
// @Router /promocode [GET]
// @Summary Get List PromoCode
// @Description Get List PromoCode
// @Tags PromoCode
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListPromoCode(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list PromoCode", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list PromoCode", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.PromoCode().GetListPromoCode(context.Background(), &models.GetListPromoCodeRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.promocode.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list PromoCode response", http.StatusOK, resp)
}

// Update PromoCode godoc
// @ID update_promocode
// @Router /promocode/{id} [PUT]
// @Summary Update PromoCode
// @Description Update PromoCode
// @Tags PromoCode
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param promocode body models.UpdatePromoCode true "UpdatePromoCodeRequest"
// @Success 202 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePromoCode(c *gin.Context) {

	var updatePromoCode models.UpdatePromoCode

	id := c.Param("id")

	err := c.ShouldBindJSON(&updatePromoCode)
	if err != nil {
		h.handlerResponse(c, "update PromoCode", http.StatusBadRequest, err.Error())
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.PromoCode.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	updatePromoCode.PromoCodeId = idInt

	rowsAffected, err := h.storages.PromoCode().UpdatePromoCode(context.Background(), &updatePromoCode)
	if err != nil {
		h.handlerResponse(c, "storage.promocode.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.promocode.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.PromoCode().GetByID(context.Background(), &models.PromoCodePrimaryKey{PromoCodeId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.promocode.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update promocode", http.StatusAccepted, resp)
}

// Delete PromoCode godoc
// @ID delete_promocode
// @Router /promocode/{id} [DELETE]
// @Summary Delete PromoCode
// @Description Delete PromoCode
// @Tags PromoCode
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error
func (h *Handler) DeletePromoCode(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.promocode.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	rowsAffected, err := h.storages.PromoCode().Delete(context.Background(), &models.PromoCodePrimaryKey{PromoCodeId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.promocode.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.promocode.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete promocode", http.StatusNoContent, nil)
}
