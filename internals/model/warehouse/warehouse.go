package warehouse

type Request struct {
	Name   string
	Detail string
}

type FitterDeleteWarehouse struct {
	Name   string
	Detail string
	Id     int32
}

type DeleteResponseWarehouse struct {
	Name   string
	Detail string
	Id     int32
}
