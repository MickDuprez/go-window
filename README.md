# go-window

### Creating an empty gui window in go for software rendering of 2D/3D graphics
This project creates a bare bones GUI window for any Windows/Linux/Mac with basic event handlers 
and graphics buffer/s ready to draw pixels to. It can be used to create simple 2D/3D graphics 
from 'absolute' scratch (i.e. pixels!) for games or image viewing and editing etc.
Yes, there definitely libraries out there ready to go for building games or GUI's but they abstract
away all the very things I want to learn.

The project uses libraries from golang.org/x/exp/shiny and golang.org/x/mobile/event and produce a
blank native styled window with title bar and menu, there are no options for border styles etc.
The GoDoc's, while very good have some unusual terminologies for things if you come from say a win32 background
and can seem more confusing than enlightening so I thought I'd document my findings as I build up this basic 
windowing library from as low a level as possible.
I'm also learning Go so this is as good a project as any to stretch the old grey cells with and I welcome any
comments/suggestions/corrections :)

### Sample of output
![sample output](https://github.com/MickDuprez/go-window/blob/master/go-window_03.JPG)

### What this project won't or can't provide
Due to limitations in the go libraries and my lack of knowledge of said lib's there will
be no bells and whistles like title bar icons, message boxes, UI widgets and other things you
would require for a 'proper' desktop application. For these I'd recomend one of the Go wrapped
GUI frameworks such as QT, GTK or WxWidgets et al.
This project is built purely for low overhead with good performance and to develop and display 
graphics using software rendering in Go so I wanted this to be as lean as possible.

### What about OpenGL/DirectX et al?
The whole purpose of this windowing library is for learning and experimentation and to discover how these
graphics libraries work under the hood. While there are good arguments for using these driver API's it's still good
to know how to do basic graphics at the lower level, not only for learning but for things like micro controllers
with a simple LCD screen say to build an old style arcade cabinet or an equipment monitoring console etc.
Anyway, if you're reading this and you love Go and want to learn graphics, hopefully this repo will be of some use.

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

__NOTE: This repo is being developed on a Windows OS, as I get time I will confirm/update differences
I find with other OS's__



