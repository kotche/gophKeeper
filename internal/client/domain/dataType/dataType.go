package dataType

// DataType data types
type DataType struct {
	string
}

var (
	LP       = DataType{"lp"}
	TEXT     = DataType{"test"}
	BINARY   = DataType{"binary"}
	BANKCARD = DataType{"bank"}
)

func (d DataType) String() string {
	return d.string
}
