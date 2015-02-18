package main

import (
    "flag"
    "image/color"
    "image/png"
    "log"
    "os"
    "github.com/lethal-guitar/go_tracer/tracing"
    "github.com/lethal-guitar/go_tracer/scene"
    "github.com/lethal-guitar/go_tracer/vecmath"
)

func main() {
    tileSize := flag.Int("tileSize", 64, "Size of a render tile")
    flag.Parse()

    rayTracer := createTracer()
    rendering := rayTracer.Render(*tileSize)

    file, err := os.Create("out.png")
    if err != nil {
        log.Fatal(err)
    }
    if err = png.Encode(file, rendering); err != nil {
        log.Fatal(err)
    }
}

func addAbletonLogo(
    model *tracing.Scene,
    material *scene.Material,
    start,
    startY,
    startZ float32,
) {
    const RADIUS = 2
    const STEP = 5.3

    for y := 0; y < 5; y++ {
        for x := 0; x < 4; x++ {
            xPos := start + STEP*float32(x)
            yPos := startY + float32(4*y)
            model.AddObject(scene.MakeSphere(vecmath.Vec3d{xPos, yPos, startZ}, RADIUS, *material))
        }
    }

    start += 4*STEP
    for y := 0; y < 4; y++ {
        for x := 0; x < 5; x++ {
            xPos := start + float32(4*x)
            yPos := startY + float32(y)*STEP
            model.AddObject(scene.MakeSphere(vecmath.Vec3d{xPos, yPos, startZ}, RADIUS, *material))
        }
    }
}

func setupScene() tracing.Scene {
    redMaterial := scene.MakeSimpleMaterial(color.RGBA{255, 0, 0, 255})
    greenMaterial := scene.MakeSimpleMaterial(color.RGBA{0, 255, 0, 255})
    greenMaterial.Specularity = 50
    greenMaterial.Reflectivity = 0.22

    transparent := scene.MakeSimpleMaterial(color.RGBA{255, 255, 255, 255})
    transparent.Transparency = 1.0
    transparent.RefractionIndex = 1.49

    grayMaterial := scene.MakeSimpleMaterial(color.RGBA{100, 100, 100, 255})
    grayMaterial.Specular = scene.FloatColor{}
    grayMaterial2 := scene.MakeSimpleMaterial(color.RGBA{100, 100, 100, 255})
    grayMaterial.Specular = scene.FloatColor{}
    grayMaterial2.Reflectivity = 0.05

    mirrorMaterial := scene.MakeSimpleMaterial(color.RGBA{255, 255, 255, 255})
    mirrorMaterial.Reflectivity = 0.75

    model := tracing.Scene{BackgroundColor: color.RGBA{100, 100, 100, 255}}

    for i := 0; i < 5; i++ {
        addAbletonLogo(&model, &mirrorMaterial, -18, -2.0, 230 + float32(i)*20)
    }

    model.AddObject(scene.MakeSphere(vecmath.Vec3d{-16.0, -11.0, 205.0}, 5.0, redMaterial))
    model.AddObject(scene.MakeSphere(vecmath.Vec3d{10.0, 13.0, 190.0}, 3.0, greenMaterial))
    model.AddObject(scene.MakeSphere(vecmath.Vec3d{20.0, 0.0, 205.0}, 5.0, greenMaterial))

    floor := scene.MakePlane(16, vecmath.Vec3d{0,-1,0})
    floor.ObjectMaterial = grayMaterial2
    floor.AssignedBounds = scene.MakeAABB(-28, 16, -50, 28, 16.1, 320)

    leftWall := scene.MakePlane(24, vecmath.Vec3d{1,0,0})
    leftWall.ObjectMaterial = grayMaterial
    leftWall.AssignedBounds = scene.MakeAABB(-24, -80, -50, -23.9, 80, 320)

    backWall := scene.MakePlane(320, vecmath.Vec3d{0,0,-1})
    backWall.ObjectMaterial = grayMaterial
    backWall.AssignedBounds = scene.MakeAABB(-28, -80, 320, 28, 80, 320.1)

    rightWall := scene.MakePlane(28, vecmath.Vec3d{-1,0,0})
    rightWall.ObjectMaterial = grayMaterial
    rightWall.AssignedBounds = scene.MakeAABB(28, -80, -50, 28.1, 80, 320)

    model.AddObject(floor)

    white := color.RGBA{255, 255, 255, 255}
    dimWhite := color.RGBA{100, 100, 100, 255}

    model.AddLight(
        scene.MakeColoredLight(vecmath.Vec3d{0.0, -30.0, 210.0}, white))
    model.AddLight(
        scene.MakeColoredLight(vecmath.Vec3d{26.0, -80.0, 80.0}, white))
    model.AddLight(
        scene.MakeColoredLight(vecmath.Vec3d{0.0, 0.0, -10.0}, dimWhite))
    return model
}

func createTracer() tracing.RayTracer {
    // A4 600 DPI: 4961 x 7016

    // A4 300 DPI
    //width := 3508
    //height := 2480

    // Full HD
    //width := 1920
    //height := 1080

    width := 1280
    height := 1280

    // NOTE: For now, the target resolution's aspect ratio must be square,
    // otherwise the view frustum wouldn't be right.
    scene := setupScene()
    tracer := tracing.RayTracer{
        Width: width,
        Height: height,
        Camera: tracing.CameraConfig{
            -1.0, 1.0,
            -1.0, 1.0,
            vecmath.Vec3d{0.0, 0.0, -10.0},
        },
        Scene: &scene,
    }

    return tracer
}
