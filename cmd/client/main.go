package main

import (
	"context"
	"fmt"
	c "github.com/mvanyushkin/go-calendar/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"os"
	"time"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial("127.0.0.1:8888", opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := c.NewCalendarClient(conn)
	nowTimestamp := time.Now()

	fmt.Println("Creating a new event for tomorrow...")
	createEvent(err, client, nowTimestamp.AddDate(0, 0, 1))

	fmt.Println("Creating a new event for a week...")
	createEvent(err, client, nowTimestamp.AddDate(0, 0, 7))

	fmt.Println("Creating a new event for a month...")
	respMonth := createEvent(err, client, nowTimestamp.AddDate(0, 0, 25))

	fmt.Println("Creating a new event for tomorrow again...")
	_, err = client.Create(context.Background(), &c.EventDto{
		Title:       "test title",
		Description: "test description",
		Time:        nowTimestamp.AddDate(0, 0, 1).Unix(),
	})

	if err != nil {
		fmt.Printf("occured exception %v it's ok\n", err.Error())
	}

	tomorrowsEvents, err := client.GetForDate(context.Background(), &c.DateRequest{
		Day: time.Now().AddDate(0, 0, 1).Unix(),
	})

	fmt.Println("Today's events:")
	for _, event := range tomorrowsEvents.Events {
		fmt.Printf("id %v title %v description %v \n", event.Id, event.Title, event.Description)
	}

	weeksEvents, err := client.GetForWeek(context.Background(), &c.Empty{})

	fmt.Println("Week's's events:")
	for _, event := range weeksEvents.Events {
		fmt.Printf("id %v title %v description %v \n", event.Id, event.Title, event.Description)
	}

	monthsEvents, err := client.GetForMonth(context.Background(), &c.Empty{})

	fmt.Println("Months's's events:")
	for _, event := range monthsEvents.Events {
		fmt.Printf("id %v title %v description %v \n", event.Id, event.Title, event.Description)
	}

	for _, event := range monthsEvents.Events {
		fmt.Printf("Modifying event %v \n", event.Id)
		event.Title = "modified title"
		_, err = client.Update(context.Background(), &c.EventDto{
			Id:          event.Id,
			Title:       event.Title,
			Description: event.Description,
			Time:        event.Time,
		})

		if err != nil {
			fmt.Printf("occured exception %v \n", err.Error())
			os.Exit(-1)
		}
	}

	_, err = client.Remove(context.Background(), &c.EventDto{
		Id: respMonth.Id,
	})

	if err != nil {
		fmt.Printf("occured exception %v it's ok\n", err.Error())
	}
}

func createEvent(err error, client c.CalendarClient, nowTimestamp time.Time) *c.CreateResponseDto {
	resp, err := client.Create(context.Background(), &c.EventDto{
		Title:       "test title",
		Description: "test description",
		Time:        nowTimestamp.Unix(),
	})

	if err != nil {
		fmt.Printf("occured exception %v \n", err.Error())
		os.Exit(-1)
	}

	fmt.Printf("Event created, id = %v. \n", resp.Id)
	return resp
}
