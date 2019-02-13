package schedule

import (
	"fmt"
	"github.com/robfig/cron"
)

func init() {
	c := cron.New()
	_ = c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	_ = c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	_ = c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	_ = c.AddFunc("@every 5s", func() { fmt.Println("Every five seconds") })
	c.Start()
}
