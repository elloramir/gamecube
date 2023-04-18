package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

var keyStates [glfw.KeyLast]bool
var mouseX, mouseY float32
var mouseRelativeX, mouseRelativeY float32

func InputUpdate() {
	mouseRelativeX, mouseRelativeY = 0, 0
}

func KeyboardCallback(_ *glfw.Window, key glfw.Key, _ int, action glfw.Action, _ glfw.ModifierKey) {
	keyStates[key] = (action == glfw.Press || action == glfw.Repeat)
}

func MouseCallback(_ *glfw.Window, xPos float64, yPos float64) {
	mouseRelativeX = float32(xPos)-mouseX
	mouseRelativeY = float32(yPos)-mouseY

	// save current mouse position
	mouseX = float32(xPos) 
	mouseY = float32(yPos) 
}

func IsKeyDown(key glfw.Key) bool {
	return keyStates[key]
}

func MousePosition() (float32, float32) {
	return mouseX, mouseY
}

func MouseRelativePosition() (float32, float32) {
	return mouseRelativeX, mouseRelativeY
}