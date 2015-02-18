package spatial

import (
    "math"
    "github.com/lethal-guitar/go_tracer/scene"
)

const MAX_OCTREE_DEPTH = 2

// A very naive and unoptimized Octree. This actually uses an AABB per node -
// normally you wouldn't do that, since there are Octree-specific ray
// intersection testing/walking algorithms. The placement heuristics are also
// super simple and not tweaked at all.
//
// I'm planning to replace this with a Kd-Tree anyway
type Octree struct {
    Root OctreeNode
}

func MakeOctree(bounds scene.AABB) Octree {
    return Octree{Root: *MakeNode(bounds)}
}

func (self *Octree) InsertObject(object *scene.Traceable) {
    bounds := (*object).Bounds()
    if bounds.Overlaps(&self.Root.Bounds) {
        self.place(object, &self.Root, 0)
    }
}

func (self *Octree) place(object *scene.Traceable, node *OctreeNode, depth int) {
    bounds := (*object).Bounds()

    if depth >= MAX_OCTREE_DEPTH {
        node.AddObject(object)
        return
    }

    if node.IsLeaf() {
        node.Split()
    }

    for _, child := range node.Children {
        if child.Bounds.Contains(&bounds) {
            self.place(object, child, depth+1)
            return
        }

        if child.Bounds.Overlaps(&bounds) {
             self.place(object, child, depth+1)
        }
    }
}

func (self *Octree) findObject(ray *scene.Ray, node *OctreeNode) (
    *scene.Traceable, float32,
) {
    return node.Objects.FindClosestObject(ray)
}

func (self *Octree) checkShadowObject(
    ray *scene.Ray,
    lightDistance float32,
    node *OctreeNode,
) bool {
    return node.Objects.FindOccluder(ray, lightDistance)
}

func (self *Octree) FindClosestObject(ray *scene.Ray) (
    foundObject *scene.Traceable, distance float32,
) {
    return self.findClosest(ray, &self.Root)
}

func (self *Octree) FindOccluder(ray *scene.Ray, lightDistance float32) bool {
    return self.findForShadow(ray, lightDistance, &self.Root)
}

func (self *Octree) findClosest(ray *scene.Ray, node *OctreeNode) (
    foundObject *scene.Traceable, distance float32,
) {
    if node.Bounds.IntersectsBasic(ray) {
        foundObject, distance = self.findObject(ray, node)

        if foundObject == nil {
            distance = float32(math.Inf(1))
        }

        if !node.IsLeaf() {
            for _, child := range node.Children {
                subFound, subDist := self.findClosest(ray, child)
                if subFound != nil && subDist < distance {
                    foundObject = subFound
                    distance = subDist
                }
            }
        }
    }
    return
}

func (self *Octree) findForShadow(
    ray *scene.Ray,
    lightDistance float32,
    node *OctreeNode,
) bool {
    if node.Bounds.Intersects(ray, lightDistance) {
        foundObject := self.checkShadowObject(ray, lightDistance, node)

        if foundObject {
            return true
        }

        if !node.IsLeaf() {
            for _, child := range node.Children {
                subFound := self.findForShadow(ray, lightDistance, child)
                if subFound {
                    return true
                }
            }
        }
    }
    return false
}
