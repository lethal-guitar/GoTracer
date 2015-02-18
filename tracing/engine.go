package tracing

import "github.com/lethal-guitar/go_tracer/scene"

type Engine interface {

    // Return closest object intersecting ray if any, nil otherwise.
    // If there is an intersection, also return distance to ray origin
    FindClosestObject(ray *scene.Ray) (*scene.Traceable, float32)

    // Return whether ray hits any object. This is an optimization for the
    // shadow calculations, where the search can be aborted as soon as any
    // occluding object is found
    FindOccluder(ray *scene.Ray, lightDistance float32) bool
}
