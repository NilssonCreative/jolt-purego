package jolt

// PhysicsSystem is the main physics simulation manager.
// It owns the broad-phase, narrow-phase, body storage, and provides
// the BodyInterface used to add/remove/manipulate bodies.
//
// Close() must be called when the system is no longer needed to free
// the underlying C resources.
type PhysicsSystem struct {
	handle uintptr

	// Keep references to prevent garbage collection of resources that the
	// C PhysicsSystem holds pointers to.
	bpLayerInterface *BroadPhaseLayerInterface
	objLayerFilter   *ObjectLayerPairFilter
	objVsBPFilter    *ObjectVsBroadPhaseLayerFilter
}

// PhysicsSystemConfig holds configuration for creating a PhysicsSystem.
type PhysicsSystemConfig struct {
	MaxBodies             uint32
	NumBodyMutexes        uint32
	MaxBodyPairs          uint32
	MaxContactConstraints uint32
	BroadPhaseLayer       *BroadPhaseLayerInterface
	ObjectLayerPairFilter *ObjectLayerPairFilter
	ObjectVsBPLayerFilter *ObjectVsBroadPhaseLayerFilter
}

// NewPhysicsSystem creates a new physics system with the given configuration.
// The caller is responsible for calling Close() when done.
func NewPhysicsSystem(cfg *PhysicsSystemConfig) *PhysicsSystem {
	if cfg.MaxBodies == 0 {
		cfg.MaxBodies = 10240
	}
	if cfg.MaxBodyPairs == 0 {
		cfg.MaxBodyPairs = 65536
	}
	if cfg.MaxContactConstraints == 0 {
		cfg.MaxContactConstraints = 10240
	}

	settings := &physicsSystemSettings{
		MaxBodies:                     cfg.MaxBodies,
		NumBodyMutexes:                cfg.NumBodyMutexes,
		MaxBodyPairs:                  cfg.MaxBodyPairs,
		MaxContactConstraints:         cfg.MaxContactConstraints,
		BroadPhaseLayerInterface:      cfg.BroadPhaseLayer.handle,
		ObjectLayerPairFilter:         cfg.ObjectLayerPairFilter.handle,
		ObjectVsBroadPhaseLayerFilter: cfg.ObjectVsBPLayerFilter.handle,
	}
	h := jphPhysicsSystemCreate(settings)
	return &PhysicsSystem{
		handle:           h,
		bpLayerInterface: cfg.BroadPhaseLayer,
		objLayerFilter:   cfg.ObjectLayerPairFilter,
		objVsBPFilter:    cfg.ObjectVsBPLayerFilter,
	}
}

// Close destroys the physics system and releases all C resources.
func (ps *PhysicsSystem) Close() {
	if ps.handle != 0 {
		jphPhysicsSystemDestroy(ps.handle)
		ps.handle = 0
	}
}

// OptimizeBroadPhase optimizes the broad-phase data structure.
// Call this after adding a batch of bodies for better performance.
func (ps *PhysicsSystem) OptimizeBroadPhase() {
	jphPhysicsSystemOptimizeBroadPhase(ps.handle)
}

// Update steps the physics simulation forward by deltaTime seconds.
// collisionSteps is the number of collision sub-steps (typically 1).
// Returns a PhysicsUpdateError indicating any issues during the step.
func (ps *PhysicsSystem) Update(deltaTime float32, collisionSteps int, jobSystem *JobSystem) PhysicsUpdateError {
	result := jphPhysicsSystemUpdate(ps.handle, deltaTime, int32(collisionSteps), jobSystem.handle)
	return PhysicsUpdateError(result)
}

// GetBodyInterface returns the BodyInterface for adding, removing, and
// manipulating physics bodies. The returned BodyInterface is valid for
// the lifetime of this PhysicsSystem.
func (ps *PhysicsSystem) GetBodyInterface() *BodyInterface {
	h := jphPhysicsSystemGetBodyInterface(ps.handle)
	return &BodyInterface{handle: h}
}

// SetGravity sets the global gravity vector.
func (ps *PhysicsSystem) SetGravity(gravity Vec3) {
	jphPhysicsSystemSetGravity(ps.handle, &gravity)
}

// GetGravity returns the current global gravity vector.
func (ps *PhysicsSystem) GetGravity() Vec3 {
	var result Vec3
	jphPhysicsSystemGetGravity(ps.handle, &result)
	return result
}

// GetNumBodies returns the total number of bodies in the system.
func (ps *PhysicsSystem) GetNumBodies() uint32 {
	return jphPhysicsSystemGetNumBodies(ps.handle)
}

// GetNumActiveBodies returns the number of active rigid bodies.
func (ps *PhysicsSystem) GetNumActiveBodies() uint32 {
	return jphPhysicsSystemGetNumActiveBodies(ps.handle, 0) // 0 = JPH_BodyType_Rigid
}

// GetMaxBodies returns the maximum number of bodies supported.
func (ps *PhysicsSystem) GetMaxBodies() uint32 {
	return jphPhysicsSystemGetMaxBodies(ps.handle)
}
