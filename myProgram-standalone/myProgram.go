package main 

import (
	"ethos/syscall"
	"ethos/ethos"
	ethosLog "ethos/log"
        "ethos/efmt"
)

func main () {
	me := syscall.GetUser()
	path := "/user/" + me + "/myDir/"
        status := ethosLog.RedirectToLog("myProgram")
	if status != syscall.StatusOk {
        efmt.Fprintf(syscall.Stderr,"Error opening %v: %v\n", path, status)
        syscall.Exit(syscall.StatusOk)
	}

	data    := MyType {10,20,30,40}
        fd, status := ethos.OpenDirectoryPath(path)
	data.Write(fd)
        data.WriteVar(path +"foobar")
        efmt .Println("output:",data.x1,data.y1,data.x2,data.y2)
        data.x1 = data.x1*10
	data.y1 = data.y1*10	
	data.x2 = data.x2*10
        data.y2 = data.y2*10
	efmt.Println("output:",data.x1,data.y1,data.x2,data.y2)
       
}
