package delta

import (
	"context"
	"github.com/BigNutJaa/user-service/internals/entity"
	model "github.com/BigNutJaa/user-service/internals/model/delta"
	"log"
	"sync"
)

func (s *DeltaService) Create(ctx context.Context, request *model.Request) (int, error) {
	c := make(chan int)
	var roundSum int
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go s.buildCreate(request, c, wg)
	roundSum = roundSum + <-c
	log.Println("Round 1: ", roundSum)
	go s.buildCreate3(request, c, wg)
	roundSum = roundSum + <-c
	log.Println("Round 2 : ", roundSum)
	wg.Wait()
	log.Println("Wait Group")

	go s.buildCreate2(roundSum, request, c)
	roundSum = roundSum + <-c
	log.Println("Round 3: ", roundSum)
	log.Println("END Process")
	//result := <-c
	//log.Println("ID :", result)
	return 1, nil
}
func (s *DeltaService) buildCreate(request *model.Request, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var i int
	for i = 1; i < 2000; i++ {
		input := &entity.Delta{
			MovieName: request.MovieName,
			Date:      request.Date,
			Time:      request.Time,
			CinemaNo:  request.CinemaNo,
		}
		s.repository.Create(input)
	}
	c <- i

}
func (s *DeltaService) buildCreate3(request *model.Request, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var i int
	for i = 1; i < 2000; i++ {
		input := &entity.Delta{
			MovieName: request.MovieName,
			Date:      request.Date,
			Time:      request.Time,
			CinemaNo:  request.CinemaNo,
		}
		s.repository.Create(input)
	}
	c <- i

}
func (s *DeltaService) buildCreate2(num int, request *model.Request, c chan int) {
	log.Println("START buildCreate2")
	var i int
	for i = num; i < 6000; i++ {
		input := &entity.Delta{
			MovieName: request.MovieName,
			Date:      request.Date,
			Time:      request.Time,
			CinemaNo:  request.CinemaNo,
		}
		s.repository.Create(input)
	}
	c <- i

}

//func (s *DeltaService) Create(ctx context.Context, request *model.Request) (int, error) {
//	//c := make(chan int)
//	wg := new(sync.WaitGroup)
//	wg.Add(1)
//
//	go s.buildCreate(request, wg)
//	wg.Wait()
//	//result := <-c
//	//log.Println("ID :", result)
//	return 1, nil
//}
//func (s *DeltaService) buildCreate(request *model.Request, wg *sync.WaitGroup) {
//	defer wg.Done()
//	var i int
//	for i = 1; i < 2000; i++ {
//		input := &entity.Delta{
//			MovieName: request.MovieName,
//			Date:      request.Date,
//			Time:      request.Time,
//			CinemaNo:  request.CinemaNo,
//		}
//		s.repository.Create(input)
//	}
//
//}
