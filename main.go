package main

import (
	"log"
	"runtime"
	_"embed"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v3.3-core/gl"
)

var (
	screenWidth = 800
	screenHeight = 600
)

var (
	//go:embed res/game.vert
	voxelVertexSource string
	//go:embed res/game.frag
	voxelFragmentSource string
)

func init() {
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		log.Fatal("can't initialize GLFW")
	}
	defer glfw.Terminate()

	// create window
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	window, err := glfw.CreateWindow(screenWidth, screenHeight, "", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	window.MakeContextCurrent()
	//window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	//window.SetKeyCallback(KeyboardCallback)
	//window.SetCursorPosCallback(MouseCallback)

	glfw.SwapInterval(1) // vsync on

	// loading opengl
	err = gl.Init()
	if err != nil {
		log.Fatal(err)
	}

	// program startup
	log.Println("welcome adventurer, to the 'gamecube'")
	log.Println(gl.GoStr(gl.GetString(gl.VERSION)))

	// creating GPU objects to hold our data
	var vao, vbo, ebo uint32

	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &ebo)

	// creating our first face
	var x, y, z float32 = 0, 0, 0

	vertices := []float32{
		x + 0, y + 0, z, // top left
		x + 1, y + 0, z, // top right
		x + 1, y - 1, z, // bottom right
		x + 0, y - 1, z} // bottom left


	// first triangle
    // 0 ------ 1
	// |         
	// 2        3

	// second triangle
	// 0        1
	//          |
	// 2 ------ 3

	triangles := []uint32{
		0, 1, 3,
		1, 2, 3}

	voxelMesh := Mesh{}
	voxelMesh.upload(vertices, triangles)

	gl.BindVertexArray(0)

	// load our programs/shaders
	voxelProgram, err := genProgram(voxelVertexSource, voxelFragmentSource)
	if err != nil {
		log.Fatal(err)
	}

	camera := newCamera(0, 0, 1)

	for !window.ShouldClose() {
		glfw.PollEvents()

		camera.sendUniforms(voxelProgram)
		gl.UseProgram(voxelProgram)
		voxelMesh.render()

		window.SwapBuffers()
	}
}