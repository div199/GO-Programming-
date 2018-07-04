# GO-Programming-
This repository contains information and programs for those who want set their hands in GO programming language.


In order to install Go on windows visit this link " https://golang.org/doc/install?download=go1.10.3.windows-amd64.msi "
If you are upgrading from an older version of Go you must first remove the existing version.

Windows
MSI installer
Open the MSI file and follow the prompts to install the Go tools. By default, the installer puts the Go distribution in c:\Go.

The installer should put the c:\Go\bin directory in your PATH environment variable. You may need to restart any open command prompts for the change to take effect.

Setting environment variables under Windows
Under Windows, you may set environment variables through the "Environment Variables" button on the "Advanced" tab of the "System" control panel. Some versions of Windows provide this control panel through the "Advanced System Settings" option inside the "System" control panel.

Test your installation
Check that Go is installed correctly by setting up a workspace and building a simple program, as follows.

Create your workspace directory, %USERPROFILE%\go. (If you'd like to use a different directory, you will need to set the GOPATH environment variable.)

Next, make the directory src/hello inside your workspace, and in that directory create a file named hello.go that looks like:

package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}

Then build it with the go tool:

C:\> cd %USERPROFILE%\go\src\hello
C:\Users\Gopher\go\src\hello> go build
The command above will build an executable named hello.exe in the directory alongside your source code. Execute it to see the greeting:

C:\Users\Gopher\go\src\hello> hello
hello, world
If you see the "hello, world" message then your Go installation is working.



OR 
you can just write the above code in any editor and save it with extension .go
and then run it in cmd using the command " go run fileName.go "
Make sure that the path is set.



Thanks for Reading..
Contacts : 9794413845 divyansh.gupta_cs15@gla.ac.in

**Repository under construction wait till it completes**



