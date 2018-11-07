package main

import (
	"fmt"
	"github.com/robfig/cron"
)

func main() {
	Test()
}

func Test() {
	c := cron.New()
	c.AddFunc("@every 1s", func() {
		fmt.Println("Every each second")
	})
	c.Start()

	inspect(c.Entries())

	var Test string
	fmt.Scan(&Test)

}
func inspect(entries []*cron.Entry) {

	for _, entry := range entries {
		entry.Job.Run();
	}
}
