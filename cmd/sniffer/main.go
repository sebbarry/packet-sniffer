package main


import (
    "fmt"
    "log"
    "os"
    "github.com/sebbarry/sniffer/capture/pcap"
)


func usage() {
    fmt.Println("--------------")
    fmt.Println("To use: ./sniffer <network interface>")
    fmt.Println("--------------")
}

func main() {

    //plumbing
    args := os.Args
    if len(args) <= 1 {
        usage()
        os.Exit(1)
    }

    nic := args[1]

    //init the handler for our sniffer
    src, err := pcap.Open(nic)

    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }

    defer src.Close()


    //connect to the network interface
    err = src.Activate()
    if err != nil {
        log.Fatal(err)
    }

    //switch the stream on
    for {
        buf, err := src.Capture()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Packet: ", buf)
    }

}








