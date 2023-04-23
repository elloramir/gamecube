#version 330

out vec4 outputColor;
in vec2 inTexCoord;

uniform sampler2D sprite;

void main() {
	outputColor = texture(sprite, inTexCoord);
}
