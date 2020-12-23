/*
* 2020年12月23日10:27:05
* 打印触发对象的id到 命令框 input 里
* 快捷输入
* */
function Write_text(obj) {
    var commit =  obj.getAttribute("id")
    $("div#commit input").attr("value",commit)
}

function Write_Maessge(str) {
    var span = document.createElement("font")
    var str  = document.createTextNode(strArr[i])
    span.style.cursor = "pointer"
    span.style.display = "block"
    span.appendChild(str)
    $("#tip").append(span)
}