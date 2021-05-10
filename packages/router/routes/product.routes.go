package routes

import (
	product "me-english/packages/product"
	"net/http"
)

var productRoutes = []Route{
	{
		Uri:     "/product",
		Method:  http.MethodGet,
		Handler: product.GetAll,
	},
}
