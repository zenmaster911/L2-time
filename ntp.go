package L2time

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	resp, err := ntp.Query("ntp0.ntp-servers.net")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ntp querry error: %v\n", err)
		os.Exit(1)
	}
	if err := resp.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "ntp querry validation error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(time.Now().Add(resp.ClockOffset))

}
