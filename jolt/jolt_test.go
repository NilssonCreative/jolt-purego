package jolt

import "testing"

func TestVec3(t *testing.T) {
	v := Vec3{X: 1.0, Y: 2.0, Z: 3.0}
	if v.X != 1.0 || v.Y != 2.0 || v.Z != 3.0 {
		t.Errorf("Vec3 fields incorrect: got %+v", v)
	}
}

func TestVec4(t *testing.T) {
	v := Vec4{X: 1.0, Y: 2.0, Z: 3.0, W: 4.0}
	if v.X != 1.0 || v.Y != 2.0 || v.Z != 3.0 || v.W != 4.0 {
		t.Errorf("Vec4 fields incorrect: got %+v", v)
	}
}

func TestQuatIdentity(t *testing.T) {
	q := QuatIdentity()
	if q.X != 0 || q.Y != 0 || q.Z != 0 || q.W != 1 {
		t.Errorf("QuatIdentity incorrect: got %+v", q)
	}
}

func TestQuat(t *testing.T) {
	q := Quat{X: 0.1, Y: 0.2, Z: 0.3, W: 0.9}
	if q.X != 0.1 || q.Y != 0.2 || q.Z != 0.3 || q.W != 0.9 {
		t.Errorf("Quat fields incorrect: got %+v", q)
	}
}

func TestMotionTypeConstants(t *testing.T) {
	if MotionTypeStatic != 0 {
		t.Errorf("MotionTypeStatic should be 0, got %d", MotionTypeStatic)
	}
	if MotionTypeKinematic != 1 {
		t.Errorf("MotionTypeKinematic should be 1, got %d", MotionTypeKinematic)
	}
	if MotionTypeDynamic != 2 {
		t.Errorf("MotionTypeDynamic should be 2, got %d", MotionTypeDynamic)
	}
}

func TestActivationConstants(t *testing.T) {
	if Activate != 0 {
		t.Errorf("Activate should be 0, got %d", Activate)
	}
	if DontActivate != 1 {
		t.Errorf("DontActivate should be 1, got %d", DontActivate)
	}
}

func TestPhysicsUpdateErrorConstants(t *testing.T) {
	if PhysicsUpdateErrorNone != 0 {
		t.Errorf("PhysicsUpdateErrorNone should be 0, got %d", PhysicsUpdateErrorNone)
	}
	if PhysicsUpdateErrorManifoldCacheFull != 1 {
		t.Errorf("PhysicsUpdateErrorManifoldCacheFull should be 1, got %d", PhysicsUpdateErrorManifoldCacheFull)
	}
	if PhysicsUpdateErrorBodyPairCacheFull != 2 {
		t.Errorf("PhysicsUpdateErrorBodyPairCacheFull should be 2, got %d", PhysicsUpdateErrorBodyPairCacheFull)
	}
	if PhysicsUpdateErrorContactConstraintFull != 4 {
		t.Errorf("PhysicsUpdateErrorContactConstraintFull should be 4, got %d", PhysicsUpdateErrorContactConstraintFull)
	}
}

func TestMotionQualityConstants(t *testing.T) {
	if MotionQualityDiscrete != 0 {
		t.Errorf("MotionQualityDiscrete should be 0, got %d", MotionQualityDiscrete)
	}
	if MotionQualityLinearCast != 1 {
		t.Errorf("MotionQualityLinearCast should be 1, got %d", MotionQualityLinearCast)
	}
}

func TestAllowedDOFsConstants(t *testing.T) {
	if AllowedDOFsAll != 0b111111 {
		t.Errorf("AllowedDOFsAll should be 0b111111, got %b", AllowedDOFsAll)
	}
	if AllowedDOFsPlane2D != (AllowedDOFsTranslationX | AllowedDOFsTranslationY | AllowedDOFsRotationZ) {
		t.Errorf("AllowedDOFsPlane2D should be TranslationX|TranslationY|RotationZ, got %b", AllowedDOFsPlane2D)
	}
}

func TestBodyIDType(t *testing.T) {
	var id BodyID = 42
	if uint32(id) != 42 {
		t.Errorf("BodyID conversion failed")
	}
}

func TestObjectLayerType(t *testing.T) {
	var layer ObjectLayer = 1
	if uint32(layer) != 1 {
		t.Errorf("ObjectLayer conversion failed")
	}
}

func TestBroadPhaseLayerType(t *testing.T) {
	var layer BroadPhaseLayer = 0
	if uint8(layer) != 0 {
		t.Errorf("BroadPhaseLayer conversion failed")
	}
}

func TestGetLibraryName(t *testing.T) {
	name := getLibraryName()
	if name == "" {
		t.Error("getLibraryName returned empty string")
	}
	// Should be one of the known library names
	switch name {
	case "joltc.dll", "libjoltc.so", "libjoltc.dylib":
		// OK
	default:
		t.Errorf("unexpected library name: %s", name)
	}
}
