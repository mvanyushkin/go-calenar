package main

import (
	"fmt"
	calendar2 "github.com/mvanyushkin/go-calendar/calendar"
	"time"
)
import store "github.com/mvanyushkin/go-calendar/calendar/store"

func main() {
	s := store.NewInMemoryEventStore()
	calendar := calendar2.NewCalendar(s)

	println("Creating events...")
	event1, _ := calendar.CreateEvent("TestEvent 1", "Description 1", time.Now())
	event2, _ := calendar.CreateEvent("TestEvent 2", "Description 2", time.Now())
	println("Done.")

	println("Iterating by events..")
	items, _ := calendar.List()
	for _, item := range items {
		fmt.Printf("Id: %v, Title %v \n", item.Id, item.Title)
	}

	println("Done.")

	println("Get by id ")

	evt, _ := calendar.Get(event1.Id)
	fmt.Printf("Given the event: %v \n", evt.Id)

	println("Done.")

	println("Removing events...")

	calendar.Remove(event1)
	calendar.Remove(event2)

	items, _ = calendar.List()
	fmt.Printf("Items left %v \n", len(items))
	println("Bye!")
}
