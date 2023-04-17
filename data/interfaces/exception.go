package interfaces

type IException interface {
	ToString() string
	ToError() error
	Business() string
	Tech() string
}
