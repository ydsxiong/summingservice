package datastore

type Store interface {
	Save(DataMessage) error
	FetchAll() ([]DataMessage, error)
}

type DataMessage struct {
	Name   string
	Input  []int64
	Output int64
}
