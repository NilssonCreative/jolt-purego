// This example demonstrates a minimal JoltPhysics simulation using the
// jolt-purego Go wrapper: creating a physics world, adding a floor and a
// falling sphere, and stepping the simulation.
//
// Prerequisites:
//   - The joltc shared library must be available at runtime.
//     See the project README for build/install instructions.
//
// Run:
//
//	go run .
package main

import (
	"fmt"
	"log"

	"github.com/NilssonCreative/jolt-purego/jolt"
)

// Collision layers
const (
	LayerNonMoving jolt.ObjectLayer = 0
	LayerMoving    jolt.ObjectLayer = 1
	NumLayers      uint32           = 2
)

// Broad-phase layers
const (
	BPLayerNonMoving jolt.BroadPhaseLayer = 0
	BPLayerMoving    jolt.BroadPhaseLayer = 1
	NumBPLayers      uint32               = 2
)

func main() {
	// 1. Initialize Jolt
	if err := jolt.Init(); err != nil {
		log.Fatalf("Failed to initialize Jolt: %v", err)
	}
	defer jolt.Shutdown()

	// 2. Create a job system (thread pool)
	jobSystem := jolt.NewJobSystem(0, 0, -1)
	defer jobSystem.Close()

	// 3. Set up collision filtering
	//    Map object layers → broad-phase layers
	bpLayerInterface := jolt.NewBroadPhaseLayerInterfaceTable(NumLayers, NumBPLayers)
	bpLayerInterface.MapObjectToBroadPhaseLayer(LayerNonMoving, BPLayerNonMoving)
	bpLayerInterface.MapObjectToBroadPhaseLayer(LayerMoving, BPLayerMoving)

	//    Define which object layers can collide
	objectFilter := jolt.NewObjectLayerPairFilterTable(NumLayers)
	objectFilter.EnableCollision(LayerNonMoving, LayerMoving)
	objectFilter.EnableCollision(LayerMoving, LayerMoving)

	//    Define which object layers collide with which broad-phase layers
	bpFilter := jolt.NewObjectVsBroadPhaseLayerFilterTable(
		bpLayerInterface, NumBPLayers,
		objectFilter, NumLayers,
	)

	// 4. Create the physics system
	physicsSystem := jolt.NewPhysicsSystem(&jolt.PhysicsSystemConfig{
		MaxBodies:             1024,
		MaxBodyPairs:          1024,
		MaxContactConstraints: 1024,
		BroadPhaseLayer:       bpLayerInterface,
		ObjectLayerPairFilter: objectFilter,
		ObjectVsBPLayerFilter: bpFilter,
	})
	defer physicsSystem.Close()

	bodyInterface := physicsSystem.GetBodyInterface()

	// 5. Create a static floor (box shape)
	floorShape := jolt.NewBoxShape(jolt.Vec3{X: 100, Y: 1, Z: 100}, 0.0)
	floorSettings := jolt.NewBodyCreationSettings(
		floorShape,
		jolt.Vec3{X: 0, Y: -1, Z: 0},
		jolt.QuatIdentity(),
		jolt.MotionTypeStatic,
		LayerNonMoving,
	)
	floorID := bodyInterface.CreateAndAddBody(floorSettings, jolt.DontActivate)
	floorSettings.Close()
	fmt.Printf("Floor body ID: %d\n", floorID)

	// 6. Create a dynamic sphere that will fall onto the floor
	sphereShape := jolt.NewSphereShape(0.5)
	sphereSettings := jolt.NewBodyCreationSettings(
		sphereShape,
		jolt.Vec3{X: 0, Y: 10, Z: 0},
		jolt.QuatIdentity(),
		jolt.MotionTypeDynamic,
		LayerMoving,
	)
	sphereSettings.SetRestitution(0.5) // Some bounciness
	sphereID := bodyInterface.CreateAndAddBody(sphereSettings, jolt.Activate)
	sphereSettings.Close()
	fmt.Printf("Sphere body ID: %d\n", sphereID)

	// Optimize after adding all bodies
	physicsSystem.OptimizeBroadPhase()

	// 7. Step the simulation
	deltaTime := float32(1.0 / 60.0)
	numSteps := 120

	fmt.Printf("\nSimulating %d steps (%.1f seconds)...\n\n", numSteps, float32(numSteps)*deltaTime)

	for step := 0; step < numSteps; step++ {
		physicsSystem.Update(deltaTime, 1, jobSystem)

		if step%10 == 0 {
			pos := bodyInterface.GetCenterOfMassPosition(sphereID)
			vel := bodyInterface.GetLinearVelocity(sphereID)
			fmt.Printf("Step %3d: pos=(%.2f, %.2f, %.2f)  vel=(%.2f, %.2f, %.2f)\n",
				step, pos.X, pos.Y, pos.Z, vel.X, vel.Y, vel.Z)
		}
	}

	// Final position
	finalPos := bodyInterface.GetCenterOfMassPosition(sphereID)
	fmt.Printf("\nFinal position: (%.2f, %.2f, %.2f)\n", finalPos.X, finalPos.Y, finalPos.Z)

	// 8. Clean up bodies
	bodyInterface.RemoveAndDestroyBody(sphereID)
	bodyInterface.RemoveAndDestroyBody(floorID)
}
