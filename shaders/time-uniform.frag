#version 330 core
out vec4 FragColor;
  
uniform float uTime;

void main()
{
    vec2 uv = gl_FragCoord.xy/vec2(800,600);
    uv.x=uv.x+sin(uTime)/2;
    uv.y=uv.y+cos(uTime)/2;
    FragColor = vec4(uv,uv.x+uv.y/2,1.0);
}  