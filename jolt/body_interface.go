package jolt

// BodyInterface provides methods for creating, adding, removing, and
// manipulating physics bodies. It is obtained from PhysicsSystem.GetBodyInterface().
//
// The BodyInterface does not own any resources itself—it is a view into the
// PhysicsSystem and remains valid for the PhysicsSystem's lifetime.
type BodyInterface struct {
	handle uintptr
}

// CreateAndAddBody creates a new body from the given settings and immediately
// adds it to the physics world. Returns the BodyID of the new body.
func (bi *BodyInterface) CreateAndAddBody(settings *BodyCreationSettings, activation Activation) BodyID {
	id := jphBodyInterfaceCreateAndAddBody(bi.handle, settings.handle, int32(activation))
	return BodyID(id)
}

// RemoveAndDestroyBody removes a body from the simulation and destroys it.
func (bi *BodyInterface) RemoveAndDestroyBody(bodyID BodyID) {
	jphBodyInterfaceRemoveAndDestroyBody(bi.handle, uint32(bodyID))
}

// RemoveBody removes a body from the simulation without destroying it.
func (bi *BodyInterface) RemoveBody(bodyID BodyID) {
	jphBodyInterfaceRemoveBody(bi.handle, uint32(bodyID))
}

// DestroyBody destroys a body that has already been removed from the simulation.
func (bi *BodyInterface) DestroyBody(bodyID BodyID) {
	jphBodyInterfaceDestroyBody(bi.handle, uint32(bodyID))
}

// IsAdded returns whether a body is currently in the simulation.
func (bi *BodyInterface) IsAdded(bodyID BodyID) bool {
	return jphBodyInterfaceIsAdded(bi.handle, uint32(bodyID))
}

// SetLinearVelocity sets the linear velocity of a body.
func (bi *BodyInterface) SetLinearVelocity(bodyID BodyID, velocity Vec3) {
	jphBodyInterfaceSetLinearVelocity(bi.handle, uint32(bodyID), &velocity)
}

// GetLinearVelocity returns the linear velocity of a body.
func (bi *BodyInterface) GetLinearVelocity(bodyID BodyID) Vec3 {
	var v Vec3
	jphBodyInterfaceGetLinearVelocity(bi.handle, uint32(bodyID), &v)
	return v
}

// GetCenterOfMassPosition returns the center-of-mass world position of a body.
func (bi *BodyInterface) GetCenterOfMassPosition(bodyID BodyID) Vec3 {
	var pos Vec3
	jphBodyInterfaceGetCenterOfMassPosition(bi.handle, uint32(bodyID), &pos)
	return pos
}

// SetPosition sets the world position of a body.
func (bi *BodyInterface) SetPosition(bodyID BodyID, position Vec3, activation Activation) {
	jphBodyInterfaceSetPosition(bi.handle, uint32(bodyID), &position, int32(activation))
}

// GetPosition returns the world position of a body.
func (bi *BodyInterface) GetPosition(bodyID BodyID) Vec3 {
	var pos Vec3
	jphBodyInterfaceGetPosition(bi.handle, uint32(bodyID), &pos)
	return pos
}

// SetRotation sets the rotation of a body.
func (bi *BodyInterface) SetRotation(bodyID BodyID, rotation Quat, activation Activation) {
	jphBodyInterfaceSetRotation(bi.handle, uint32(bodyID), &rotation, int32(activation))
}

// GetRotation returns the rotation of a body.
func (bi *BodyInterface) GetRotation(bodyID BodyID) Quat {
	var q Quat
	jphBodyInterfaceGetRotation(bi.handle, uint32(bodyID), &q)
	return q
}

// ActivateBody wakes a sleeping body.
func (bi *BodyInterface) ActivateBody(bodyID BodyID) {
	jphBodyInterfaceActivateBody(bi.handle, uint32(bodyID))
}

// DeactivateBody puts a body to sleep.
func (bi *BodyInterface) DeactivateBody(bodyID BodyID) {
	jphBodyInterfaceDeactivateBody(bi.handle, uint32(bodyID))
}

// IsActive returns whether a body is currently active (not sleeping).
func (bi *BodyInterface) IsActive(bodyID BodyID) bool {
	return jphBodyInterfaceIsActive(bi.handle, uint32(bodyID))
}

// AddForce adds a force (in Newtons) to the body's center of mass.
// The force is applied for the duration of the next simulation step.
func (bi *BodyInterface) AddForce(bodyID BodyID, force Vec3) {
	jphBodyInterfaceAddForce(bi.handle, uint32(bodyID), &force)
}

// AddImpulse applies an instantaneous impulse to the body's center of mass.
func (bi *BodyInterface) AddImpulse(bodyID BodyID, impulse Vec3) {
	jphBodyInterfaceAddImpulse(bi.handle, uint32(bodyID), &impulse)
}

// SetFriction sets the friction coefficient of a body.
func (bi *BodyInterface) SetFriction(bodyID BodyID, friction float32) {
	jphBodyInterfaceSetFriction(bi.handle, uint32(bodyID), friction)
}

// GetFriction returns the friction coefficient of a body.
func (bi *BodyInterface) GetFriction(bodyID BodyID) float32 {
	return jphBodyInterfaceGetFriction(bi.handle, uint32(bodyID))
}

// SetRestitution sets the restitution (bounciness) of a body.
func (bi *BodyInterface) SetRestitution(bodyID BodyID, restitution float32) {
	jphBodyInterfaceSetRestitution(bi.handle, uint32(bodyID), restitution)
}

// GetRestitution returns the restitution of a body.
func (bi *BodyInterface) GetRestitution(bodyID BodyID) float32 {
	return jphBodyInterfaceGetRestitution(bi.handle, uint32(bodyID))
}

// SetGravityFactor sets how much gravity affects a body (1.0 = normal).
func (bi *BodyInterface) SetGravityFactor(bodyID BodyID, factor float32) {
	jphBodyInterfaceSetGravityFactor(bi.handle, uint32(bodyID), factor)
}

// GetGravityFactor returns the gravity factor of a body.
func (bi *BodyInterface) GetGravityFactor(bodyID BodyID) float32 {
	return jphBodyInterfaceGetGravityFactor(bi.handle, uint32(bodyID))
}

// GetMotionType returns the motion type of a body.
func (bi *BodyInterface) GetMotionType(bodyID BodyID) MotionType {
	return MotionType(jphBodyInterfaceGetMotionType(bi.handle, uint32(bodyID)))
}

// SetMotionType changes the motion type of a body.
func (bi *BodyInterface) SetMotionType(bodyID BodyID, motionType MotionType, activation Activation) {
	jphBodyInterfaceSetMotionType(bi.handle, uint32(bodyID), int32(motionType), int32(activation))
}
