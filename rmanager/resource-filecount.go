package rmanager

import (
	"io/ioutil"
	"sync"
	"time"
)

type FileCountResource struct {
	ID		  			string
	Path      			string
	Limit     			int
	available 			bool
	heartbeat 			int
	mutex     			sync.RWMutex
	tickerChan 			<-chan time.Time
}

func (r FileCountResource) GetID() string {
	return r.ID
}

func (r *FileCountResource) Start(heartbeatMilliseconds int) {
	r.tickerChan = time.NewTicker(time.Millisecond * time.Duration(heartbeatMilliseconds)).C
	r.checkAvailable()
	go r.collectData()
}

func (r FileCountResource) GetProgress() float32{
	files, _ := ioutil.ReadDir("/temp")
	if len(files) < r.Limit {
		return float32((len(files)/r.Limit) * 100)
	} else {
		return float32(100)
	}
}

func (r *FileCountResource) IsAvailable() bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.checkAvailable()
}

func (r FileCountResource) collectData() {
	for {
	 <- r.tickerChan
		r.mutex.Lock()
		r.checkAvailable()
		r.mutex.Unlock()
	}
}

func (r *FileCountResource) checkAvailable() bool {
	files, _ := ioutil.ReadDir("/temp")
	// validate if resource is available
	if len(files) > r.Limit {	
		r.available = false
	} else {
		r.available = true
	}
	return r.available
}
