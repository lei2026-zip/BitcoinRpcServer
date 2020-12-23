
import("jquery-3.1.1(1)")
//list的显示描述
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
//list的关闭描述
function Undescribe(obj){
    var des =  document.getElementsByClassName(obj.id)[0]
    des.style.display = "none"
}

/*
* obj : 触发对象
* maxlen : 最大返回结果数量
* */
//COMMIT 关键字查询
function Query_List(obj,maxlen) {
   var str = obj.value
    $("#tip").empty()
   if(str.length>0){
        Query_String(str,$("#list ul li"),maxlen-1)
   }
}
/*
* ahead : 要搜索的字段
* obj : 要搜索的对象
* maxlen : 最大返回结果数量
* */
//遍历列表  2020年12月22日19:09:36
function Query_String(str,obj,maxlen) {
    var lenth = obj.length
    var strArr =  Compare_StrAhead(str,obj,maxlen)
    $("#tip").css({"display":"block"})
    for(var i=0;i<strArr.length;i++){
        var span = document.createElement("span")
        var str  = document.createTextNode(strArr[i])
        span.style.cursor = "pointer"
        span.style.display = "block"
        span.appendChild(str)
        $("#tip").append(span)
    }
}

function clearspan(){
    $("#tip").empty()
    $("#select input").empty()
}

/*
* ahead : 要搜索的字段
* obj : 要搜索的对象
* maxlen : 最大返回结果数量
* */
/*关键字查询算法 2020年12月22日18:06:55*/
function Compare_StrAhead(ahead,obj,maxlen) {
    var index = new Array()
    var i=0,pin=1,len = ahead.length;
    for (i=0;i<obj.length;i++){
        if(ahead[0]==obj[i].id[0]){
            index.push(i)
        }
    }
    if(len>1){
        var lenth  = index.length
        for (i=0;i<=lenth;i++){
            if(i==lenth){
                i=0
            }
            if(obj[index[i]].id.length>pin){
                if (ahead[pin] != obj[index[i]].id[pin]) {
                    index.splice(i, 1)   //切割
                    i--
                    lenth--
                }
            }else{
                index.splice(i, 1)   //切割
            }
            if(i==(lenth-1)){
                pin++
                if(pin>=len){
                    break
                }
            }
        }
    }
    var strarr = new Array()
    for(i=0;i<index.length;i++)
    {
        strarr.push(obj[index[i]].id)
        if(i>=(maxlen)){
            break
        }
    }
    return strarr
}


