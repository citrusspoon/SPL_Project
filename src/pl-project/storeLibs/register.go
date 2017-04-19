package storeLibs

type Register struct {
	ID                 int
	QueueSize          int
	MaxQueueSize       int
	CurrentlyServicing bool
	Money              *Money
	Line			   *Queue
}

func MakeRegister() *Register {
	return &Register{0, 0, 0, false, &Money{0, 0}, MakeQueue()}
}


