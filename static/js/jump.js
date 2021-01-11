/*
* 2020年12月23日10:27:05
* 打印触发对象的id到 命令框 input 里
* 快捷输入
* */
function Write_text(obj) {
    var commit =  obj.getAttribute("id")
    $("#commit input").val(commit)
    $("#params input").val("")
}




function Write_Maessge(master,str,ahead_color,font_color) {
    var contant = document.createElement("div")
    var ahead = document.createElement("span")
    var data = document.createElement("font")
    contant.setAttribute("id","massage")
    ahead.innerText = master
    ahead.style.color =ahead_color
    contant.appendChild(ahead)
    data.innerHTML = str
    data.style.color = font_color
    data.style.wordBreak = "break-all"
    data.style.display = "block"
    data.style.float = "right"
    data.style.width = "900px"
    contant.appendChild(data)
    contant.style.float = "left"
    $("#View").prepend(contant)
}

function Clear() {
   var view =  document.getElementById("View")
    var child = view.children
    for(i=0;i<child.length;i++){
        view.removeChild(child[i])
        i--
    }
   // var commit =   $("div#commit input").attr("value",undefined).val()
   //  if(commit==""){
   //      return
   //  }
   //  Write_Maessge("$User=>","help " +`"`+commit+`"`,"red","yellow");  //打印的对话框中
   //  post("http://127.0.0.1:8080/RPC","help",`"`+commit+`"`)
}