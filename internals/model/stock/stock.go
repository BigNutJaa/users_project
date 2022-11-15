package category

type Request struct {
	Name   string
	Detail string
}

type FitterUpdateCategory struct {
	Name         string
	Detail       string
	Id           int32
	NameUpdate   string
	DetailUpdate string
}

type UpdateResponseCategory struct {
	Name   string
	Detail string
	Id     int32
}
