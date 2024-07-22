package exception

type BaseError interface {
	Error() string
	Code() string
	HttpStatusCode() int
}
