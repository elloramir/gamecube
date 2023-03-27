#version 330

layout (location = 0) in vec3 positionAttribute;
layout (location = 1) in vec2 uvAttribute;

uniform mat4 viewProjMatrix;

out vec2 uv;

void main() {
	uv = uvAttribute;
	gl_Position = viewProjMatrix * vec4(positionAttribute, 1.0);	
}