// 8g try.go && 8l try.8 && ./8.out

package main

import (
	"bitbucket.org/fhs/goplumb/plumb"
	"goplan9.googlecode.com/hg/plan9"
)

func main() {
	port, err := plumb.Open("send", plan9.OWRITE)
	if err != nil {
		panic(err)
	}
	defer port.Close()

	msg := &plumb.Msg{
		Src:  "goplumb",
		Dst:  "edit",
		WDir: "/home/fhs",
		Kind: "text",
		Attr: map[string]string{},
		Data: []byte("/etc/passwd:9"),
	}
	port.Send(msg)
}
