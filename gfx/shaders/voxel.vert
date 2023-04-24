#version 330

layout (location = 0) in vec3 position;
layout (location = 1) in vec3 normals;
layout (location = 2) in vec2 uv;

out vec2 texCoords;

uniform mat4 transform;
uniform vec3 offset;

void main() {
	texCoords = uv;
	gl_Position = transform * vec4(position+offset, 1.0);
}
