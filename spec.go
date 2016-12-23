package output

type Output interface {
	Text() string
}

// Service provides a communication channel to send information sequences.
type Service interface {
	Channel() chan Output
}
