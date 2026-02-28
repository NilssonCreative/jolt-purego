# jolt-purego

A Go-idiomatic wrapper for the [JoltPhysics](https://github.com/jrouwe/JoltPhysics) engine, using the [joltc](https://github.com/amerkoleci/joltc) C API and [purego](https://github.com/ebitengine/purego) for Go↔C calls — **no cgo required**.

## Features

- **No cgo** — uses purego for dynamic linking, enabling simple cross-compilation.
- **Go-idiomatic API** — all C pointers and handles are wrapped in Go structs with methods; callers use native Go types (`float32`, `Vec3`, `Quat`, etc.).
- **Resource management** — `Close()` and `Destroy()` methods for deterministic release of C resources.
- **Extensible** — follows a consistent pattern (symbols → raw functions → Go wrapper types) that makes adding new Jolt features straightforward.

## Supported Platforms

| OS      | Architecture | Notes                          |
|---------|-------------|--------------------------------|
| Linux   | amd64, arm64 | Requires `libjoltc.so`         |
| macOS   | amd64, arm64 | Requires `libjoltc.dylib`      |
| Windows | amd64       | Requires `joltc.dll`           |

Platform support matches [purego's tier-1 platforms](https://github.com/ebitengine/purego#supported-platforms).

## Prerequisites

### 1. Go 1.21+

Install from [go.dev](https://go.dev/dl/).

### 2. joltc shared library

You need the joltc native library built as a shared library (`.so`, `.dylib`, or `.dll`).

#### Building joltc from source

```bash
git clone https://github.com/amerkoleci/joltc.git
cd joltc
cmake -B build -DCMAKE_BUILD_TYPE=Release -DBUILD_SHARED_LIBS=ON -DJPH_SHARED_LIBRARY_BUILD=ON
cmake --build build --config Release
```

Then ensure the resulting library is in your library search path:

```bash
# Linux
export LD_LIBRARY_PATH=/path/to/joltc/build:$LD_LIBRARY_PATH

# macOS
export DYLD_LIBRARY_PATH=/path/to/joltc/build:$DYLD_LIBRARY_PATH

# Windows — add the directory containing joltc.dll to PATH
```

## Installation

```bash
go get github.com/NilssonCreative/jolt-purego
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    "github.com/NilssonCreative/jolt-purego/jolt"
)

const (
    LayerNonMoving jolt.ObjectLayer = 0
    LayerMoving    jolt.ObjectLayer = 1
)

func main() {
    // Initialize the engine
    if err := jolt.Init(); err != nil {
        log.Fatal(err)
    }
    defer jolt.Shutdown()

    // Create a job system for parallel physics
    jobSystem := jolt.NewJobSystem(0, 0, -1) // -1 = auto-detect threads
    defer jobSystem.Close()

    // Set up collision layers
    bpLayer := jolt.NewBroadPhaseLayerInterfaceTable(2, 2)
    bpLayer.MapObjectToBroadPhaseLayer(0, 0)
    bpLayer.MapObjectToBroadPhaseLayer(1, 1)

    objFilter := jolt.NewObjectLayerPairFilterTable(2)
    objFilter.EnableCollision(0, 1)
    objFilter.EnableCollision(1, 1)

    bpFilter := jolt.NewObjectVsBroadPhaseLayerFilterTable(bpLayer, 2, objFilter, 2)

    // Create the physics system
    physics := jolt.NewPhysicsSystem(&jolt.PhysicsSystemConfig{
        MaxBodies:             1024,
        MaxBodyPairs:          1024,
        MaxContactConstraints: 1024,
        BroadPhaseLayer:       bpLayer,
        ObjectLayerPairFilter: objFilter,
        ObjectVsBPLayerFilter: bpFilter,
    })
    defer physics.Close()

    bi := physics.GetBodyInterface()

    // Create a falling sphere
    sphere := jolt.NewSphereShape(0.5)
    settings := jolt.NewBodyCreationSettings(
        sphere,
        jolt.Vec3{X: 0, Y: 10, Z: 0},
        jolt.QuatIdentity(),
        jolt.MotionTypeDynamic,
        LayerMoving,
    )
    bodyID := bi.CreateAndAddBody(settings, jolt.Activate)
    settings.Close()

    physics.OptimizeBroadPhase()

    // Step the simulation
    for i := 0; i < 60; i++ {
        physics.Update(1.0/60.0, 1, jobSystem)
    }

    pos := bi.GetCenterOfMassPosition(bodyID)
    fmt.Printf("Sphere position after 1s: (%.2f, %.2f, %.2f)\n", pos.X, pos.Y, pos.Z)

    bi.RemoveAndDestroyBody(bodyID)
}
```

## Package Structure

```
jolt-purego/
├── jolt/                           # Main Go wrapper package
│   ├── doc.go                      # Package documentation
│   ├── types.go                    # Core types (Vec3, Quat, enums)
│   ├── library.go                  # Library loading and symbol registration
│   ├── symbols.go                  # Raw C function variable declarations
│   ├── jolt.go                     # Init/Shutdown, JobSystem, layer filters
│   ├── physics_system.go           # PhysicsSystem wrapper
│   ├── shapes.go                   # Shape wrappers (Box, Sphere, Capsule)
│   ├── body_creation_settings.go   # BodyCreationSettings wrapper
│   └── body_interface.go           # BodyInterface wrapper
├── examples/
│   └── basic/                      # Minimal simulation example
│       └── main.go
├── LICENSE                         # MIT License
├── README.md                       # This file
├── go.mod
└── go.sum
```

## Extending the Wrapper

To add support for more joltc features, follow this pattern:

1. **Add raw function variables** in `jolt/symbols.go` matching the C signature.
2. **Register the symbols** in `jolt/library.go` inside `registerSymbols()`.
3. **Create a Go wrapper type** (if needed) in a new or existing file, hiding the `uintptr` handle.
4. **Add methods** that call the raw functions and convert between Go types and C types.

Example — adding `JPH_CylinderShape_Create`:

```go
// In symbols.go:
var jphCylinderShapeCreate func(halfHeight float32, radius float32) uintptr

// In library.go registerSymbols():
purego.RegisterLibFunc(&jphCylinderShapeCreate, handle, "JPH_CylinderShape_Create")

// In shapes.go:
func NewCylinderShape(halfHeight, radius float32) *Shape {
    h := jphCylinderShapeCreate(halfHeight, radius)
    return &Shape{handle: h}
}
```

## Limitations

- **Native dependency**: requires the joltc shared library at runtime.
- **No double precision**: only single-precision joltc builds are supported.
- **Partial API coverage**: this wrapper covers core functionality (bodies, shapes, simulation stepping). Additional features (constraints, characters, raycasting, etc.) can be added following the same pattern.
- **Thread safety**: Go wrapper types are not safe for concurrent use without external synchronization.

## License

[MIT](LICENSE)

## Acknowledgments

- [JoltPhysics](https://github.com/jrouwe/JoltPhysics) by Jorrit Rouwe
- [joltc](https://github.com/amerkoleci/joltc) by Amer Koleci — C API for JoltPhysics
- [purego](https://github.com/ebitengine/purego) by Ebitengine — Go↔C without cgo
