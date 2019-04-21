package response

import (
	"encoding/json"
	"fmt"
	"github.com/Jamsek-m/config-server/errors"
	"net/http"
)

type JsonArgs struct {
	Data    interface{}
	Status  int
	Headers map[string]string
}

func Json(res http.ResponseWriter, args JsonArgs) {
	if args.Status == 0 {
		args.Status = http.StatusOK
	}

	if args.Headers != nil {
		if _, present := args.Headers[CONTENT_TYPE]; !present {
			res.Header().Add(CONTENT_TYPE, APP_JSON)
		}
		for headerName, headerValue := range args.Headers {
			res.Header().Add(headerName, headerValue)
		}
	} else {
		res.Header().Add(CONTENT_TYPE, APP_JSON)
	}

	res.WriteHeader(args.Status)

	if args.Data != nil {
		jsonEncodeErr := json.NewEncoder(res).Encode(args.Data)
		if jsonEncodeErr != nil {
			fmt.Println("Error setting json response!")
		}
	}
}

func HandleError(res http.ResponseWriter, err error) {
	res.Header().Add(CONTENT_TYPE, APP_JSON)
	statusCode := determineStatusCode(err)
	res.WriteHeader(statusCode)
	entity := make(map[string]interface{})
	entity["status"] = statusCode
	entity["message"] = err.Error()
	jsonEncodeErr := json.NewEncoder(res).Encode(entity)
	if jsonEncodeErr != nil {
		fmt.Println("Error setting json response!")
	}
}

func determineStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case errors.NotFoundError:
		return http.StatusNotFound
	case errors.BadRequestError:
		return http.StatusBadRequest
	case errors.ConflictError:
		return http.StatusConflict
	case errors.NilValidationFailedError, errors.SemanticValidationFailedError:
		return http.StatusUnprocessableEntity
	case errors.UnAuthorizedError:
		return http.StatusUnauthorized
	case errors.ForbiddenError:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
