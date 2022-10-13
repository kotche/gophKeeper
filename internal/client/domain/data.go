package domain

type LoginPass struct {
	ID       int
	Login    string
	Password string
	MetaInfo string
}

type Text struct {
	ID       int
	Text     string
	MetaInfo string
}

type Binary struct {
	ID       int
	Binary   string
	MetaInfo string
}

type BankCard struct {
	ID       int
	Number   string
	MetaInfo string
}
