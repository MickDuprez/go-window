# go-window
### Creating an empty gui window in go for rendering 2D/3D graphics

The goal is to create a bare bones window for any operating system (OS) with basic event handlers 
and graphics buffer/s ready to draw pixels to to create simple 2D/3D graphics 
from 'absolute' scratch (i.e. pixels!) for games or image viewing and editing etc.
Yes, there definitely libraries out there ready to go for building games or GUI's but they abstract
away all the very things I want to learn.

Creating a window and getting the 'buffer' is more trivial in other languages like Java, Python, C# C/C++ etc 
and it's a well troden path. Go however, being more focused on systems programming rather than desktop 
GUI applications, seems to be a lot harder.
The GoDoc's, while very good have some unusual terminologies for things if you come from say a win32 background
and can seem more confusing the enlightening so I thought I'd document my findings as I build up this basic 
windowing library from as low a level as possible.
I'm also learning Go so this is as good a project as any to stretch the old grey cells with :)

### What about OpenGL/DirectX et al?
The whole purpose of this windowing library is for learning and experimentation and to discover how these
graphics libraries work under the hood.

## Repo Structure and Usage
This repo is structured in lesson like folders that start from displaying a bare bones empty window with
very basic low level events up to a fully working example library ready to go with all of the previous steps
being abstracted away and ready to use.

Study the README in each folder for some more detailed explanations and observations I had while coding
up the source.



