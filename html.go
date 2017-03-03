package main
import "net/http"
import "fmt"

func init(){
  http.HandleFunc("/"+kultcenter+"/index", h_index)
}

func h_index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, html())
}


func run(){
  http.ListenAndServe(":"+port, nil)
}


func html() string{
	return `
  <html>
  <head>
        <script src="https://code.jquery.com/jquery-3.1.1.min.js"></script>
    <script>

      
      THRUST.remote.listen(function (event) {
        console.log(event.payload);
      })

                 var Jrep = function(arg,data)
          {
            this.arg = arg;
            this.data = data || ["null"];
          }

          Jrep.prototype.send = function(){
            var x = JSON.stringify(this);
            THRUST.remote.send(x);
          }

        function todo()
        {
          var x = new Jrep("todo");
          x.send();
        }
    </script>
  </head>
  <style>

  ul,li{
    list-style-type: none;
    padding: 0px;
  }


  html,body,section{
    font-family: Helvetica, Arial, sans-serif;
    background-color : #27272c;    
    }
    .title_blue{
      color : #0DACE0;
      text-shadow: 0px  0px  4px rgba(13, 172, 224, 0.4);
    }

    #index{
      text-align: center;
    }


  </style>
  <body>
  <section id="index">
    <h1 class="title_blue" style="text-align:center">Kult Center</h1>

<ul>
  <li> <button onclick="todo()">Todo</button></li>
</ul>
      
  </section>

</body>
</html>
`
}
