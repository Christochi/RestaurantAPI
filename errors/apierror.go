package errror

import (
	"errors"
	"net/http"
	"restaurantapi/utils"
	"strconv"

	"github.com/Christochi/error-handler/service"
)

var requestLogger = utils.InfoLog() // return info field

type apiError struct {
	err    string
	status int
}

// Custom API Error
func RestError(rw http.ResponseWriter, err error) {
	var apierr apiError
	var svcErr *service.ServiceError

	if errors.As(err, &svcErr) {
		apierr.err = svcErr.AppError().Error()
		var atoiErr error
		apierr.status, atoiErr = strconv.Atoi(svcErr.ErrDesc())
		if atoiErr != nil {
			requestLogger.Println(atoiErr)
		}
	}

	http.Error(rw, apierr.err, apierr.status)

}
