package main

import (
	"github.com/miketheprogrammer/go-thrust/thrust"
	"github.com/miketheprogrammer/go-thrust/lib/bindings/window"
	"github.com/miketheprogrammer/go-thrust/lib/commands"
	"fmt"
	"encoding/json"
)

var (
	kultcenter = "kultcenter"
	port = "8989"
)



func main() {
	thrust.InitLogger()
	thrust.Start()


	thrustWindow := thrust.NewWindow(thrust.WindowOptions{
		RootUrl:  "http://127.0.0.1:"+port+"/"+kultcenter+"/index",
		HasFrame: true,
	})
	thrustWindow.Show()
	thrustWindow.Maximize()
	thrustWindow.Focus()
	thrustWindow.OpenDevtools()


	thrustWindow.HandleRemote(func(er commands.EventResult, this *window.Window) {
		    res := Jrep{}
    		json.Unmarshal([]byte(er.Message.Payload), &res)


    		snd := Jrep{}
    		snd.Arg = res.Arg
    		switch res.Arg {
    		case "todo" :
    			snd.Data = append(snd.Data, db_get("todo"))
    			v, _ :=json.Marshal(snd)
    			this.SendRemoteMessage(string(v[:]))
    		case "todo_end" :
    			fmt.Println(res)
    			todos, _  := json.Marshal(res.Data)
    			db_set("todo", string(todos[:]))
    		}
	})
	run()    
}
