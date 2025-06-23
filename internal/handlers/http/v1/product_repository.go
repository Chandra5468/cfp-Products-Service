package v1

import (
	"log"
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
	router.Get("/v1/api/product/{productId}", h.getAProduct)
	router.Post("/v1/api/product/{productId}/buy", h.updateProducts) // This will be serving both checkout after adding to cart and directly clicking on buy
	router.Get("/v1/api/cart/{productId}/check-quantity", h.productsCount)
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

	pId := chi.URLParam(r, "productId")
	log.Println("THis is pid--------", pId)
	if pId == "" {
		responses.WriteJson(w, http.StatusBadRequest, "please send productId in url")
		return
	}

	pd, err := h.store.GetProduct(pId)

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

}

func (h *Handler) deleteProduct(w http.ResponseWriter, r *http.Request) {

}
