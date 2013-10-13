package plumb

import (
	"code.google.com/p/goplan9/plan9/client"
	"strconv"
)

// Port represents an open plumber port.
type Port client.Fid

// Msg is a message that can be send to an open plumber port.
type Msg struct {
	Src  string            // source application
	Dst  string            // destination port
	WDir string            // working directory
	Kind string            // type of data
	Attr map[string]string // attributes
	Data []byte
}

func packMsg(msg *Msg) []byte {
	attr := ""
	for name, val := range msg.Attr {
		if attr != "" {
			attr += " "
		}
		attr += name + "=" + val
	}
	return append([]byte(msg.Src+"\n"+
		msg.Dst+"\n"+
		msg.WDir+"\n"+
		msg.Kind+"\n"+
		attr+"\n"+
		strconv.Itoa(len(msg.Data))+"\n"),
		msg.Data...)
}
