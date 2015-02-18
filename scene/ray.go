package scene

import "github.com/lethal-guitar/go_tracer/vecmath"

type Ray struct {
    Origin vecmath.Vec3d
    Direction vecmath.Vec3d
}

func (self *Ray) PointAt(distance float32) vecmath.Vec3d {
    vector := self.Direction.Scaled(distance)
    return self.Origin.Added(vector)
}
