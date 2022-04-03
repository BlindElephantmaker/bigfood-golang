package cafe

type Cafe struct {
	Id Id `json:"id" example:"uuid"`
}

func New() *Cafe {
	return &Cafe{
		Id: newId(),
	}
}
