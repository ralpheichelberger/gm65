package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"

	"bezahl.online/gm65/device"
	"github.com/radovskyb/watcher"
)

const WATCHER_POLL_TIME_MS = 100
const TAN_FILE = "/tmp/tan"

func readTANFile() string {
	b, err := ioutil.ReadFile(TAN_FILE)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = os.Remove(TAN_FILE)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(b)
}

func TANFileWatcher(scanner *device.Scanner) {
	w := watcher.New()
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Create)
	w.AddFilterHook(watcher.RegexFilterHook(regexp.MustCompile("^tan$"), false))

	go func() {
		for {
			select {
			case <-w.Event:
				scanner.SetCode(readTANFile())
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.Add("/tmp"); err != nil {
		log.Fatalln(err)
	}

	if err := w.Start(time.Millisecond * WATCHER_POLL_TIME_MS); err != nil {
		log.Fatalln(err)
	}
}
