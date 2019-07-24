package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sisimogangg/supermarket.discount.api/discount"
	"github.com/sisimogangg/supermarket.discount.api/utils"
)

type discountHandler struct {
	discountService discount.ServiceLayer
}

// NewDiscountHandler creates routes and starts listening
func NewDiscountHandler(router *mux.Router, service discount.ServiceLayer) {
	handler := &discountHandler{
		discountService: service,
	}

	router.HandleFunc("api/discount/{id}", handler.productDicount).Methods("GET")
	//router.HandleFunc("/")

}

func (h *discountHandler) productDicount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	ctx := r.Context()
	if ctx != nil {
		ctx = context.Background()
	}

	disc, err := h.discountService.GetDiscountByProductID(ctx, int32(productID))
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	resp := utils.Message(true, "success")
	resp["discount"] = disc
	utils.Respond(w, resp)
}
