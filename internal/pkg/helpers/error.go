package helpers

type FamlyErr struct {
	Code    int
	Message string
	Error   error
}

var NoFamErr = FamlyErr{}
