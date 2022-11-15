package movie

type Request struct {
	MovieName string
	Date      string
	Time      string
	CinemaNo  string
}

type FitterReadMovie struct {
	MovieName string
	Date      string
	Time      string
	CinemaNo  string
	Id        int32
}

type ReadResponseMovie struct {
	MovieName string
	Date      string
	Time      string
	CinemaNo  string
	Id        int32
}

type FitterDeleteMovie struct {
	Id int32
}

type DeleteResponseMovie struct {
	MovieName string
	Date      string
	Time      string
	CinemaNo  string
	Id        int32
}

type FitterUpdateMovie struct {
	MovieName      string
	Date           string
	Time           string
	CinemaNo       string
	Id             int32
	DateUpdate     string
	TimeUpdate     string
	CinemaNoUpdate string
}

type UpdateResponseMovie struct {
	MovieName string
	Date      string
	Time      string
	CinemaNo  string
	Id        int32
}

type FitterListMovie struct {
	MovieName string
	Date      string
	Time      string
	CinemaNo  string
	Page      int64
	Per_page  int64
}

type ListResponseMovie struct {
	MovieName string
	Date      string
	Time      string
	CinemaNo  string
}
