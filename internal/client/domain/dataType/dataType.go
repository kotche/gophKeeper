package dataType

import "errors"

// DataType data types
type DataType struct {
	string
}

var (
	ErrInvalidDataType = errors.New("unknown data type")

	UNKNOWN  = DataType{"UNKNOWN"}
	LP       = DataType{"lp"}
	TEXT     = DataType{"test"}
	BINARY   = DataType{"binary"}
	BANKCARD = DataType{"bank"}
)

func (d DataType) String() string {
	return d.string
}

func GetDataType(s string) (DataType, error) {
	switch s {
	case LP.string, TEXT.string, BINARY.string, BANKCARD.string:
		return DataType{s}, nil
	default:
		return UNKNOWN, ErrInvalidDataType
	}
}
