package storeLibs

type Register struct {
	ID                     int
	QueueSize              int
	MaxQueueSize           int
	CurrentlyServicing     bool
	Money                  *Money
	Line			       *Queue
	TotalCustomersServiced int
}

func MakeRegister(i int, qs int, mqs int, ser bool) *Register {
	return &Register{i, qs, mqs, ser, &Money{0, 0}, MakeQueue(), 0}
}


