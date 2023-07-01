package errors

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorReport struct {
	StatusCode int
	Response   *ErrorResponse
}

func ErrorHandler(err error, c echo.Context) {
	report := UnwrapError(err)
	//pperr.Print(err)

	if !c.Response().Committed {
		_ = c.JSON(report.StatusCode, report.Response)
	}
}

func UnwrapError(err error) *ErrorReport {

	for i := 0; i < 10; i++ {
		switch e := err.(type) {

		case interface{ Unwrap() error }:
			if e.Unwrap() != nil {
				err = e.Unwrap()
			}

		case interface{ Cause() error }:
			if e.Cause() != nil {
				err = e.Cause()
			}

		// OpenAPI validation error
		case openapi3.MultiError:
			if len(e) > 0 {
				err = e[0]
			} else {
				break
			}

		// OpenAPIs authentication error
		case *openapi3filter.SecurityRequirementsError:
			if len(e.Errors) > 0 {
				err = e.Errors[0]
			} else {
				break
			}

		default:
			return &ErrorReport{
				StatusCode: 500,
				Response: &ErrorResponse{
					Code:    0,
					Message: e.Error(),
				},
			}
		}
	}

	return &ErrorReport{
		StatusCode: 500,
		Response: &ErrorResponse{
			Code:    1,
			Message: "unexpected error.",
		},
	}
}
