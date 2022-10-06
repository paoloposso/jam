package controllers

import (
	customerrors "api/src/core/custom-errors"
	"net/http"
)

func GetHttpError(err error) (httpCode int, message string) {
	switch err.(type) {
	case *customerrors.ValidationError:
		return http.StatusBadRequest, err.Error()
	default:
		return http.StatusInternalServerError, err.Error()
	}
}
