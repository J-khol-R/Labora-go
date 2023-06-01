package models

type Message struct {
	Mensaje string `json:"message"`
}

func (m *Message) Ok() {
	m.Mensaje = "Transaccion exitosa"
}
