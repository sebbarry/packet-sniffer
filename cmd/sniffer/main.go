package main

import (
    "fmt"
    "os"
    "net"
)



// interfaces and structs


//Network 

type NetworkSource struct {
    Default bool
    SrcNet  *net.IPNet
    DstNet  *net.IPNet
    Gateway net.IP
    Iface   *net.Interface
}


//ERRORS

type error interface {
    PrintError() string
}
type inputError struct {
    error string
}
type connectionError struct {
    error string
}
func (e *inputError) PrintError() {
    fmt.Println(e.error)
}
func (e *connectionError) PrintError() {
    fmt.Println(e.error)
}

//helpers.

func handleInput(arguments []string) (*inputError) {
    if len(arguments) - 1 <= 0 {
        return makeNetworkError()
    }
    return nil
}

func makeNetworkError() (*inputError) {
    return &inputError { error: "Invalid input." }
}
func makeConnectionError(con string) (*connectionError) {
    err := "Could not connect to: " + con
    return &connectionError{ error: err }
}




//main 

func main() {
    //check for params
    arguments := os.Args
    err := handleInput(arguments)
    if err != nil {
        err.PrintError()
        os.Exit(1)
    }

    //open the connection 
    src, openerr := Open(os.Args[1])
    if openerr != nil { 
        openerr.PrintError() 
        os.Exit(1)
    } else {
        fmt.Println(src)
    }


    defer Close(src) 


    //establish the src's stream (switch on)

    //start capturing the packets

}



//handlers 

//open the open a connection to the nic
func Open(con string) (*NetworkSource, *connectionError) {
    if len(con) == 0 {
        return nil, makeConnectionError(con)
    }
    return &NetworkSource{}, nil
}

func Capture() {}
func Close(*NetworkSource) {}






