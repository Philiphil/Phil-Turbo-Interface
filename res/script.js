      var Jrep = function(arg,data)
      {
        this.arg = arg;
        this.data = data || ["null"];
      }
      Jrep.prototype.send = function(){
        var x = JSON.stringify(this);
            astilectron.send(x)
      }

      if (document.readyState === 'complete' || document.readyState !== 'loading') {
        onLoad();
      } else {
        document.addEventListener('DOMContentLoaded', onLoad);
      }
      function onLoad()
      {
        swift("index");
      }

      function swift(cible)
      {
        var s = document.querySelectorAll('section');
        for (var e = 0; e < s.length; e++) {
          s[e].classList.add("invisible")
        }
        var t = document.querySelector('#'+cible);
        var n = document.querySelector('#nav')

        t.classList.remove("invisible");
        if(cible === "index"){
          n.classList.add("invisible");
        }else{
          n.classList.remove("invisible");
        }
      }