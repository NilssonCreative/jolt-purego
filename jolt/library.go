package jolt

import (
	"fmt"
	"runtime"

	"github.com/ebitengine/purego"
	"golang.org/x/sys/windows"
)

// lib holds the handle to the loaded joltc shared library.
// It is set once during Init() and never changed afterward.
var lib uintptr

// getLibraryName returns the platform-specific shared library file name for joltc.
func getLibraryName() string {
	switch runtime.GOOS {
	case "windows":
		return "joltc.dll"
	case "darwin":
		return "libjoltc.dylib"
	case "linux":
		return "libjoltc.so"
	default:
		panic(fmt.Sprintf("jolt: unsupported GOOS=%s", runtime.GOOS))
	}
}

// loadLibrary opens the joltc shared library and registers all C function symbols
// used by this wrapper. It is called once from Init().
func loadLibrary() error {
	joltDll := windows.NewLazyDLL(getLibraryName())
	if joltDll == nil {
		return fmt.Errorf("jolt: failed to load joltc library (%s)", getLibraryName())
	}
	lib = joltDll.Handle()
	registerSymbols(joltDll.Handle())
	return nil
}

// registerSymbols uses purego.RegisterLibFunc to bind Go function variables
// to their corresponding C symbols in the joltc shared library.
func registerSymbols(handle uintptr) {
	// --- Core ---
	purego.RegisterLibFunc(&jphInit, handle, "JPH_Init")
	purego.RegisterLibFunc(&jphShutdown, handle, "JPH_Shutdown")

	// --- JobSystem ---
	purego.RegisterLibFunc(&jphJobSystemThreadPoolCreate, handle, "JPH_JobSystemThreadPool_Create")
	purego.RegisterLibFunc(&jphJobSystemDestroy, handle, "JPH_JobSystem_Destroy")

	// --- BroadPhaseLayerInterface ---
	purego.RegisterLibFunc(&jphBroadPhaseLayerInterfaceTableCreate, handle, "JPH_BroadPhaseLayerInterfaceTable_Create")
	purego.RegisterLibFunc(&jphBroadPhaseLayerInterfaceTableMapObjectToBroadPhaseLayer, handle, "JPH_BroadPhaseLayerInterfaceTable_MapObjectToBroadPhaseLayer")

	// --- ObjectLayerPairFilter ---
	purego.RegisterLibFunc(&jphObjectLayerPairFilterTableCreate, handle, "JPH_ObjectLayerPairFilterTable_Create")
	purego.RegisterLibFunc(&jphObjectLayerPairFilterTableEnableCollision, handle, "JPH_ObjectLayerPairFilterTable_EnableCollision")
	purego.RegisterLibFunc(&jphObjectLayerPairFilterTableDisableCollision, handle, "JPH_ObjectLayerPairFilterTable_DisableCollision")

	// --- ObjectVsBroadPhaseLayerFilter ---
	purego.RegisterLibFunc(&jphObjectVsBroadPhaseLayerFilterTableCreate, handle, "JPH_ObjectVsBroadPhaseLayerFilterTable_Create")

	// --- PhysicsSystem ---
	purego.RegisterLibFunc(&jphPhysicsSystemCreate, handle, "JPH_PhysicsSystem_Create")
	purego.RegisterLibFunc(&jphPhysicsSystemDestroy, handle, "JPH_PhysicsSystem_Destroy")
	purego.RegisterLibFunc(&jphPhysicsSystemOptimizeBroadPhase, handle, "JPH_PhysicsSystem_OptimizeBroadPhase")
	purego.RegisterLibFunc(&jphPhysicsSystemUpdate, handle, "JPH_PhysicsSystem_Update")
	purego.RegisterLibFunc(&jphPhysicsSystemGetBodyInterface, handle, "JPH_PhysicsSystem_GetBodyInterface")
	purego.RegisterLibFunc(&jphPhysicsSystemSetGravity, handle, "JPH_PhysicsSystem_SetGravity")
	purego.RegisterLibFunc(&jphPhysicsSystemGetGravity, handle, "JPH_PhysicsSystem_GetGravity")
	purego.RegisterLibFunc(&jphPhysicsSystemGetNumBodies, handle, "JPH_PhysicsSystem_GetNumBodies")
	purego.RegisterLibFunc(&jphPhysicsSystemGetNumActiveBodies, handle, "JPH_PhysicsSystem_GetNumActiveBodies")
	purego.RegisterLibFunc(&jphPhysicsSystemGetMaxBodies, handle, "JPH_PhysicsSystem_GetMaxBodies")

	// --- Shapes ---
	purego.RegisterLibFunc(&jphBoxShapeCreate, handle, "JPH_BoxShape_Create")
	purego.RegisterLibFunc(&jphBoxShapeGetHalfExtent, handle, "JPH_BoxShape_GetHalfExtent")
	purego.RegisterLibFunc(&jphSphereShapeCreate, handle, "JPH_SphereShape_Create")
	purego.RegisterLibFunc(&jphSphereShapeGetRadius, handle, "JPH_SphereShape_GetRadius")
	purego.RegisterLibFunc(&jphCapsuleShapeCreate, handle, "JPH_CapsuleShape_Create")
	purego.RegisterLibFunc(&jphShapeDestroy, handle, "JPH_Shape_Destroy")

	// --- BodyCreationSettings ---
	purego.RegisterLibFunc(&jphBodyCreationSettingsCreate3, handle, "JPH_BodyCreationSettings_Create3")
	purego.RegisterLibFunc(&jphBodyCreationSettingsDestroy, handle, "JPH_BodyCreationSettings_Destroy")
	purego.RegisterLibFunc(&jphBodyCreationSettingsSetLinearVelocity, handle, "JPH_BodyCreationSettings_SetLinearVelocity")
	purego.RegisterLibFunc(&jphBodyCreationSettingsGetLinearVelocity, handle, "JPH_BodyCreationSettings_GetLinearVelocity")
	purego.RegisterLibFunc(&jphBodyCreationSettingsSetFriction, handle, "JPH_BodyCreationSettings_SetFriction")
	purego.RegisterLibFunc(&jphBodyCreationSettingsGetFriction, handle, "JPH_BodyCreationSettings_GetFriction")
	purego.RegisterLibFunc(&jphBodyCreationSettingsSetRestitution, handle, "JPH_BodyCreationSettings_SetRestitution")
	purego.RegisterLibFunc(&jphBodyCreationSettingsGetRestitution, handle, "JPH_BodyCreationSettings_GetRestitution")
	purego.RegisterLibFunc(&jphBodyCreationSettingsSetGravityFactor, handle, "JPH_BodyCreationSettings_SetGravityFactor")
	purego.RegisterLibFunc(&jphBodyCreationSettingsGetGravityFactor, handle, "JPH_BodyCreationSettings_GetGravityFactor")
	purego.RegisterLibFunc(&jphBodyCreationSettingsSetAllowSleeping, handle, "JPH_BodyCreationSettings_SetAllowSleeping")
	purego.RegisterLibFunc(&jphBodyCreationSettingsGetAllowSleeping, handle, "JPH_BodyCreationSettings_GetAllowSleeping")
	purego.RegisterLibFunc(&jphBodyCreationSettingsSetMotionQuality, handle, "JPH_BodyCreationSettings_SetMotionQuality")
	purego.RegisterLibFunc(&jphBodyCreationSettingsGetMotionQuality, handle, "JPH_BodyCreationSettings_GetMotionQuality")

	// --- BodyInterface ---
	purego.RegisterLibFunc(&jphBodyInterfaceCreateAndAddBody, handle, "JPH_BodyInterface_CreateAndAddBody")
	purego.RegisterLibFunc(&jphBodyInterfaceRemoveAndDestroyBody, handle, "JPH_BodyInterface_RemoveAndDestroyBody")
	purego.RegisterLibFunc(&jphBodyInterfaceRemoveBody, handle, "JPH_BodyInterface_RemoveBody")
	purego.RegisterLibFunc(&jphBodyInterfaceDestroyBody, handle, "JPH_BodyInterface_DestroyBody")
	purego.RegisterLibFunc(&jphBodyInterfaceIsAdded, handle, "JPH_BodyInterface_IsAdded")
	purego.RegisterLibFunc(&jphBodyInterfaceSetLinearVelocity, handle, "JPH_BodyInterface_SetLinearVelocity")
	purego.RegisterLibFunc(&jphBodyInterfaceGetLinearVelocity, handle, "JPH_BodyInterface_GetLinearVelocity")
	purego.RegisterLibFunc(&jphBodyInterfaceGetCenterOfMassPosition, handle, "JPH_BodyInterface_GetCenterOfMassPosition")
	purego.RegisterLibFunc(&jphBodyInterfaceSetPosition, handle, "JPH_BodyInterface_SetPosition")
	purego.RegisterLibFunc(&jphBodyInterfaceGetPosition, handle, "JPH_BodyInterface_GetPosition")
	purego.RegisterLibFunc(&jphBodyInterfaceSetRotation, handle, "JPH_BodyInterface_SetRotation")
	purego.RegisterLibFunc(&jphBodyInterfaceGetRotation, handle, "JPH_BodyInterface_GetRotation")
	purego.RegisterLibFunc(&jphBodyInterfaceActivateBody, handle, "JPH_BodyInterface_ActivateBody")
	purego.RegisterLibFunc(&jphBodyInterfaceDeactivateBody, handle, "JPH_BodyInterface_DeactivateBody")
	purego.RegisterLibFunc(&jphBodyInterfaceIsActive, handle, "JPH_BodyInterface_IsActive")
	purego.RegisterLibFunc(&jphBodyInterfaceAddForce, handle, "JPH_BodyInterface_AddForce")
	purego.RegisterLibFunc(&jphBodyInterfaceAddImpulse, handle, "JPH_BodyInterface_AddImpulse")
	purego.RegisterLibFunc(&jphBodyInterfaceSetFriction, handle, "JPH_BodyInterface_SetFriction")
	purego.RegisterLibFunc(&jphBodyInterfaceGetFriction, handle, "JPH_BodyInterface_GetFriction")
	purego.RegisterLibFunc(&jphBodyInterfaceSetRestitution, handle, "JPH_BodyInterface_SetRestitution")
	purego.RegisterLibFunc(&jphBodyInterfaceGetRestitution, handle, "JPH_BodyInterface_GetRestitution")
	purego.RegisterLibFunc(&jphBodyInterfaceSetGravityFactor, handle, "JPH_BodyInterface_SetGravityFactor")
	purego.RegisterLibFunc(&jphBodyInterfaceGetGravityFactor, handle, "JPH_BodyInterface_GetGravityFactor")
	purego.RegisterLibFunc(&jphBodyInterfaceGetMotionType, handle, "JPH_BodyInterface_GetMotionType")
	purego.RegisterLibFunc(&jphBodyInterfaceSetMotionType, handle, "JPH_BodyInterface_SetMotionType")
}
