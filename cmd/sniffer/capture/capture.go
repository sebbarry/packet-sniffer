package capture

type Handler interface {
    SetMTU(mtu int) error
    SetPromiscMode(promisc bool) error
    SetMonitorMode(monitor bool) error
    Activate() error
    Capture() ([]byte, error)
    Inject(buf []byte) error
    Close()
}
