package stock

type Request struct {
	movieName string
	date      string
	time      string
	cinemaNo  string
}

type FitterReadStock struct {
	movieName string
	date      string
	time      string
	cinemaNo  string
	Id        int32
}

type ReadResponseStock struct {
	movieName string
	date      string
	time      string
	cinemaNo  string
	Id        int32
}

type FitterDeleteStock struct {
	Id int32
}

type DeleteResponseStock struct {
	movieName string
	date      string
	time      string
	cinemaNo  string
	Id        int32
}

type FitterUpdateStock struct {
	movieName      string
	date           string
	time           string
	cinemaNo       string
	Id             int32
	dateUpdate     string
	timeUpdate     string
	cinemaNoUpdate string
}

type UpdateResponseStock struct {
	movieName string
	date      string
	time      string
	cinemaNo  string
	Id        int32
}

type FitterListStock struct {
	movieName string
	date      string
	time      string
	cinemaNo  string
	Page      int64
	Per_page  int64
}

type ListResponseStock struct {
	movieName string
	date      string
	time      string
	cinemaNo  string
}
