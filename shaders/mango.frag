#version 330 core
out vec4 FragColor;


void main() {
	vec2 st = gl_FragCoord.xy/vec2(800,600);
	FragColor = vec4(st.x,st.y,0.0,1.0);
}

