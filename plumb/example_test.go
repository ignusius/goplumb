package plumb

import (
	"code.google.com/p/goplan9/plan9"
)

func ExampleSend() {
	port, err := Open("send", plan9.OWRITE)
	if err != nil {
		panic(err)
	}
	defer port.Close()

	msg := &Msg{
		Src:  "goplumb",
		Dst:  "edit",
		WDir: "/home/fhs",
		Kind: "text",
		Data: []byte("/etc/passwd:9"),
	}
	port.Send(msg)
}
