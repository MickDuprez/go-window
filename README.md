# go-window

### Creating an empty gui window in go for rendering 2D/3D graphics
The goal is to create a bare bones window for any operating system (OS) with basic event handlers 
and graphics buffer/s ready to draw pixels to. It can be used to create simple 2D/3D graphics 
from 'absolute' scratch (i.e. pixels!) for games or image viewing and editing etc.
Yes, there definitely libraries out there ready to go for building games or GUI's but they abstract
away all the very things I want to learn.

Creating a window and getting to the 'buffer' is more trivial in other languages like Java, Python, C# C/C++ etc 
and it's a well troden path. Go however, being more focused on systems programming rather than desktop 
GUI applications, is proving to be a lot harder.
The GoDoc's, while very good have some unusual terminologies for things if you come from say a win32 background
and can seem more confusing than enlightening so I thought I'd document my findings as I build up this basic 
windowing library from as low a level as possible.
I'm also learning Go so this is as good a project as any to stretch the old grey cells with and I welcome any
comments/suggestions/corrections :)

### What this project won't or can't provide
Due to limitations in the go libraries and my lack of knowledge of said lib's there will
be no bells and whistles like title bar icons, message boxes, UI widgets and other things you
would require for a 'proper' desktop application. For these I'd recomend one of the Go wrapped
GUI frameworks such as QT, GTK or WxWidgets et al.
This project is built purely for low overhead with good performance and to develop and display 
graphics using software rendering in Go so I wanted this to be as lean as possible.

### What about OpenGL/DirectX et al?
The whole purpose of this windowing library is for learning and experimentation and to discover how these
graphics libraries work under the hood. While combing the internets, a questions such as "I want to create
a graphics engine from scratch" is usually replied with "don't waste your time, that tech is dead" and other
such un-inspiring comments but I think there is real value in learning these things from the bottom up.
If you're reading this, you love go and you want to learn graphics, hopefully this repo will be of some use.

## Repo Structure and Usage
This repo is structured in lesson like folders that start from displaying a bare bones empty window with
very basic low level events up to a fully working example library ready to go with all of the previous steps
being abstracted away and ready to use.

Study the README in each folder for some more detailed explanations and observations I had while coding
up the source.

## Important!
For futher documentation, see the imported lib's and look for them in GoDoc, particulalry golang.org/x/exp/shiny,
golang.org/x/exp/shiny/screen and mobile/event to get started with. 
Reading them along while reading this code will be very helpful!
NOTE: This repo is being developed on a Windows OS, as I get time I will confirm/update differences
I find with other OS's



