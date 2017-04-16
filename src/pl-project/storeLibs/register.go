package storeLibs

type Register struct {
	ID                 int
	QueueSize          int
	MaxQueueSize       int
	CurrentlyServicing bool
}

const NUMOFREGISTERS int = 5
