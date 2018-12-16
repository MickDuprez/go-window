## Creating the screen and pixel buffers
This is as bare bones as you can get to have a simple window to draw pixels to.

### The Process - Get Screen Buffer -> Get a Pixel Buffer -> Draw, Upload and Publish
Once you have set up the main window you need to get one or more screen Buffers.
From what I can tell, the screen buffer is what the window calls to display the image.
The screen buffer/s can be of arbitrary sizes and could be arranged to fit within
the main window to create multiple panels of different graphics buffers. These can be
used for things like different views or scales of the same image.
The screen buffer/s needs to be renewed on resizing of the main window to resize itself.
there doesn't appear to be any methods to resize it which makes sense as it would
have to return a new buffer anyway to organise memory.
There is also a Texture you can grab from the screen which I haven't looked into yet.

You can then get a pixel buffer (*image.RGBA) from the screen buffer, this will be the same size as
the screen buffer until a resize is called so you will need a new one of these as well on resize.
This is where you draw your pixels to.
All x, y values start at the top left corner of the pixel buffer and are positive in the right
and down directions.

Once you are done drawing pixels you tell the window to upload the screen buffer and publish
the results.

So far this is working well but I would love any feedback on how to do this better. The examples
mentioned in the doc's are a bit ambitious to take in just to get a basic window and they use
different libraries than used here.


