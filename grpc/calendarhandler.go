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
	return g.getByCriteria(func(x entities.Event) bool {
		return x.Time.Unix() >= time.Now().Unix() && x.Time.Unix() <= x.Time.AddDate(0, 0, 1).Unix()
	})
}

func (g CalendarHandler) GetForWeek(context.Context, *Empty) (*EventsResponse, error) {
	return g.getByCriteria(func(x entities.Event) bool {
		return x.Time.Unix() >= time.Now().Unix() && x.Time.Unix() <= x.Time.AddDate(0, 0, 7).Unix()
	})
}

func (g CalendarHandler) GetForMonth(context.Context, *Empty) (*EventsResponse, error) {
	return g.getByCriteria(func(x entities.Event) bool {
		return x.Time.Unix() >= time.Now().Unix() && x.Time.Unix() <= x.Time.AddDate(0, 0, 30).Unix()
	})
}

func (g CalendarHandler) getByCriteria(criteria func(x entities.Event) bool) (*EventsResponse, error) {
	events, err := g.Calendar.List()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Something went wrong")
	}

	var values = funk.Filter(events, criteria).([]entities.Event)
	var items = funk.Map(values, func(x entities.Event) *EventDto {
		item := &EventDto{
			Id:          int32(x.Id),
			Title:       string(x.Title),
			Description: string(x.Description),
			Time:        x.Time.Unix(),
		}
		return item
	}).([]*EventDto)

	return &EventsResponse{
		Events: items,
	}, nil
}
