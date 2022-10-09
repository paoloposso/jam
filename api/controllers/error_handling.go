package controllers

import (
	"net/http"

	customerrors "github.com/paoloposso/jam/api/src/core/custom-errors"
)

func GetHttpError(err error) (httpCode int, message string) {
	switch err.(type) {
	case *customerrors.ValidationError:
		return http.StatusBadRequest, err.Error()
	case *customerrors.NotFoundError:
		return http.StatusNotFound, err.Error()
	default:
		return http.StatusInternalServerError, err.Error()
	}
}
