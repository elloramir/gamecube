package main

import (
	// "fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var keyStates [glfw.KeyLast]bool
var mouseX, mouseY float64
var lastMouseX, lastMouseY float64

func KeyboardCallback(_ *glfw.Window, key glfw.Key, _ int, action glfw.Action, _ glfw.ModifierKey) {
	keyStates[key] = (action == glfw.Press || action == glfw.Repeat)
}

func MouseCallback(_ *glfw.Window, xPos float64, yPos float64) {
	lastMouseX, lastMouseY = mouseX, mouseY 
	mouseX, mouseY = xPos, yPos 
}

func IsKeyPressed(key glfw.Key) bool {
	return keyStates[key]
}

func MousePosition() (float32, float32) {
	return float32(mouseX), float32(mouseY)
}

func MouseRelativePosition() (float32, float32) {
	return float32(lastMouseX), float32(lastMouseY)
}