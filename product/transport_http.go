package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getProductByIdHandler := kithttp.NewServer(makeProductByIdEndpoint(s),
		getProductByIdRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getProductByIdHandler)

	getProducthandler := kithttp.NewServer(makeGetProductsEndPoint(s),
		getProductsRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getProducthandler)

	addProducthandler := kithttp.NewServer(MakeAddproductEndPoint(s),
		addProductRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", addProducthandler)

	updateProducthandler := kithttp.NewServer(makeUpdateProductEndPoint(s),
		updateproductRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", updateProducthandler)

	deleteProductHandler := kithttp.NewServer(makeDeleteProductendPoint(s),
		getDeleteProductRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{id}", deleteProductHandler)
	return r
}

func getProductByIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	productId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getProductByIdRequest{
		ProductID: productId,
	}, nil

}

func getProductsRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}

func addProductRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getAddProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}

func updateproductRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := updateProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}

func getDeleteProductRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {

	return deleteProductRequest{
		ProductID: chi.URLParam(r, "id"),
	}, nil

}
