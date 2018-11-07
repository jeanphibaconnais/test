package testmodules

import (
	"fmt"
	"github.com/robfig/cron"
)

func main() {

}

func Test() {
	c := cron.New()
	c.AddFunc("0 0 10 * * *", func() {
		fmt.Println("Every ten secondes")
	})

}
