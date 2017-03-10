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
  return `  <html>
  <head>
  <script>
      var Jrep = function(arg,data)
      {
        this.arg = arg;
        this.data = data || ["null"];
      }

      Jrep.prototype.send = function(){
        var x = JSON.stringify(this);
        THRUST.remote.send(x);
      }

      if (document.readyState === 'complete' || document.readyState !== 'loading') {
        onLoad();
      } else {
        document.addEventListener('DOMContentLoaded', onLoad);
      }

      function onLoad()
      {
        set_visible("index");
      }

      function set_visible(cible)
      {

        var s = document.querySelectorAll('section');
        for (var e = 0; e < s.length; e++) {
          s[e].classList.add("invisible")
        }
        var t = document.querySelector('#'+cible);
        t.classList.remove("invisible");
      }

      THRUST.remote.listen(function (event) {
        var x = JSON.parse(event.payload);
        switch (x.arg){
          case "todo":
          todo_rep(x.data);
          break;
        }
      })

      function swift(e)
      {
        set_visible(e);
      }



      function todo()
      {
        var x = new Jrep("todo");
        x.send();
      }

      function todo_rep(data)
      {
        data = JSON.parse(data)
        if(data[0] != "")
        {
           var input = document.querySelector("#todo_inp");
          for (var i = 0; i < data.length; i++) {
            input.insertAdjacentHTML('beforebegin','<li><input type="checkbox" onclick="todo_done(this)"><span>'+data[i]+'</span></li>');


          }
        }
      }

      function todo_done(el)
      {
        var v = el.parentElement;
        v.classList.toggle("crossed");
      }

      function todo_add(ev,el)
      {
        if(ev.keyCode !=13) return;
        var txt = el.value;
        var ul = el.parentNode ;
        ul.insertAdjacentHTML('beforeend','<li><input type="checkbox" onclick="todo_done(this)"><span>'+txt+'</span></li>');
        ul.removeChild(el);
        ul.appendChild(el)
        el.value = "";
        todo_end();
      }

      function todo_end()
      {
        var e = document.querySelectorAll("#todo ul li span");
        var z = [];
        for (var i = 0; i < e.length; i++) {
          var tmp = e[i].parentNode;
          if(!tmp.classList.contains("crossed"))
          {
            z[i] = e[i].innerHTML;
          }
          
        }
        var r = new Jrep("todo_end", z);
        r.send();
      }

    </script>
  </head>
  <style>

    /* RESET */

    html,body{
      padding: 0% ;
      margin: 0% ;
      width: 100% ;
      height: 100% ;
    }
    ul,li{
      list-style-type: none;
      padding: 0px;
    }


    html,body,section{
      font-family: Helvetica, Arial, sans-serif;
      background-color : #27272c;    
      color : white;
    }

    nav{
      background-color: #19191d;
      height: 100%;
      max-width: 300px;
float: left;
position: absolute;
    }
    h1{
      margin : 0px;
      padding: 0px;
    }
    section{
  //   display: inline-block;
    }
    /*STYLE*/
    .title_blue{
      color : #0DACE0;
      text-shadow: 0px  0px  4px rgba(13, 172, 224, 0.4);
    }

    .invisible{
      display: none;
    }

    #index{
      text-align: center;
    }

    /*TODO*/
    #todo{
      text-align: center;
    }
    #todo ul {
      display: inline-block;
    }
    #todo ul li{
      text-align: left;
    }

    #todo ul li.crossed{
      text-decoration: line-through;
    }

  </style>
  <body>

<nav>
lorem
</nav>

    <section id="index">
      <h1 class="title_blue" style="text-align:center">Kult Center</h1>

      <ul>
        <li> <button onclick="swift('todo');todo()">Todo</button></li>
      </ul>
      
    </section>

    <section id="todo">
      <ul>
        <li id="todo_inp"> <input type="text" placeholder="todo" name="todo_inp" onkeydown="todo_add(event,this)"> </li>
      </ul>
    </section>

  </body>
  </html>`
}
