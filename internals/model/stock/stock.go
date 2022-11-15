package stock

type Request struct {
	Name   string
	Detail string
	Qty    int32
}

type FitterReadStock struct {
	Name   string
	Detail string
	Id     int32
}

type ReadResponseStock struct {
	Name   string
	Detail string
	Qty    int32
	Id     int32
}

type FitterDeleteStock struct {
	Id int32
}

type DeleteResponseStock struct {
	Name   string
	Detail string
	Qty    int32
	Id     int32
}

type FitterUpdateStock struct {
	Name      string
	Detail    string
	Qty       int32
	Id        int32
	QtyUpdate int32
}

type UpdateResponseStock struct {
	Name   string
	Detail string
	Qty    int32
	Id     int32
}

type FitterListStock struct {
	Name     string
	Detail   string
	Page     int64
	Per_page int64
}

type ListResponseStock struct {
	Id     int32
	Name   string
	Detail string
	Qty    int32
}
