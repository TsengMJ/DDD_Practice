package department

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}
