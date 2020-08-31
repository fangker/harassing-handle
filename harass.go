package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Harass struct {
	config    *HarassingConfig
	total     int
	doneCount int
	errCount  int
	sync.Mutex
}

func NewHarass(config *HarassingConfig) *Harass {
	config.Info()
	return &Harass{config: config, total: len(config.List)}
}

type submitInfo struct {
	url  string
	name string
}

func getSubmitInfo(s submitInfo) string {
	return fmt.Sprintf("%+v", s)
}

func (h *Harass) Do() {
	var d = make(chan submitInfo)
	var e = make(chan submitInfo)
	var quitTag = make(chan int, h.total)
	for i := 0; i < h.total; i++ {
		i := i
		go func() {
			client := &http.Client{}
			item := h.config.List[i]
			resp, err := client.Do(ParseWebConfigItem(item))
			if err != nil {
				log.Fatalln(err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			fmt.Sprintf("%s", body)
			h.Lock()
			//fmt.Printf(string(body), err)
			if err != nil {
				h.errCount++
				e <- submitInfo{name: item.Name, url: item.RequestURL}
			} else {
				h.doneCount++
				d <- submitInfo{name: item.Name, url: item.RequestURL}
			}
			quitTag <- 1
			h.Unlock()
		}()
	}
ForEnd:
	for {
		select {
		case v := <-d:
			fmt.Println("【Submit】:", getSubmitInfo(v))
		case v := <-e:
			fmt.Println("【Error】:", getSubmitInfo(v))
		case <-quitTag:
			if h.errCount+h.doneCount == h.total {
				fmt.Printf("【Result】submit ended |  ok: %d err: %d \n", h.doneCount, h.errCount)
				break ForEnd
			}
		case <-time.After(time.Second * 100):
			fmt.Println("Time Out")
		}
	}
}
