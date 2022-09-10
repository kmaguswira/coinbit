package error_mapping

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

var BadRequestError = fmt.Errorf("bad request")
var UnauthorizedError = fmt.Errorf("unauthorized")
var ForbiddenError = fmt.Errorf("forbidden")
var NotFoundError = fmt.Errorf("resource could not be found")
var InternalServerError = fmt.Errorf("internal server error")

type errorMapping struct {
	FromErrors []error
	StatusCode int
	Response   func(ctx *gin.Context, err error)
}

func (r *errorMapping) ToStatusCode(statusCode int) *errorMapping {
	r.StatusCode = statusCode
	r.Response = func(ctx *gin.Context, err error) {
		ctx.JSON(statusCode, err.Error())
	}
	return r
}

func (r *errorMapping) ToResponse(response func(ctx *gin.Context, err error)) *errorMapping {
	r.Response = response
	return r
}

func Map(err ...error) *errorMapping {
	return &errorMapping{
		FromErrors: err,
	}
}

func IsType(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

var Errors = []errorMapping{
	*Map(BadRequestError).ToStatusCode(http.StatusBadRequest),
	*Map(UnauthorizedError).ToStatusCode(http.StatusUnauthorized),
	*Map(ForbiddenError).ToStatusCode(http.StatusForbidden),
	*Map(NotFoundError).ToStatusCode(http.StatusNotFound),
	*Map(InternalServerError).ToStatusCode(http.StatusInternalServerError),
}
