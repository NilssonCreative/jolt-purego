// Package jolt provides a Go-idiomatic wrapper for the JoltPhysics engine.
//
// It communicates with the joltc C library using purego (no cgo required).
// All C pointers are wrapped in Go structs; callers use native Go types.
//
// # Quick Start
//
//	import "github.com/NilssonCreative/jolt-purego/jolt"
//
//	func main() {
//	    if err := jolt.Init(); err != nil { panic(err) }
//	    defer jolt.Shutdown()
//
//	    jobSystem := jolt.NewJobSystem(0, 0, -1)
//	    defer jobSystem.Close()
//
//	    // Set up collision layers ...
//	    // Create PhysicsSystem, shapes, bodies ...
//	    // Step the simulation ...
//	}
//
// # Resource Management
//
// Resources that own C memory provide a Close() or Destroy() method.
// Always defer these calls to avoid leaks:
//
//   - [JobSystem].Close
//   - [PhysicsSystem].Close
//   - [BodyCreationSettings].Close
//   - [Shape].Destroy (only for shapes not referenced by any body)
//
// # Thread Safety
//
// The underlying joltc library is designed for multi-threaded use via its
// own job system. However, Go wrapper types are not safe for concurrent use
// from multiple goroutines without external synchronization.
//
// # Limitations
//
//   - Requires the joltc shared library (libjoltc.so / libjoltc.dylib / joltc.dll)
//     to be available at runtime (in LD_LIBRARY_PATH, DYLD_LIBRARY_PATH, or PATH).
//   - Supported platforms: linux/amd64, linux/arm64, darwin/amd64, darwin/arm64,
//     windows/amd64 (matching purego's tier-1 support).
//   - Double-precision builds of joltc are not supported by this wrapper.
package jolt
