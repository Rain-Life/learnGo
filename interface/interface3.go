package learn_interface

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type WriteCloser interface {
	Writer
	Closer
}

type Device struct {

}

func (d *Device) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (d *Device) Close() error {
	return nil
}


