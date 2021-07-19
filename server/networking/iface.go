package networking

type WSWriter interface {
	WriteJSON(v interface{}) error
}
