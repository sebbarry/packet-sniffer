package pcap

import (
    "fmt"
    "unsafe"
)

// #cgo LDFLAGS: -lpcap
// #include <stdint.h>
// #include <stdlib.h>
// #include <pcap.h>
import "C"

type Handler struct {
    Device string
    pcap *C.pcap_t
}

//creates a new network capture handle from the interface
func Open(device_name string) (*Handler, error){

    handle := &Handler{ Device: device_name }      //make a new handler
    dev_str := C.CString(device_name)              //convert to C string
    defer C.free(unsafe.Pointer(dev_str))          //defer freeing memory of dev_str
    err_str := (*C.char)(C.calloc(256, 1))         //allocate 256 memory slots to the buffer
    defer C.free(unsafe.Pointer(err_str))          //free memory space when closed loop
    handle.pcap = C.pcap_create(dev_str, err_str)  //make a new pcap value
    if handle == nil {
        return nil, fmt.Errorf(
            "Could not open device: %s", C.GoString(err_str),
        )
    }
    return handle, nil
}

func (h *Handler) Close() {
    C.pcap_close(h.pcap);
}


func (h *Handler) Activate() error {
    err := C.pcap_activate(h.pcap)
    if err < 0 {
        C.pcap_close(h.pcap)
        return nil
    }
    return nil
}



func (h *Handler) Capture() ([]byte, error)  {

    var buf *C.uint8_t
    var pkt_hdr *C.struct_pcap_pkthdr

    for {
        err := C.pcap_next_ex(h.pcap, &pkt_hdr, &buf)
        switch err {
        case -2:
            return nil, nil
        case -1:
            return nil, fmt.Errorf(
                "Could not read packet: %d", err,
            )
        case 0:
            continue
        case 1:
            return C.GoBytes(unsafe.Pointer(buf),
                C.int(pkt_hdr.len)), nil
        }
        return nil, fmt.Errorf("Something happened.")
    }

}
