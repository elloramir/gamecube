#version 330

layout (location = 0) in vec3 positionAttribute;

uniform mat4 transform;

out vec3 inPos;

void main() {
	inPos = positionAttribute;
	gl_Position = transform * vec4(positionAttribute, 1.0);	
}