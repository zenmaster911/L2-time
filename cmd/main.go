package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

// ntp0.ntp-servers.net
func main() {
	resp, err := ntp.Query("ntp.msk-ix.ru")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ntp querry error: %v\n", err)
		os.Exit(1)
	}
	if err := resp.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "ntp querry validation error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(resp.Time.Local().Format("2006-01-02 15:04:05"))

}
