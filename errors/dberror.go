package errors

import (
	"errors"
	"fmt"

	"github.com/Christochi/error-handler/service"
)

type dbError struct {
	err     string
	message string
}

func (db *dbError) Error() string {
	return fmt.Sprintf("DB Error: %v -- %v", db.err, db.message)
}

// Custom DB Error
func DatabaseError(errs error) *dbError {
	var dberr dbError
	var svcErr *service.ServiceError

	if errors.As(errs, &svcErr) {
		dberr.err = svcErr.AppError().Error()
		dberr.message = svcErr.ErrDesc()
	}

	return &dberr

}
