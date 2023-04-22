#version 330

out vec4 outputColor;
in vec3 inPos;

void main() {
	outputColor = vec4(mod(inPos.yxz/25.0, 1.0), 1.0);
}