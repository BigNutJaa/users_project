package token

type Request struct {
	User_name string
	Password  string
}

type ReadResponseToken struct {
	Id        int32
	User_name string
	Token     string
}

type FitterDeleteToken struct {
	Id int32
}

type DeleteResponseToken struct {
	Id int32
}
