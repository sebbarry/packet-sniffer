package errors




import (
    "fmt"
)


//Error types
type error interface {
    PrintError() string
}
type InputError struct {
    error string
}
type ConnectionError struct {
    error string
}



//Generic functions
func (e *InputError) PrintError() {
    fmt.Println(e.error)
}
func (e *ConnectionError) PrintError() {
    fmt.Println(e.error)
}


//Handlers 
func MakeNetworkError() (*InputError) {
    return &InputError { error: "Invalid input." }
}
func MakeConnectionError(con string) (*ConnectionError) {
    err := "Could not connect to: " + con
    return &ConnectionError{ error: err }
}
