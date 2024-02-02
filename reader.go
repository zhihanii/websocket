package websocket

type Reader interface {
	ReadLine() ([]byte, bool, error)
}
