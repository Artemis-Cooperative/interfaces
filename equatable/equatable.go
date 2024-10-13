package equatable

type equatable interface {
	Equals(equatable) bool
}