package jolt

// BodyCreationSettings configures a new physics body before it is added
// to the simulation. Create one with NewBodyCreationSettings, adjust
// properties as needed, then pass it to BodyInterface.CreateAndAddBody.
//
// Close() must be called to free the underlying C resource if you are done
// with the settings without adding a body, or after the body has been added.
type BodyCreationSettings struct {
	handle uintptr
}

// NewBodyCreationSettings creates body creation settings from a shape,
// position, rotation, motion type, and collision layer.
func NewBodyCreationSettings(shape *Shape, position Vec3, rotation Quat, motionType MotionType, objectLayer ObjectLayer) *BodyCreationSettings {
	h := jphBodyCreationSettingsCreate3(
		shape.handle,
		&position,
		&rotation,
		int32(motionType),
		uint32(objectLayer),
	)
	return &BodyCreationSettings{handle: h}
}

// Close releases the underlying C body creation settings.
func (bcs *BodyCreationSettings) Close() {
	if bcs.handle != 0 {
		jphBodyCreationSettingsDestroy(bcs.handle)
		bcs.handle = 0
	}
}

// SetLinearVelocity sets the initial linear velocity.
func (bcs *BodyCreationSettings) SetLinearVelocity(v Vec3) {
	jphBodyCreationSettingsSetLinearVelocity(bcs.handle, &v)
}

// GetLinearVelocity returns the initial linear velocity.
func (bcs *BodyCreationSettings) GetLinearVelocity() Vec3 {
	var v Vec3
	jphBodyCreationSettingsGetLinearVelocity(bcs.handle, &v)
	return v
}

// SetFriction sets the friction coefficient.
func (bcs *BodyCreationSettings) SetFriction(v float32) {
	jphBodyCreationSettingsSetFriction(bcs.handle, v)
}

// GetFriction returns the friction coefficient.
func (bcs *BodyCreationSettings) GetFriction() float32 {
	return jphBodyCreationSettingsGetFriction(bcs.handle)
}

// SetRestitution sets the restitution (bounciness).
func (bcs *BodyCreationSettings) SetRestitution(v float32) {
	jphBodyCreationSettingsSetRestitution(bcs.handle, v)
}

// GetRestitution returns the restitution (bounciness).
func (bcs *BodyCreationSettings) GetRestitution() float32 {
	return jphBodyCreationSettingsGetRestitution(bcs.handle)
}

// SetGravityFactor sets how much gravity affects this body (1.0 = normal).
func (bcs *BodyCreationSettings) SetGravityFactor(v float32) {
	jphBodyCreationSettingsSetGravityFactor(bcs.handle, v)
}

// GetGravityFactor returns the gravity factor.
func (bcs *BodyCreationSettings) GetGravityFactor() float32 {
	return jphBodyCreationSettingsGetGravityFactor(bcs.handle)
}

// SetAllowSleeping controls whether the body is allowed to go to sleep.
func (bcs *BodyCreationSettings) SetAllowSleeping(v bool) {
	jphBodyCreationSettingsSetAllowSleeping(bcs.handle, v)
}

// GetAllowSleeping returns whether sleeping is allowed.
func (bcs *BodyCreationSettings) GetAllowSleeping() bool {
	return jphBodyCreationSettingsGetAllowSleeping(bcs.handle)
}

// SetMotionQuality sets the motion quality (discrete or linear cast).
func (bcs *BodyCreationSettings) SetMotionQuality(v MotionQuality) {
	jphBodyCreationSettingsSetMotionQuality(bcs.handle, int32(v))
}

// GetMotionQuality returns the motion quality setting.
func (bcs *BodyCreationSettings) GetMotionQuality() MotionQuality {
	return MotionQuality(jphBodyCreationSettingsGetMotionQuality(bcs.handle))
}
