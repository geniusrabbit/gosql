//
// @project GeniusRabbit
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016
//

package gosql

import "errors"

// Set of errors
var (
	ErrInvalidScan         = errors.New("Invalid field scan")
	ErrNullValueNotAllowed = errors.New("nil value not allowed")
	ErrInvalidDecodeValue  = errors.New("Invalid decode value")
)
