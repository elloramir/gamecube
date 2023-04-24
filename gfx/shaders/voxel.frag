#version 330

in vec2 texCoords;

out vec4 color;

uniform sampler2D sprite;

void main() {
	color = texture(sprite, texCoords);
}
