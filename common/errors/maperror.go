package errors

import (
	responseModel "antrian-golang/common/response/model"

	"github.com/pkg/errors"

	"net/http"
)

var (
	ErrCustomResponse   Response
	ErrSQLQueryBuilder  = errors.New("error query builder")
	ErrSQLExec          = errors.New("error sql exec")
	ErrRequiredMessage  = errors.New("require message to start money recon")
	ErrMigrate          = errors.New("failed when migrating database")
	ErrFailedParseToCSV = errors.New("failed when converting data to csv")
	ErrFailedUploadToS3 = errors.New("failed when uploading file to s3")
	ErrCustom           = errors.New("custom message")
)

type Response struct {
	StatusCode int
	Response   interface{}
}

var MapErrorResponse = map[error]Response{
	ErrSQLQueryBuilder: {
		StatusCode: http.StatusInternalServerError,
		Response: responseModel.Response{
			Message: http.StatusText(http.StatusInternalServerError),
			Status:  responseModel.StatusFail,
		},
	},

	ErrSQLExec: {
		StatusCode: http.StatusInternalServerError,
		Response: responseModel.Response{
			Message: "Database Server Failed to Execute, Please Try Again",
			Status:  responseModel.StatusFail,
		},
	},

	ErrRequiredMessage: {
		StatusCode: http.StatusBadRequest,
		Response: responseModel.Response{
			Message: "Message Required",
			Status:  responseModel.StatusFail,
		},
	},

	ErrMigrate: {
		StatusCode: http.StatusInternalServerError,
		Response: responseModel.Response{
			Message: "Failed When Migrating The Database",
			Status:  responseModel.StatusFail,
		},
	},
}

func SetDataErrCustom(statusCode int, message string, data any) {
	ErrCustomResponse = Response{
		StatusCode: statusCode,
		Response: responseModel.Response{
			Status:  responseModel.StatusFail,
			Message: message,
			Data:    data,
		},
	}
}

func SetErrCustomResponse() {
	MapErrorResponse[ErrCustom] = ErrCustomResponse
}
