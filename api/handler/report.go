package handler

import (
	"app/api/models"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UPDATE Product Transaction godoc
// @ID update_product_transaction
// @Router /product_transaction [PUT]
// @Summary Update Product Transaction
// @Description Update Product Transaction
// @Tags Report
// @Accept json
// @Produce json
// @Param productTransaction body models.FromToStock true "UpdateFromToStockRequest"
// @Success 204 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) FromStock_ToStock(c *gin.Context) {

	var Body models.FromToStock

	err := c.ShouldBindJSON(&Body)
	if err != nil{
		h.handlerResponse(c, "Product Transaction", http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(Body)

	rowsAffected, err := h.storages.Raport().FromStock_ToStock(context.Background(), &Body)
	if err != nil{
		h.handlerResponse(c, "Storaga Product Transaction", 500, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Product Transaction", 400, "No rows affected")
		return
	}

	h.handlerResponse(c, "Product Transaction", http.StatusAccepted, models.FromToStock{})
}

// Get List Staff Report godoc
// @ID get_list_staff_report
// @Router /staff_report [GET]
// @Summary Get List Staff Report
// @Description Get List Staff Report
// @Tags Report
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetStaffSell_Report(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list order", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list order", http.StatusBadRequest, "invalid limit")
		return
	}


	resp, err := h.storages.Raport().GetListStaffSell(context.Background(), &models.GetListStaffSellRequest{
		Offset: offset,
		Limit: limit,
	})
	if err != nil{
		h.handlerResponse(c, "Get Staff Sell Raport", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Get Staff Sell Raport", http.StatusOK, resp)
}


// Get order ID Total_sum godoc
// @ID get_by_orderid_report
// @Router /order_total_sum [GET]
// @Summary Get By order TotalSum
// @Description Get By ID TotalSum
// @Tags Report
// @Accept json
// @Produce json
// @Param order_id query string true "order_id"
// @Param promocode query string false "promocode"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) Order_Total_sum(c *gin.Context) {
	
	orderid := c.Query("order_id")
	order_id, err := strconv.Atoi(orderid)
	if err != nil{
		h.handlerResponse(c, "Order Total Sum", 400, err.Error())
		return
	}


	resp, err := h.storages.Raport().Total_sum(context.Background(), &models.TotalSum{
		Order_id: order_id,
		Promo_code: c.Query("promocode"),
	})

	if err != nil{
		h.handlerResponse(c, "Stotage Report Total Sum", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Report Total Sum", 200, resp)

}