// Client library for plumbing messages to plumber(4).
package plumb

import (
	"os"
	"once"
	"goplan9.googlecode.com/hg/plan9/client"
)

var fsys *client.Fsys
var fsysErr os.Error


func mountPlumb() {
	fsys, fsysErr = client.MountService("plumb")
}

// Open opens the plumb port names plumb using access mode omode.
func Open(port string, omode uint8) (*Port, os.Error) {
	once.Do(mountPlumb)
	if fsysErr != nil {
		return nil, fsysErr
	}
	fid, err := fsys.Open(port, omode)
	return (*Port)(fid), err

}

// Send write the message msg to the plumb port.
func (port *Port) Send(msg *Msg) os.Error {
	b := packMsg(msg)
	fid := (*client.Fid)(port)
	n, err := fid.Write(b)
	if n != len(b) {
		return err
	}
	return nil
}

// Close closes the plumb port.
func (port *Port) Close() os.Error {
	return (*client.Fid)(port).Close()
}
