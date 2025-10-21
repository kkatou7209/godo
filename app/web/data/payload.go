package data

type Status bool

const (
	StatusSuccess Status = true
	StatusFail    Status = false
)

type Payload[T any] struct {
	Status Status  			 `json:"status"`
	Message string 			 `json:"message"`
	Data *T           		 `json:"data"`
	Errors map[string]string `json:"errors"`
}

type PayloadOption[T any] func(p *Payload[T])

func NewPayload[T any](status Status, data T) *Payload[T] {

	p := &Payload[T]{
		Status: status,
		Message: "",
		Data: &data,
		Errors: nil,
	}

	return p
}

func (p *Payload[T]) WithMessage(msg string) *Payload[T] {

	p.Message = msg

	return p
}

func (p *Payload[T]) WithErrors(key string, msg string) *Payload[T] {

	if p.Errors == nil {
		p.Errors = map[string]string{}
	}

	p.Errors[key] = msg
	
	return p
}