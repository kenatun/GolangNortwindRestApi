package order

import (
	"context"
	"net/http"
	"strconv"

	"github.com/GolangNortwindRestApi/helper"
	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getOrderByIdRequest := kithttp.NewServer(makeGetOrderByIdEndPoint(s),
		getOrderByIdRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getOrderByIdRequest)
	return r

}

func getOrderByIdRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	orderId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	helper.Catch(err)
	return getOrderByIdRequest{
		orderId: orderId,
	}, nil
}
