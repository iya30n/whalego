package Proxy

type Proxy struct {
	Link   string
	Server string
	Port   int32
	Secret string
}

func New() *Proxy {
	return &Proxy{}
}