package employee

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/GolangNortwindRestApi/helper"
	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()
	getEmployeesHandler := kithttp.NewServer(makeGetEmployeesEndPoint(s),
		getEmployeesRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getEmployeesHandler)

	getEmployeeByIdHandler := kithttp.NewServer(makeEmployeeByIdEndpoint(s),
		getEmployeeByIdRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getEmployeeByIdHandler)

	getBestEmployeehandler := kithttp.NewServer(makeGetbestEmployeeEndPoint(s),
		getBestEmployeeRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/best", getBestEmployeehandler)

	addEmployeeHandler := kithttp.NewServer(makeInsertemployeeEndpoint(s),
		getAddEmployeeRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", addEmployeeHandler)

	updateEmployeeHandler := kithttp.NewServer(makeUpdateEmployeeEndpoint(s),
		getUpdateEmployeeRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", updateEmployeeHandler)

	return r
}

func getEmployeesRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getEmployeesRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getEmployeeByIdRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return getEmployeeByIdRequest{
		EmployeeId: chi.URLParam(r, "id"),
	}, nil
}

func getBestEmployeeRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return getBestEmployeeRequest{}, nil
}

func getAddEmployeeRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := addEmployeeRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)

	return request, nil
}

func getUpdateEmployeeRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := updateEmployeeRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}
