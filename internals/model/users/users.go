package users

type Request struct {
	FirstName   string
	LastName    string
	Address     string
	PhoneNumber string
	Gender      string
}

type FitterReadUsers struct {
	FullName string
	Id       int32
}

type ReadResponseUsers struct {
	Id          int32
	FullName    string
	Address     string
	PhoneNumber string
	Gender      string
}
