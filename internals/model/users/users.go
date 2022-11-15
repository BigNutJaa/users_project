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

type FitterListUsers struct {
	FullName     string
	Phone_number string
	Gender       string
	Page         int64
	Per_page     int64
}

type ReadResponseUsers struct {
	Id          int32
	FullName    string
	Address     string
	PhoneNumber string
	Gender      string
}

type ListResponseUsers struct {
	Id          int32
	FullName    string
	Address     string
	PhoneNumber string
	Gender      string
}
