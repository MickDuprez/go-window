## Loading an image
In this sections we'll load a classic image and display it.
We'll also add a line to the same frame buffer that draws over the image as expected.
This will be handy for loading textures at a later date I think.

I've also added some timeing code to measure how milliseconds per frame per second.
An ideal target is around 16.666 which equates to 60 frames per second.
I'm not happy with this yet though and it will nedd more work.

There's also a sleep line of code to make the loop sleep for 16ms (60fps) to try
and set a base speed for the loop. This is placed after all rendering is done although
it probably doesn't matter that much.

### Sample of output
![sample output](https://github.com/MickDuprez/go-window/blob/master/04 - Loading an image/doom.JPG)


