#version 330 core
layout (location = 0) in vec2 aPos;   // the position variable has attribute position 0
layout (location = 1) in vec3 aColor; // the color variable has attribute position 1
  
out vec3 ourColor; // output a color to the fragment shader
// uniform float uOffset;
// out vec3 pos;

void main()
{
    gl_Position = vec4(aPos,0.0, 1.0);
    // gl_Position = vec4(aPos.x+uOffset,aPos.y,0.0, 1.0);
    // gl_Position = vec4(-aPos.x,-aPos.y,0.0, 1.0);
    ourColor = aColor; // set ourColor to the input color we got from the vertex data
    // pos=vec3(aPos,0.0);
}