package main

import (
	"github.com/miketheprogrammer/go-thrust/thrust"
	"github.com/miketheprogrammer/go-thrust/lib/bindings/window"
	"github.com/miketheprogrammer/go-thrust/lib/commands"
	//"fmt"
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
		    var res map[string]interface{}
    		json.Unmarshal([]byte(er.Message.Payload), &res)
    		arg := res["arg"].(string)
    		strs := res["data"].([]interface{})
    		var str []string
    		for _,e := range strs{
    			str = append(str,e.(string))
    			
    		}
    		rep := Jrep{arg,str}
    		switch rep.arg {
    		case "todo" :
    			tmp := db_get("todo")
    			this.SendRemoteMessage(tmp)
    		}
	})
	run()    
}
