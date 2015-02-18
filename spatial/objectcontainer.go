package spatial

import (
    "math"
    "github.com/lethal-guitar/go_tracer/scene"
)

const MAX_SPHERES_PER_CONTAINER = 64


// Back end for spatial data structures: A flat list of traceables, that allows
// checking against rays.
// When used directly as tracing Engine, you get the classic "naive ray tracing"
// algorithm.
type ObjectContainer struct {

    // Cache-friendly list of sphere intersection data
    SphereInfos [MAX_SPHERES_PER_CONTAINER]scene.SphereIntersectionInfo

    // Less frequently accessed sphere data. This is assumed to have the
    // same length and order as SphereInfos, so that the remaining sphere
    // data corresponding to a certain intersection info can be retrieved using
    // an index into SphereInfos
    SphereRemainders []*scene.TraceableBase

    // Non-spheres (not optimized)
    Objects []*scene.Traceable
}

func (self *ObjectContainer) Insert(object *scene.Traceable) {
    sphere, ok := (*object).(*scene.Sphere)
    if ok {
        self.SphereInfos[len(self.SphereRemainders)] = sphere.SphereIntersectionInfo

        self.SphereRemainders = append(self.SphereRemainders, &sphere.TraceableBase)
    } else {
        self.Objects = append(self.Objects, object)
    }
}

// Due to the sphere data being stored in two separate lists, we need to
// reassemble the original sphere once we need to return it
func (self *ObjectContainer) reassembleSphereAt(index int) *scene.Sphere {
    return &scene.Sphere{
        SphereIntersectionInfo: self.SphereInfos[index],
        TraceableBase: *self.SphereRemainders[index],
    }
}


func (self *ObjectContainer) FindClosestObject(ray *scene.Ray) (*scene.Traceable, float32) {
    var minDistance float32 = float32(math.Inf(1))
    var foundObject *scene.Traceable

    for i:=0; i<len(self.SphereRemainders); i++ {
        intersected, distance := self.SphereInfos[i].CheckIntersection(ray)
        if intersected && distance < minDistance {
            minDistance = distance
            t := scene.Traceable(self.reassembleSphereAt(i))
            foundObject = &t
        }
    }

    for _, object := range self.Objects {
        intersected, distance := (*object).CheckIntersection(ray)

        if intersected && distance < minDistance {
            minDistance = distance
            foundObject = object
        }
    }

    return foundObject, minDistance
}

func (self *ObjectContainer) FindOccluder(ray *scene.Ray, lightDistance float32) bool {
    for i:=0; i<len(self.SphereRemainders); i++ {
        intersected, distance := self.SphereInfos[i].CheckIntersection(ray)
        if intersected && distance < lightDistance {
            return true
        }
    }

    for _, object := range self.Objects {
        intersected, distance := (*object).CheckIntersection(ray)
        if intersected && distance < lightDistance {
            return true
        }
    }
    return false
}
