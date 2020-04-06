package grpc

import (
	"context"
	"github.com/mvanyushkin/go-calendar/calendar"
	"github.com/mvanyushkin/go-calendar/calendar/entities"
	"github.com/thoas/go-funk"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type CalendarHandler struct {
	Calendar *calendar.Calendar
}

func (g CalendarHandler) Create(ctx context.Context, e *CreateEventRequestDto) (*CreateResponseDto, error) {
	title := entities.Title(e.Data.Title)
	desc := entities.Description(e.Data.Description)
	time := time.Unix(e.Data.Time, 0)
	evt, err := g.Calendar.CreateEvent(title, desc, time)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Something went wrong")
	}

	return &CreateResponseDto{
		Id: int32(evt.Id),
	}, nil
}

func (g CalendarHandler) Update(ctx context.Context, e *UpdateEventRequestDto) (*Empty, error) {
	title := entities.Title(e.Data.Title)
	desc := entities.Description(e.Data.Description)
	time := time.Unix(e.Data.Time, 0)
	err := g.Calendar.Update(entities.Event{
		Id:          entities.Id(e.Id),
		Title:       title,
		Description: desc,
		Time:        time,
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Something went wrong")
	}

	return &Empty{}, nil
}

func (g CalendarHandler) Remove(ctx context.Context, r *RemoveEventRequestDto) (*Empty, error) {
	evt, err := g.Calendar.Get(entities.Id(r.Id))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Something went wrong")
	}

	err = g.Calendar.Remove(evt)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Something went wrong")
	}
	return &Empty{}, nil
}

func (g CalendarHandler) GetForDate(context.Context, *Empty) (*EventsResponse, error) {
	events, err := g.Calendar.List()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Something went wrong")
	}

	var values = funk.Filter(events, func(x entities.Event) bool {
		return x.Time.Unix() >= time.Now().Unix() && x.Time.Unix() <= x.Time.AddDate(0, 0, 1).Unix()
	}).([]entities.Event)

	var items = funk.Map(values, func(x entities.Event) *CreateEventRequestDto {
		 s := CreateEventRequestDto{
			 Data:                 ,
			 XXX_NoUnkeyedLiteral: struct{}{},
			 XXX_unrecognized:     nil,
			 XXX_sizecache:        0,
		 }
	}).([]entities.Event)

	return &EventsResponse{
		Events: ,
	}, nil
}

func (g CalendarHandler) GetForWeek(context.Context, *Empty) (*EventsResponse, error) {
	panic("implement me")
}

func (g CalendarHandler) GetForMonth(context.Context, *Empty) (*EventsResponse, error) {
	panic("implement me")
}
