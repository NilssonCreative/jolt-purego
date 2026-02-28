// Package jolt provides a Go-idiomatic wrapper for the JoltPhysics engine
// via the joltc C API, using purego for Go↔C calls without cgo.
//
// All C pointers and handles are hidden behind Go structs. Users interact with
// native Go types (float32, structs, slices) rather than unsafe.Pointer or uintptr.
package jolt

// Vec3 represents a 3D vector (single-precision).
type Vec3 struct {
	X, Y, Z float32
}

// Vec4 represents a 4D vector (single-precision).
type Vec4 struct {
	X, Y, Z, W float32
}

// Quat represents a quaternion rotation.
// Component order matches joltc: (X, Y, Z, W) where W is the scalar part.
type Quat struct {
	X, Y, Z, W float32
}

// QuatIdentity returns the identity quaternion (no rotation).
func QuatIdentity() Quat {
	return Quat{X: 0, Y: 0, Z: 0, W: 1}
}

// BodyID is an opaque identifier for a physics body.
type BodyID uint32

// ObjectLayer identifies which collision layer an object belongs to.
type ObjectLayer uint32

// BroadPhaseLayer identifies a broad-phase collision layer.
type BroadPhaseLayer uint8

// MotionType specifies how a body moves in the simulation.
type MotionType int32

const (
	MotionTypeStatic    MotionType = 0
	MotionTypeKinematic MotionType = 1
	MotionTypeDynamic   MotionType = 2
)

// Activation specifies whether a body should be activated when added or modified.
type Activation int32

const (
	Activate     Activation = 0
	DontActivate Activation = 1
)

// PhysicsUpdateError represents errors from a physics update step.
type PhysicsUpdateError int32

const (
	PhysicsUpdateErrorNone                  PhysicsUpdateError = 0
	PhysicsUpdateErrorManifoldCacheFull     PhysicsUpdateError = 1 << 0
	PhysicsUpdateErrorBodyPairCacheFull     PhysicsUpdateError = 1 << 1
	PhysicsUpdateErrorContactConstraintFull PhysicsUpdateError = 1 << 2
)

// MotionQuality controls the quality of motion detection.
type MotionQuality int32

const (
	MotionQualityDiscrete   MotionQuality = 0
	MotionQualityLinearCast MotionQuality = 1
)

// AllowedDOFs specifies which degrees of freedom a body is allowed to move in.
type AllowedDOFs int32

const (
	AllowedDOFsAll          AllowedDOFs = 0b111111
	AllowedDOFsTranslationX AllowedDOFs = 0b000001
	AllowedDOFsTranslationY AllowedDOFs = 0b000010
	AllowedDOFsTranslationZ AllowedDOFs = 0b000100
	AllowedDOFsRotationX    AllowedDOFs = 0b001000
	AllowedDOFsRotationY    AllowedDOFs = 0b010000
	AllowedDOFsRotationZ    AllowedDOFs = 0b100000
	AllowedDOFsPlane2D      AllowedDOFs = AllowedDOFsTranslationX | AllowedDOFsTranslationY | AllowedDOFsRotationZ
)
