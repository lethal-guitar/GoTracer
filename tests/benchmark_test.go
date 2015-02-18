package tests

import (
    "testing"
    "github.com/lethal-guitar/go_tracer/scene"
    "github.com/lethal-guitar/go_tracer/vecmath"
)

func makeExampleSphere() *scene.Sphere {
    return scene.MakeSimpleSphere(vecmath.Vec3d{-19.0, 12.0, 205.0}, 5.0)
}

func BenchmarkSphereRayIntersectionHitting(b *testing.B) {
    sphere := makeExampleSphere()
    ray := scene.Ray{
        vecmath.Vec3d{0, 0, -10},
        vecmath.Vec3d{-0.09082417, 0.032597166, 0.9953333},
    }

    for n := 0; n < b.N; n++ {
        sphere.CheckIntersection(&ray)
    }
}

func BenchmarkSphereRayIntersectionMissing(b *testing.B) {
    sphere := makeExampleSphere()
    ray := scene.Ray{
        vecmath.Vec3d{0, 0, -10},
        vecmath.Vec3d{-0.09901476, -0.09901476, 0.9901476},
    }

    for n := 0; n < b.N; n++ {
        sphere.CheckIntersection(&ray)
    }
}


func BenchmarkAABBRayIntersectionHitting(b *testing.B) {
    sphere := makeExampleSphere()
    bounds := sphere.Bounds()
    ray := scene.Ray{
        vecmath.Vec3d{0, 0, -10},
        vecmath.Vec3d{-0.09082417, 0.032597166, 0.9953333},
    }

    for n := 0; n < b.N; n++ {
        bounds.IntersectsBasic(&ray)
    }
}

func BenchmarkAABBRayIntersectionMissing(b *testing.B) {
    sphere := makeExampleSphere()
    ray := scene.Ray{
        vecmath.Vec3d{0, 0, -10},
        vecmath.Vec3d{-0.09901476, -0.09901476, 0.9901476},
    }

    bounds := sphere.Bounds()

    for n := 0; n < b.N; n++ {
        bounds.IntersectsBasic(&ray)
    }
}
