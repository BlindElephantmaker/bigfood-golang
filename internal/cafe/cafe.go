package cafe

type Cafe struct {
	Id
}

func New() *Cafe {
	return &Cafe{
		Id: NewId(),
	}
}
