// Copyright (c) 2023 Ellora.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package gfx

import (
	"embed"
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

//go:embed shaders/*
var shadersFS embed.FS

func LoadShader(vertexFilename, fragmentFilename string) (uint32, error) {
	data, err := shadersFS.ReadFile(vertexFilename)
	if err != nil {
		return 0, err
	}
	vertexSource := string(data)
	data, err = shadersFS.ReadFile(fragmentFilename)
	if err != nil {
		return 0, err
	}
	fragmentSource := string(data)

	return newProgram(vertexSource, fragmentSource)
}

func newProgram(vertexSource, fragmentSource string) (uint32, error) {
	// Compule shaders
	vertexShader, err := newShader(vertexSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}
	fragmentShader, err := newShader(fragmentSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	// Create program
	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	// Link result
	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		// Log size
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)
		// Log content
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
}

func newShader(source string, shaderType uint32) (uint32, error) {
	// Create shader
	shader := gl.CreateShader(shaderType)
	csources, free := gl.Strs(source + "\x00")
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	// Compile result
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		// Log size
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		// Log content
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
