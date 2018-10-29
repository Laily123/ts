package main

import (
	"flag"
	"strconv"
	"time"

	"github.com/deanishe/awgo"
)

const (
	layout = "2006-01-02 15:04:05"
)

var (
	wf  *aw.Workflow
	val string
)

func init() {
	flag.StringVar(&val, "v", "", "")
	flag.Parse()
	wf = aw.New()
}

func run() {
	ts, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		date, err1 := time.Parse(layout, val)
		if err1 != nil {
			wf.NewItem("can not transfer")
		} else {
			// 如果是时间字符串
			wf.NewItem(strconv.FormatInt(date.Unix(), 10))
		}
		wf.SendFeedback()
	} else {
		// 如果是时间戳
		t := time.Unix(ts, 0)
		wf.NewItem(t.Format(layout))
		wf.SendFeedback()
	}
}

func main() {
	wf.Run(run)
}
