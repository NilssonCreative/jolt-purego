package jolt

import "fmt"

// initialized tracks whether Init() has been called successfully.
var initialized bool

// Init initializes the Jolt physics engine.
// This must be called before any other Jolt function.
// It loads the joltc shared library and calls JPH_Init.
func Init() error {
	if initialized {
		return nil
	}
	if err := loadLibrary(); err != nil {
		return err
	}
	if !jphInit() {
		return fmt.Errorf("jolt: JPH_Init failed")
	}
	initialized = true
	return nil
}

// Shutdown shuts down the Jolt physics engine and releases global resources.
// After calling Shutdown, no other Jolt functions should be called.
func Shutdown() {
	if !initialized {
		return
	}
	jphShutdown()
	initialized = false
}

// JobSystem manages a thread pool used by the physics engine for parallel work.
// The underlying C resource must be released by calling Close().
type JobSystem struct {
	handle uintptr
}

// NewJobSystem creates a new thread pool–based job system.
// numThreads specifies the number of worker threads; use 0 or -1 to let
// joltc auto-detect based on available CPU cores.
// maxJobs and maxBarriers control internal capacity (use 0 for defaults).
func NewJobSystem(maxJobs, maxBarriers uint32, numThreads int32) *JobSystem {
	if maxJobs == 0 {
		maxJobs = 2048
	}
	if maxBarriers == 0 {
		maxBarriers = 8
	}
	cfg := &jobSystemThreadPoolConfig{
		MaxJobs:    maxJobs,
		MaxBarrier: maxBarriers,
		NumThreads: numThreads,
	}
	h := jphJobSystemThreadPoolCreate(cfg)
	return &JobSystem{handle: h}
}

// Close releases the underlying C job system.
func (js *JobSystem) Close() {
	if js.handle != 0 {
		jphJobSystemDestroy(js.handle)
		js.handle = 0
	}
}

// BroadPhaseLayerInterface maps object layers to broad-phase layers.
// The underlying C resource is owned by the PhysicsSystem that uses it;
// it does not need to be independently freed, but must remain valid for
// the lifetime of the PhysicsSystem.
type BroadPhaseLayerInterface struct {
	handle uintptr
}

// NewBroadPhaseLayerInterfaceTable creates a table-based mapping from
// object layers to broad-phase layers.
func NewBroadPhaseLayerInterfaceTable(numObjectLayers, numBroadPhaseLayers uint32) *BroadPhaseLayerInterface {
	h := jphBroadPhaseLayerInterfaceTableCreate(numObjectLayers, numBroadPhaseLayers)
	return &BroadPhaseLayerInterface{handle: h}
}

// MapObjectToBroadPhaseLayer maps an object layer to a broad-phase layer.
func (b *BroadPhaseLayerInterface) MapObjectToBroadPhaseLayer(objectLayer ObjectLayer, broadPhaseLayer BroadPhaseLayer) {
	jphBroadPhaseLayerInterfaceTableMapObjectToBroadPhaseLayer(b.handle, uint32(objectLayer), uint8(broadPhaseLayer))
}

// ObjectLayerPairFilter determines which object layer pairs should collide.
type ObjectLayerPairFilter struct {
	handle uintptr
}

// NewObjectLayerPairFilterTable creates a table-based object layer pair filter.
func NewObjectLayerPairFilterTable(numObjectLayers uint32) *ObjectLayerPairFilter {
	h := jphObjectLayerPairFilterTableCreate(numObjectLayers)
	return &ObjectLayerPairFilter{handle: h}
}

// EnableCollision enables collision between two object layers.
func (f *ObjectLayerPairFilter) EnableCollision(layer1, layer2 ObjectLayer) {
	jphObjectLayerPairFilterTableEnableCollision(f.handle, uint32(layer1), uint32(layer2))
}

// DisableCollision disables collision between two object layers.
func (f *ObjectLayerPairFilter) DisableCollision(layer1, layer2 ObjectLayer) {
	jphObjectLayerPairFilterTableDisableCollision(f.handle, uint32(layer1), uint32(layer2))
}

// ObjectVsBroadPhaseLayerFilter determines which object layers collide with which broad-phase layers.
type ObjectVsBroadPhaseLayerFilter struct {
	handle uintptr
}

// NewObjectVsBroadPhaseLayerFilterTable creates a table-based filter using the
// given broad-phase layer interface and object layer pair filter.
func NewObjectVsBroadPhaseLayerFilterTable(
	bpInterface *BroadPhaseLayerInterface,
	numBroadPhaseLayers uint32,
	objectFilter *ObjectLayerPairFilter,
	numObjectLayers uint32,
) *ObjectVsBroadPhaseLayerFilter {
	h := jphObjectVsBroadPhaseLayerFilterTableCreate(
		bpInterface.handle, numBroadPhaseLayers,
		objectFilter.handle, numObjectLayers,
	)
	return &ObjectVsBroadPhaseLayerFilter{handle: h}
}
