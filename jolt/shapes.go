package jolt

// Shape wraps an opaque C physics shape pointer.
// Shapes are reference-counted in joltc; once assigned to a body the body
// holds a reference. Call Destroy() only on shapes that have NOT been assigned
// to any body, or after all referencing bodies have been removed.
type Shape struct {
	handle uintptr
}

// Destroy releases the underlying C shape resource.
// Only call this on shapes that are no longer referenced by any body.
func (s *Shape) Destroy() {
	if s.handle != 0 {
		jphShapeDestroy(s.handle)
		s.handle = 0
	}
}

// NewBoxShape creates a box collision shape with the given half extents.
// convexRadius adds rounding to edges for smoother collision (use 0.05 as default).
func NewBoxShape(halfExtent Vec3, convexRadius float32) *Shape {
	h := jphBoxShapeCreate(&halfExtent, convexRadius)
	return &Shape{handle: h}
}

// NewSphereShape creates a sphere collision shape with the given radius.
func NewSphereShape(radius float32) *Shape {
	h := jphSphereShapeCreate(radius)
	return &Shape{handle: h}
}

// NewCapsuleShape creates a capsule collision shape.
// halfHeight is half the height of the cylindrical part, radius is the
// radius of the hemispherical caps and the cylinder.
func NewCapsuleShape(halfHeight, radius float32) *Shape {
	h := jphCapsuleShapeCreate(halfHeight, radius)
	return &Shape{handle: h}
}
