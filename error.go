package astigandi

import "fmt"

// Error represents an error
type Error struct {
	Cause   string `json:"cause"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Object  string `json:"object"`
}

// Error implements the error interface
func (e Error) Error() string {
	return fmt.Sprintf("code: %d - message: %s - object: %s - cause: %s", e.Code, e.Message, e.Object, e.Cause)
}
