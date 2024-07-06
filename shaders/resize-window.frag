#version 330 core
out vec4 FragColor;
  
uniform float uTime;
uniform vec2 screenResolution;

void main()
{
    vec2 uv = gl_FragCoord.xy/screenResolution;
    uv.x=uv.x+sin(uTime)/2;
    uv.y=uv.y+cos(uTime)/2;
    FragColor = vec4(uv,uv.x+uv.y/2,1.0);
}  