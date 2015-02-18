package scene

import (
    "github.com/lethal-guitar/go_tracer/vecmath"
)

// Thing that can be intersected with rays and thus rendered via
// ray tracing
type Traceable interface {

    // Do we intersect ray? If so, return distance from ray origin to
    // point of intersection
    CheckIntersection(ray *Ray) (bool, float32)

    // Surface normal at given intersection point
    Normal(intersection *vecmath.Vec3d) vecmath.Vec3d

    Material() *Material

    Bounds() AABB
}

// Base class for traceables
type TraceableBase struct {
    ObjectMaterial Material
}

func (self *TraceableBase) Material() *Material {
    return &self.ObjectMaterial
}
