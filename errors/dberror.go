package errors

type dbError struct {
	err     string
	message string
}

// Custom DB Error
func DatabaseError(errs error) {

}
