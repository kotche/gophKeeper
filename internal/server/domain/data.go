package domain

// LoginPass login-password data
type LoginPass struct {
	ID       int
	UserID   int
	Login    string
	Password string
	MetaInfo string
}

// Text text data
type Text struct {
	ID       int
	UserID   int
	Text     string
	MetaInfo string
}

// Binary binary data
type Binary struct {
	ID       int
	UserID   int
	Binary   string
	MetaInfo string
}

// BankCard bank card data
type BankCard struct {
	ID       int
	UserID   int
	Number   string
	MetaInfo string
}
