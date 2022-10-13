package domain

type LoginPass struct {
	ID       int
	UserID   int
	Login    string
	Password string
	MetaInfo string
}

type Text struct {
	ID       int
	UserID   int
	Text     string
	MetaInfo string
}

type Binary struct {
	ID       int
	UserID   int
	Binary   string
	MetaInfo string
}

type BankCard struct {
	ID       int
	UserID   int
	Number   string
	MetaInfo string
}
