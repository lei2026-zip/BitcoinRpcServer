
import("static/js/jquery-3.1.1(1)")
//显示描述
function Describe(obj,evt){
    var e = evt||window.event
    var x = e.clientX
    var y = e.clientY
    var des =  document.getElementsByClassName(obj.id)[0]
    des.style.display = "block"
    if(des.innerText.length<=0){
        des.style.display = "none"
        return
    }
    des.style.left = x+30+"px"
    des.style.top = y+"px"
}
//关闭描述
function Undescribe(obj){
    var des =  document.getElementsByClassName(obj.id)[0]
    des.style.display = "none"
}

//