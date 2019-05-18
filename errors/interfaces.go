package errors

type VKErrors interface {
	GetCode() int
	GetDescription() string
}
