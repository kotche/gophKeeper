package transport

const (
	//data access
	registration       = "-reg"
	registrationDesc   = "registration: <-reg> login password"
	authentication     = "-auth"
	authenticationDesc = "authentication: <-auth> login password"

	//data type, data format
	loginPassDataType     = "lp"
	loginPassDataTypeDesc = "login password <-meta>(optional) ..."
	textDataType          = "text"
	textDataTypeDesc      = "text data <-meta>(optional) ..."
	binaryDataType        = "binary"
	binaryDataTypeDesc    = "binary data <-meta>(optional) ..."
	bankCardDataType      = "bank"
	bankCardDataTypeDesc  = "number <-meta>(optional) ..."

	metaInfo = "-meta"

	//data management
	create     = "-create"
	createDesc = "<-create> 'data type' 'data format'"
	read       = "-read"
	readDesc   = "<-read> 'data type'"
	update     = "-update"
	updateDesc = "<-update> 'data type' 'id data' 'data format'"
	delete     = "-delete"
	deleteDesc = "<-delete> 'data type' 'id data'"

	//data errors
	invalidFormat = "invalid format"
)
