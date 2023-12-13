package errors

import (
	"errors"
	"fmt"
	"net/http"
	"restaurantapi/utils"
	"strconv"

	"github.com/Christochi/error-handler/service"
)

var errorLogger = utils.InfoLog() // return info field

type apiError struct {
	err    string
	status int
}

func (a *apiError) Error() string {
	return fmt.Sprintf("Error in JSON: %v -- %v", a.err, http.StatusText(a.status))
}

// Custom API Error
func RestError(rw http.ResponseWriter, errs error) *apiError {
	var apierr apiError
	var svcErr *service.ServiceError

	if errors.As(errs, &svcErr) {
		apierr.err = svcErr.AppError().Error()
		var atoiErr error
		apierr.status, atoiErr = strconv.Atoi(svcErr.ErrDesc())
		if atoiErr != nil {
			errorLogger.Println(atoiErr)
		}
	}

	http.Error(rw, apierr.err, apierr.status)

	return &apierr

}
