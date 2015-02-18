package tracing

import (
    "math"
    "image/color"
    "github.com/lethal-guitar/go_tracer/vecmath"
    "github.com/lethal-guitar/go_tracer/spatial"
    "github.com/lethal-guitar/go_tracer/scene"
)

// Data model for scene description
type Scene struct {
    Traceables []*scene.Traceable
    Lights []*scene.LightSource

    BackgroundColor color.RGBA
}

func (self *Scene) AddObject(object scene.Traceable) {
    self.Traceables = append(self.Traceables, &object)
}

func (self *Scene) AddLight(light scene.LightSource) {
    self.Lights = append(self.Lights, &light)
}

func (self *Scene) bounds() scene.AABB {
    INFINITY := float32(math.Inf(1))

    bounds := scene.MakeAABBV(
        vecmath.MakeVec3d(INFINITY),
        vecmath.MakeVec3d(-INFINITY),
    )
    for _, object := range self.Traceables {
        objBounds := (*object).Bounds()
        bounds.Expand(&objBounds)
    }

    return bounds
}

func (self *Scene) ToOctree() spatial.Octree {
    tree := spatial.MakeOctree(self.bounds())

    for _, object := range self.Traceables {
        tree.InsertObject(object)
    }

    return tree
}
