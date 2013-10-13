// Package plumb provides a client side interface for plumbing
// messages to Plan 9 plumber(4).
package plumb

import (
	"code.google.com/p/goplan9/plan9/client"
	"sync"
)

var once sync.Once
var fsys *client.Fsys
var fsysErr error

func mountPlumb() {
	fsys, fsysErr = client.MountService("plumb")
}

// Open opens the plumb port named port using access mode omode.
func Open(port string, omode uint8) (*Port, error) {
	once.Do(mountPlumb)
	if fsysErr != nil {
		return nil, fsysErr
	}
	fid, err := fsys.Open(port, omode)
	return (*Port)(fid), err

}

// Send writes the message msg to the plumb port.
func (port *Port) Send(msg *Msg) error {
	b := packMsg(msg)
	fid := (*client.Fid)(port)
	n, err := fid.Write(b)
	if n != len(b) {
		return err
	}
	return nil
}

// Close closes the plumb port.
func (port *Port) Close() error {
	return (*client.Fid)(port).Close()
}
