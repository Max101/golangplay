package main

import (
	"encoding/json"
	"fmt"
	r "mas/db"
)

func main() {

	guestRepo := r.GuestRepo{}

	guestFilter := r.NewGuestFilter()
	guestFilter.Country = "us"

	guests, totals, _ := guestRepo.Get(guestFilter)
	//fmt.Println(guests)

	b, _ := json.Marshal(guests)

	fmt.Printf("Totals: %d \n", totals)
	fmt.Println(string(b))

	//for _, val := range all {
	//	fmt.Println(string(json.Marshal(val)))
	//}

}

//// package main
//
//// import (
//// 	"fmt"
//
//// 	"./repository"
//// )
//
//// type Salutation struct {
//// 	name     string
//// 	greeting string
//// 	times    int
//// }
//
//// type Salutations []Salutation
//
//// type Printer func(string)
//
//// func (salutations Salutations) Greet(do Printer, isFormal bool, times int) {
//// 	for _, s := range salutations {
//// 		fmt.Println(s.name)
//// 	}
//// }
//
//// func (salutations Salutations) ChannelGreeter(c chan Salutation) {
//// 	for _, s := range salutations {
//// 		c <- s
//// 	}
//// 	close(c)
//// }
//
//// func main() {
//// 	repository.InitDB("root", "", "127.0.0.1")
//
//// 	salutations := Salutations{
//// 		{"lol", "ek", 1},
//// 		{"kal", "el", 2},
//// 		{"super", "mman", 3},
//// 	}
//
//// 	c := make(chan Salutation)
//// 	c2 := make(chan Salutation)
//
//// 	go salutations.ChannelGreeter(c)
//// 	go salutations.ChannelGreeter(c2)
//
//// 	for {
//// 		select {
//// 		case s, ok := <-c:
//// 			if ok {
//// 				fmt.Println(s, 1)
//// 			} else {
//// 				return
//// 			}
//
//// 		case s, ok := <-c:
//// 			if ok {
//// 				fmt.Println(s, 2)
//// 			} else {
//// 				return
//// 			}
//
//// 		default:
//// 			fmt.Println("default")
//// 		}
//
//// 	}
//// }
//
//package main
//
//import (
//	"fmt"
//
//	"./repository"
//)
//
//type Salutation struct {
//	name     string
//	greeting string
//	times    int
//}
//
//type Salutations []Salutation
//
//type Printer func(string)
//
//func (salutations Salutations) Greet(do Printer, isFormal bool, times int) {
//	for _, s := range salutations {
//		fmt.Println(s.name)
//	}
//}
//
//func (salutations Salutations) ChannelGreeter(c chan Salutation) {
//	for _, s := range salutations {
//		c <- s
//	}
//	close(c)
//}
//
//func main() {
//	repository.InitDB("root", "", "127.0.0.1")
//	doSequential()
//	doConcurrent()
//}
//
//func doSequential() {
//	allGuests, _ := repository.GetAllGuests()
//	allBooks, _ := repository.GetAll()
//
//	fmt.Println("LenBooks", len(allBooks))
//	fmt.Println("LenGuests", len(allGuests))
//
//}
//
//func doConcurrent() {
//	c := make(chan bool, 2)
//
//	go getBooks(c)
//	go func() {
//		allGuests, _ := repository.GetAllGuests()
//
//		fmt.Println("LenGuests", len(allGuests))
//
//		c <- true
//	}()
//
//	<-c
//	<-c
//}
//
//func getBooks(c chan){
//	allBooks, _ := repository.GetAll()
//
//	fmt.Println("LenBooks", len(allBooks))
//
//	c <- true
//}
