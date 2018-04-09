package rmanager

import "sync"
import "runtime"
import "fmt"

// ResourceCallerStats - defines a struct for stats collection for resources
type ResourceCallerStats struct {
	mutex sync.Mutex
	SuccessMap map[string]int64
	FailedMap map[string]int64
}

// Init - initializes the resource stats collection struct 
func (rcs *ResourceCallerStats) Init(){
	rcs.mutex.Lock()
	defer rcs.mutex.Unlock()
	rcs.SuccessMap = make(map[string]int64)
	rcs.FailedMap = make(map[string]int64)
}

// Set - set stats for a specific caller
func (rcs *ResourceCallerStats) Set(caller string, available bool) {
	rcs.mutex.Lock()
	defer rcs.mutex.Unlock()
	if _,ok := rcs.SuccessMap[caller]; !ok {
		rcs.SuccessMap[caller] = 0
	}
	if _,ok := rcs.FailedMap[caller]; !ok {
		rcs.FailedMap[caller] = 0
	}
	if available {
		rcs.SuccessMap[caller]++
	} else {
		rcs.FailedMap[caller]++
	}
}

// PrintStats - prints the stats for successful and failed resource queries
func (rcs *ResourceCallerStats) PrintStats() {
	rcs.mutex.Lock()
	defer rcs.mutex.Unlock()
	fmt.Println("successful requests:")
	for i, count := range rcs.SuccessMap {
		fmt.Printf(" - [%s]: %v\n", i, count)
	}
	fmt.Println("failed requests:")
	for i, count := range rcs.FailedMap {
		fmt.Printf(" - [%s]: %v\n", i, count)
	}
}


// GetCaller - support function which points to the caller of a function (used as an indicator of stats)
// level should be 3 in general
func GetCaller(level int) string {
    fpcs := make([]uintptr, 1)
    n := runtime.Callers(level, fpcs)
    if n == 0 {
       return ""
    }
    fun := runtime.FuncForPC(fpcs[0]-1)
    if fun == nil {
      return ""
    }
    return fun.Name()
}
