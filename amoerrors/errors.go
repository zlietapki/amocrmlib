package amoerrors

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type ErrorApi struct {
	Title            string                    `json:"title"`
	Detail           string                    `json:"detail"`
	Hint             string                    `json:"hint"`
	Type             string                    `json:"type"`
	Status           int                       `json:"status"`
	ValidationErrors []ErrorApiValidationError `json:"validation-errors"`
}

type ErrorApiValidationError struct {
	RequestId string                         `json:"request_id"`
	Errors    []ErrorApiValidationErrorError `json:"errors"`
}

type ErrorApiValidationErrorError struct {
	Code   string `json:"code"`
	Path   string `json:"path"`
	Detail string `json:"detail"`
}

func (e ErrorApi) Error() string {
	vErrs := make([]string, 0)
	for _, vErr := range e.ValidationErrors {
		vErrDetails := make([]string, 0)
		for _, ext := range vErr.Errors {
			vErrDetails = append(vErrDetails, fmt.Sprintf("code: %s path: %s detail: %s", ext.Code, ext.Path, ext.Detail))
		}
		vErrDetailsStr := strings.Join(vErrDetails, " ")
		validationErrs := fmt.Sprintf("Request id: %s Details: %s", vErr.RequestId, vErrDetailsStr)
		vErrs = append(vErrs, validationErrs)
	}
	validationErrors := strings.Join(vErrs, " ")
	return fmt.Sprintf("Title: %s Detail: %s Hint: %s ValidationErrors: %s", e.Title, e.Detail, e.Hint, validationErrors)
}

func NewErrorApi(body []byte) error {
	errApi := new(ErrorApi)
	if err := json.Unmarshal(body, errApi); err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(errApi)
}
