#version 330

layout (location = 0) in vec3 positionAttribute;
layout (location = 1) in vec2 texCoordAttribute;

uniform mat4 transform;

out vec2 inTexCoord;

void main() {
	inTexCoord = texCoordAttribute;
	gl_Position = transform * vec4(positionAttribute, 1.0);
}
