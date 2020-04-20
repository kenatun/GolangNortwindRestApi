package order

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GolangNortwindRestApi/helper"
	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getOrderByIdHandler := kithttp.NewServer(makeGetOrderByIdEndPoint(s),
		getOrderByIdRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getOrderByIdHandler)

	getOrdersHandler := kithttp.NewServer(makeGetOrdersEndPoint(s),
		getOrderRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getOrdersHandler)
	return r

}

func getOrderByIdRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	orderId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	helper.Catch(err)
	return getOrderByIdRequest{
		orderId: orderId,
	}, nil
}

func getOrderRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getOrdersRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)

	return request, nil
}
