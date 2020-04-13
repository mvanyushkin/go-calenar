package messages

type WorkItem struct {
	Message      MessageDto
	DoneCallback func()
}
