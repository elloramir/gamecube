#version 330

out vec4 outputColor;

in vec2 uv;

uniform sampler2D tex;

void main() {
    // outputColor = texture(tex, vec2(uv.x, 1 - uv.y));
    outputColor = vec4(uv.x, uv.y, 0.0, 1.0);
}