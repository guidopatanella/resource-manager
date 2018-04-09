package rmanager


import "sync"
import "fmt"
import "time"

var RM ResourceManager

// ResourceManager - main resource manager struct
type ResourceManager struct {
	mutex sync.Mutex
	// setting a fixed max number of resources
	Resources [1000]Resource
	RoutineManager RoutineManager
	Stats map[string]ResourceCallerStats
	lastResourceIndex int
}

// Init - initializes the resource manager
func (rm *ResourceManager) Init() {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()
	rm.lastResourceIndex = 0
	rm.Stats = make(map[string]ResourceCallerStats)
}

// AddResource - adds a new resource to the resouce manager
func (rm *ResourceManager) AddResource(r Resource) {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()
	rm.Resources[rm.lastResourceIndex] = r
	rm.lastResourceIndex++
	// init stats for resort
	resourceCallerStats := ResourceCallerStats{}
	resourceCallerStats.Init()
	rm.Stats[r.GetID()] = resourceCallerStats
}

// Start - starts a specified resource 
func (rm *ResourceManager) Start(resourceID string, heartbeat time.Duration) error {
	// retrieve caller for stats purposes
	// locking
	rm.mutex.Lock()
	defer rm.mutex.Unlock()
	// caller := GetCaller(3)
	for i:=0; i<rm.lastResourceIndex; i++ {
		if rm.Resources[i].GetID()==resourceID {
			return rm.Resources[i].Start(heartbeat)
		}
	}
	// not found
	// returns any resource implementer and an error
	return fmt.Errorf("unable to find resource [%s]", resourceID)
}

// IsAvailable - queries if a resource is available
func (rm *ResourceManager) IsAvailable(resourceID string) bool {
	// retrieve caller for stats purposes
	// locking
	rm.mutex.Lock()
	defer rm.mutex.Unlock()
	caller := GetCaller(3)
	for i:=0; i<rm.lastResourceIndex; i++ {
		if rm.Resources[i].GetID()==resourceID {
			available := rm.Resources[i].IsAvailable()
			// set status for resource
			stats := rm.Stats[resourceID]
			stats.Set(caller, available)
			rm.Stats[resourceID] = stats
			return available
		}
	}
	// not found
	// returns any resource implementer and an error
	fmt.Printf("ResourceManager.IsAvailable - unable to find resource %s\n", resourceID)
	return false
}

// ListResources - lists resources
func (rm *ResourceManager) ListResources() {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()
	for i:=0; i<rm.lastResourceIndex; i++ {
		fmt.Printf("resource [%s] available: %t\n", rm.Resources[i].GetID(), rm.Resources[i].IsAvailable())
	}
}

// ListStats - for each resource, lists its stats
func (rm *ResourceManager) ListStats() {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()
	for i,v := range rm.Stats {
		fmt.Printf("stats for resource [%s]: \n", i)
		v.PrintStats()
	}
}