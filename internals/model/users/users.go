package users

type Request struct {
	User_name  string
	Password   string
	First_name string
	Last_name  string
	Email      string
}

type FitterReadUsers struct {
	User_name  string
	First_name string
	Email      string
	Id         int32
}

type FitterListUsers struct {
	User_name  string
	First_name string
	Email      string
	Page       int64
	Per_page   int64
}

type ReadResponseUsers struct {
	Id         int32
	User_name  string
	Password   string
	First_name string
	Last_name  string
	Email      string
	Role_code  string
}

type ListResponseUsers struct {
	Id         int32
	User_name  string
	First_name string
	Email      string
}

type FitterDeleteUsers struct {
	Id        int32
	User_name string
}

type DeleteResponseUsers struct {
	User_name  string
	First_name string
	Last_name  string
	Id         int32
}

type FitterUpdateUsers struct {
	User_name        string
	Id               int32
	PasswordUpdate   string
	First_nameUpdate string
	Last_nameUpdate  string
	EmailUpdate      string
	Role_codeUpdate  string
}

type UpdateResponseUsers struct {
	User_name  string
	Password   string
	First_name string
	Last_name  string
	Email      string
	Role_code  string
}
