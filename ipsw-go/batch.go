package ipsw_go

import (
	"github.com/cavaliergopher/grab/v3"
	"github.com/dustin/go-humanize"
	"github.com/spf13/cast"
	"ipsw-go/logger"
	"os"
	"sync"
	"time"
)

func doRequest(c *grab.Client, req *grab.Request) (resp *grab.Response) {
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()
	resp = c.Do(req)
	_ = os.Chmod(resp.Filename, 0777)
	var lastBytesComplete int64 = 0
	for {
		select {
		case <-t.C:
			logger.Infof("file: %s transferred %v / %v (%.2f%%) speed: %v/s",
				resp.Filename,
				humanize.Bytes(cast.ToUint64(resp.BytesComplete())),
				humanize.Bytes(cast.ToUint64(resp.Size())),
				100*resp.Progress(),
				humanize.Bytes(cast.ToUint64(resp.BytesComplete()-lastBytesComplete)/5))
			lastBytesComplete = resp.BytesComplete()

		case <-resp.Done:
			return
		}
	}
}

func doBatch(c *grab.Client, workers int, reqs []*grab.Request) <-chan *grab.Response {
	tasksNum := len(reqs)
	if workers < 1 {
		workers = 1
	}

	reqch := make(chan *grab.Request, tasksNum)

	respch := make(chan *grab.Response, tasksNum)
	wg := sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for req := range reqch {
				respch <- doRequest(c, req)
			}
			wg.Done()
		}()
	}

	// queue requests
	go func() {
		for _, req := range reqs {
			reqch <- req
		}
		close(reqch)
		wg.Wait()
		close(respch)
	}()

	return respch
}
