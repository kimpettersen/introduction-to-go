package main

import (
	"flag"
	"fmt"

	tm "github.com/buger/goterm"
)

func main() {
	tm.Clear()
	box := tm.NewBox(60|tm.PCT, 20, 0)

	str := `
Go Meetup
https://www.meetup.com/Go-Oslo-User-Group/

Golang Slack Group
Join #norway
https://invite.slack.golangbridge.org/

If you want to give a talk or help organizing
e-mail: %s
`
	email := flag.String("email", "", "Email address passed in as flag so it's not checked in on Github")
	flag.Parse()
	fmt.Fprint(box, fmt.Sprintf(str, *email))
	tm.Print(tm.MoveTo(box.String(), 20|tm.PCT, 20|tm.PCT))
	tm.Flush()
}