package rmanager

import "testing"
import "time"
import "fmt"

type mokResource struct {
    id string
    progress float32
    up bool
}

func (mr *mokResource) SetID(id string) {
    mr.id = id
}

func (mr *mokResource) GetID() string{
    return mr.id
}

func (mr *mokResource) SetProgress(p float32){
    mr.progress = p
}

func (mr *mokResource) GetProgress() float32{
    return mr.progress
}

func (mr *mokResource) IsAvailable() bool{
    return mr.up
}

func (mr *mokResource) Start(heartbeat time.Duration) error{
    mr.up = true
    return nil
}

//TestResourceManager - creates a resource manager and an internal resource
func TestResourceManager(t *testing.T) {
    // creates a resource manager
    rm := new(ResourceManager)
    rm.Init()
    // creates a mok resource
    mr := new(mokResource)
    mr.SetID("test resource")
    rm.AddResource(mr)
    // at this point it was not started, so it should not be available
    if rm.IsAvailable(mr.GetID()) {
        t.Fatalf("Error. Resource [%s] was not started and results as available", mr.GetID())
    }
    rm.Start(mr.GetID(), 1)
    if !rm.IsAvailable(mr.GetID()) {
        t.Fatalf("Error. Resource [%s] was  started and does not results as available", mr.GetID())
    }
    // test progress setter/getter
    progress := float32(50)
    mr.SetProgress(progress)
    if mr.GetProgress() != progress {
        t.Fatalf("Error. Resource [%s] progress is [%v] instead of expected [%v]", mr.GetID(), mr.GetProgress(), progress)
    }
    // test list
    fmt.Println("test ====")
    rm.ListResources()
    // test print stats
    rm.ListStats()
}