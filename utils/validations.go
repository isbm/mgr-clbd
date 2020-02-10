package utils

import (
	"fmt"
	"net/http"
	"strings"
)

type Validators struct{}

func NewValidators() *Validators {
	return new(Validators)
}

// VerifyRequired verifies required fields in the values object and
// returns an error code with a message (or 200 and no message).
func (v *Validators) VerifyRequired(request *http.Request, fields ...string) (int, string) {
	var status int
	var msg string
	missing := make([]string, 0)

	for _, field := range fields {
		if request.Form.Get(field) == "" {
			missing = append(missing, field)
		}
	}

	if len(missing) > 0 {
		status = http.StatusUnprocessableEntity
		msg = fmt.Sprintf("Missing parameters: %s.", strings.Join(missing, ", "))
	} else {
		status = http.StatusOK
	}

	return status, msg
}
