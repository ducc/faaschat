package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) (string, error) {
	return fmt.Sprintf("Hello, Go. You said: %s", string(req)), nil
}
