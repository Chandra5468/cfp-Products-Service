package v1

import (
	"encoding/json"
	"net/http"

	"github.com/Chandra5468/cfp-Products-Service/internal/types"
	"github.com/Chandra5468/cfp-Products-Service/internal/utils/responses"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store types.ProductsStore
}

func NewHandler(store types.ProductsStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	// User level routes
	router.Get("/v1/api/product/{name}", h.getAProduct)
	router.Get("/v1/api/product/{productId}/check-quantity", h.productsCount)

	router.Post("/v1/api/cart/{productId}/buy", h.updateProducts) // This will be serving both checkout after adding to cart and directly clicking on buy
	// router.Post("/v1/api/cart/{productId}/remove") // This is for removing products back from cart at checkout

	// Routes for admin purposes
	router.Get("/v1/api/admin/products", h.getAllProducts)
	router.Post("/v1/api/admin/product", h.addProducts)                 // Add a product
	router.Put("/v1/api/admin/product/{productId}", h.updateProducts)   // Update product may be quantity, may be locations of availability
	router.Delete("/v1/api/admin/product/{productId}", h.deleteProduct) // Delete a product
}

func (h *Handler) getAProduct(w http.ResponseWriter, r *http.Request) {
	//
	// gpd := &types.GetProduct{}
	// err := json.NewDecoder(r.Body).Decode(gpd)
	// if err != nil {
	// 	responses.WriteJson(w, http.StatusBadRequest, "body is not getting deserialized")
	// 	return
	// }
	// pd, err := h.store.GetProduct(gpd.Name)

	name := chi.URLParam(r, "name")
	if name == "" {
		responses.WriteJson(w, http.StatusBadRequest, "please send product name in url")
		return
	}

	pd, err := h.store.GetProduct(name)

	if err != nil {
		responses.WriteJson(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		responses.WriteJson(w, http.StatusOK, pd)
	}
}

func (h *Handler) productsCount(w http.ResponseWriter, r *http.Request) {

}

// Admin level handlers

func (h *Handler) getAllProducts(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) addProducts(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) updateProducts(w http.ResponseWriter, r *http.Request) {
	// Deserialze req.body

	cart := &types.BuyCart{}

	err := json.NewDecoder(r.Body).Decode(cart)
	if err != nil {
		responses.WriteJson(w, http.StatusBadRequest, "unable to deserialize body while puchasing products")
		return
	}

	productId := cart.ProductId

	pDetails, err := h.store.GetProductByID(&productId)
	if err != nil {
		responses.WriteJson(w, http.StatusInternalServerError, "error while fetching product information before purchase")
		return
	}
	// Check for quantity available. And reduce the total available quantity by asked and available.
	var status int8
	var totalAmount float32
	var quantity int16

	if pDetails.Quantity == 0 {
		status = 0
		totalAmount = 0
		quantity = 0
	} else if cart.Quantity < pDetails.Quantity {
		totalAmount = float32(cart.Quantity) * pDetails.Price
		status = 1
		quantity = pDetails.Quantity - cart.Quantity
	} else {
		totalAmount = float32(pDetails.Quantity) * pDetails.Price
		status = 2
		quantity = 0
	}
	// quantity * price

	// Another query to deduct from products. productId, quantity

	err = h.store.UpdateProductsQuantity(pDetails.Id, &quantity)
	if err != nil {
		responses.WriteJson(w, http.StatusInternalServerError, "there is some error while updating products")
		return
	}
	// Return PurchasedProducts

	purchased := &types.PurchasedProducts{
		ProductId:   cart.ProductId,
		Quantity:    quantity,
		TotalAmount: totalAmount,
		Status:      status,
	}

	responses.WriteJson(w, http.StatusOK, purchased)

}

func (h *Handler) deleteProduct(w http.ResponseWriter, r *http.Request) {

}
