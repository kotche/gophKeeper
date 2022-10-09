package dataType

import "errors"

type DataType struct {
	string
}

var (
	ErrPlatformInvalidParam = errors.New("unknown data type")

	LP       = DataType{"lp"}
	TEXT     = DataType{"test"}
	BINARY   = DataType{"binary"}
	BANKCARD = DataType{"bank"}
	UNKNOWN  = DataType{"unknown"}
)

func (d DataType) String() string {
	return d.string
}
