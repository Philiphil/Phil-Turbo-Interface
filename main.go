package main

import (
"github.com/asticode/go-astilectron"
"fmt"
"encoding/json"
)

var folder = "PhilTurboSoft"
var	port = "8789"
var ws *astilectron.Window
var index = "index"


func send_windows(msg string){
	fmt.Println(msg)
	ws.Send(string(msg))
}


func main() {
	go run()
	a, _ := astilectron.New(astilectron.Options{AppName:"PTS"})
	defer a.Close()
	a.HandleSignals()
	a.Start()


	ws, _ = a.NewWindow("http://127.0.0.1:"+port+"/"+folder+"/"+index, &astilectron.WindowOptions{
			Center: astilectron.PtrBool(true),
			Height: astilectron.PtrInt(600),
			Width:  astilectron.PtrInt(800), 
			})

	  ws.On(astilectron.EventNameWindowEventMessage, func(e astilectron.Event) (deleteListener bool) {
	    var m string
	    e.Message.Unmarshal(&m)
	    res := Jrep{}
	    json.Unmarshal([]byte(m), &res)

	    snd := Jrep{}
		snd.Arg = res.Arg
		switch res.Arg {
		case "todo" :
			snd.Data = append(snd.Data, db_get("todo"))
			v, _ :=json.Marshal(snd)
			send_windows(string(v[:]))
		case "todo_end" :
			fmt.Println(res)
			todos, _  := json.Marshal(res.Data)
			db_set("todo", string(todos[:]))
		}
	    return
	})

	ws.Create()
  	a.Wait()  
}