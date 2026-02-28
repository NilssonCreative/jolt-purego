package jolt

// This file declares the raw C function variables that purego binds to.
// They are unexported because callers should use the Go wrapper types instead.
//
// Each variable corresponds to a symbol exported by the joltc shared library.
// The types use uintptr for opaque C pointers and native Go types for scalars/structs.

// --- Core ---
var jphInit func() bool
var jphShutdown func()

// --- JobSystem ---

// jobSystemThreadPoolConfig mirrors the C struct JobSystemThreadPoolConfig.
type jobSystemThreadPoolConfig struct {
	MaxJobs    uint32
	MaxBarrier uint32
	NumThreads int32
}

var jphJobSystemThreadPoolCreate func(config *jobSystemThreadPoolConfig) uintptr
var jphJobSystemDestroy func(jobSystem uintptr)

// --- BroadPhaseLayerInterface ---
var jphBroadPhaseLayerInterfaceTableCreate func(numObjectLayers, numBroadPhaseLayers uint32) uintptr
var jphBroadPhaseLayerInterfaceTableMapObjectToBroadPhaseLayer func(bpInterface uintptr, objectLayer uint32, broadPhaseLayer uint8)

// --- ObjectLayerPairFilter ---
var jphObjectLayerPairFilterTableCreate func(numObjectLayers uint32) uintptr
var jphObjectLayerPairFilterTableEnableCollision func(filter uintptr, layer1, layer2 uint32)
var jphObjectLayerPairFilterTableDisableCollision func(filter uintptr, layer1, layer2 uint32)

// --- ObjectVsBroadPhaseLayerFilter ---
var jphObjectVsBroadPhaseLayerFilterTableCreate func(bpInterface uintptr, numBroadPhaseLayers uint32, objectFilter uintptr, numObjectLayers uint32) uintptr

// --- PhysicsSystem ---

// physicsSystemSettings mirrors the C struct JPH_PhysicsSystemSettings.
type physicsSystemSettings struct {
	MaxBodies                     uint32
	NumBodyMutexes                uint32
	MaxBodyPairs                  uint32
	MaxContactConstraints         uint32
	_padding                      uint32
	BroadPhaseLayerInterface      uintptr
	ObjectLayerPairFilter         uintptr
	ObjectVsBroadPhaseLayerFilter uintptr
}

var jphPhysicsSystemCreate func(settings *physicsSystemSettings) uintptr
var jphPhysicsSystemDestroy func(system uintptr)
var jphPhysicsSystemOptimizeBroadPhase func(system uintptr)
var jphPhysicsSystemUpdate func(system uintptr, deltaTime float32, collisionSteps int32, jobSystem uintptr) int32
var jphPhysicsSystemGetBodyInterface func(system uintptr) uintptr
var jphPhysicsSystemSetGravity func(system uintptr, gravity *Vec3)
var jphPhysicsSystemGetGravity func(system uintptr, result *Vec3)
var jphPhysicsSystemGetNumBodies func(system uintptr) uint32
var jphPhysicsSystemGetNumActiveBodies func(system uintptr, bodyType int32) uint32
var jphPhysicsSystemGetMaxBodies func(system uintptr) uint32

// --- Shapes ---
var jphBoxShapeCreate func(halfExtent *Vec3, convexRadius float32) uintptr
var jphBoxShapeGetHalfExtent func(shape uintptr, result *Vec3)
var jphSphereShapeCreate func(radius float32) uintptr
var jphSphereShapeGetRadius func(shape uintptr) float32
var jphCapsuleShapeCreate func(halfHeight float32, radius float32) uintptr
var jphShapeDestroy func(shape uintptr)

// --- BodyCreationSettings ---
var jphBodyCreationSettingsCreate3 func(shape uintptr, position *Vec3, rotation *Quat, motionType int32, objectLayer uint32) uintptr
var jphBodyCreationSettingsDestroy func(settings uintptr)
var jphBodyCreationSettingsSetLinearVelocity func(settings uintptr, velocity *Vec3)
var jphBodyCreationSettingsGetLinearVelocity func(settings uintptr, velocity *Vec3)
var jphBodyCreationSettingsSetFriction func(settings uintptr, value float32)
var jphBodyCreationSettingsGetFriction func(settings uintptr) float32
var jphBodyCreationSettingsSetRestitution func(settings uintptr, value float32)
var jphBodyCreationSettingsGetRestitution func(settings uintptr) float32
var jphBodyCreationSettingsSetGravityFactor func(settings uintptr, value float32)
var jphBodyCreationSettingsGetGravityFactor func(settings uintptr) float32
var jphBodyCreationSettingsSetAllowSleeping func(settings uintptr, value bool)
var jphBodyCreationSettingsGetAllowSleeping func(settings uintptr) bool
var jphBodyCreationSettingsSetMotionQuality func(settings uintptr, value int32)
var jphBodyCreationSettingsGetMotionQuality func(settings uintptr) int32

// --- BodyInterface ---
var jphBodyInterfaceCreateAndAddBody func(bi uintptr, settings uintptr, activation int32) uint32
var jphBodyInterfaceRemoveAndDestroyBody func(bi uintptr, bodyID uint32)
var jphBodyInterfaceRemoveBody func(bi uintptr, bodyID uint32)
var jphBodyInterfaceDestroyBody func(bi uintptr, bodyID uint32)
var jphBodyInterfaceIsAdded func(bi uintptr, bodyID uint32) bool
var jphBodyInterfaceSetLinearVelocity func(bi uintptr, bodyID uint32, velocity *Vec3)
var jphBodyInterfaceGetLinearVelocity func(bi uintptr, bodyID uint32, velocity *Vec3)
var jphBodyInterfaceGetCenterOfMassPosition func(bi uintptr, bodyID uint32, position *Vec3)
var jphBodyInterfaceSetPosition func(bi uintptr, bodyID uint32, position *Vec3, activation int32)
var jphBodyInterfaceGetPosition func(bi uintptr, bodyID uint32, result *Vec3)
var jphBodyInterfaceSetRotation func(bi uintptr, bodyID uint32, rotation *Quat, activation int32)
var jphBodyInterfaceGetRotation func(bi uintptr, bodyID uint32, result *Quat)
var jphBodyInterfaceActivateBody func(bi uintptr, bodyID uint32)
var jphBodyInterfaceDeactivateBody func(bi uintptr, bodyID uint32)
var jphBodyInterfaceIsActive func(bi uintptr, bodyID uint32) bool
var jphBodyInterfaceAddForce func(bi uintptr, bodyID uint32, force *Vec3)
var jphBodyInterfaceAddImpulse func(bi uintptr, bodyID uint32, impulse *Vec3)
var jphBodyInterfaceSetFriction func(bi uintptr, bodyID uint32, friction float32)
var jphBodyInterfaceGetFriction func(bi uintptr, bodyID uint32) float32
var jphBodyInterfaceSetRestitution func(bi uintptr, bodyID uint32, restitution float32)
var jphBodyInterfaceGetRestitution func(bi uintptr, bodyID uint32) float32
var jphBodyInterfaceSetGravityFactor func(bi uintptr, bodyID uint32, value float32)
var jphBodyInterfaceGetGravityFactor func(bi uintptr, bodyID uint32) float32
var jphBodyInterfaceGetMotionType func(bi uintptr, bodyID uint32) int32
var jphBodyInterfaceSetMotionType func(bi uintptr, bodyID uint32, motionType int32, activation int32)
