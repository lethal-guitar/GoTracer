package scene

import (
    "github.com/lethal-guitar/go_tracer/vecmath"
)

// An infinite plane
type Plane struct {
    TraceableBase

    DistanceToOrigin float32
    PlaneNormal vecmath.Vec3d

    // Due to the plane being infinite, it's impossible (or rather, it's
    // pointless) to determine its bounding box. So we allow specifying one
    // instead.
    AssignedBounds AABB
}

func MakePlane(distance float32, normal vecmath.Vec3d) *Plane {
    return &Plane{PlaneNormal: normal, DistanceToOrigin: distance}
}

func (self *Plane) CheckIntersection(ray *Ray) (bool, float32) {
    a := self.PlaneNormal.Dot(ray.Direction)
    if a >= 0.0 {
        return false, 0.0
    }

    b := self.PlaneNormal.Dot(ray.Origin) + self.DistanceToOrigin
    t := -(b) / a

    return t > 0.0, t
}

func (self *Plane) Normal(intersection *vecmath.Vec3d) vecmath.Vec3d {
    return self.PlaneNormal
}

func (self *Plane) Bounds() AABB {
    return self.AssignedBounds
}
