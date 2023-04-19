#version 330

layout (location = 0) in vec3 positionAttribute;

uniform mat4 transformMatrix;

void main() {
	gl_Position = transformMatrix * vec4(positionAttribute, 1.0);	
}