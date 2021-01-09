
import("jquery-3.1.1(1)");

function RpcPost() {
    var commit =  $("div#commit input").attr("value",undefined).val();
    var params =  $("div#params input").attr("value",undefined).val();
    if(commit==""){
        return
    }
    Write_Maessge("$User=>",commit+" " +params,"red","green");  //打印的对话框中
    post("http://127.0.0.1:8080/RPC",commit,params)
}

var running = 0;
function post(request_url,commit,params) {
    if(!running) {
        running=1;   //防止重复运行
        var request = new XMLHttpRequest();
        request.open('POST', request_url);
        request.setRequestHeader("Content-type", "application/json");
        var send_data = {'url': request_url, 'commit': commit, 'params': params};
        request.send(JSON.stringify(send_data));
        request.onload = function (e) {
            if (request.status === 200) {
                var result =  JSON.parse(request.response)
                if((result!=null)) {
                    if ((result["error"] != 0)) {
                        alert(result["ErrorCode"])
                    } else {
<<<<<<< HEAD

                        Write_Maessge("$Server=>", JSON.stringify(result["result"]), "blue", "yellow")
=======
                        Write_Maessge("$Server=>", JSON.stringify(result["result"]), "blue", "gray")
>>>>>>> remotes/origin/master
                    }
                }else{
                    alert("服务器 json 解析错误")
                }
            } else {
                alert('服务器请求失败，请重试！');
            }
        }
    }
    running = 0
}
function New_Bitcoin(requst_target){
        var send_data = {"requst_target": requst_target, "params": ""};
        rpc_quest("http://127.0.0.1:8080/Verify", send_data, function (a) {
            document.getElementById("Addr").innerHTML = a
            document.getElementById("status").innerText = "状态 : OK!"
        })
}

function Verify_Bitcoin(requst_target) {
    var addr = document.getElementById("BitcoinAddr").value
    if(addr.length<20){
        document.getElementById("BitcoinAddr").value = ""
        document.getElementById("res").innerHTML = "结果:长度应为20字节"
    }else {
        var send_data = {"requst_target": requst_target, "params": addr};
        rpc_quest("http://127.0.0.1:8080/Verify", send_data, function (a) {
            document.getElementById("res").innerHTML = "结果:" + a
        })
    }
}

function  rpc_quest(url,body,resu) {
    var request = new XMLHttpRequest();
    request.open('POST', url);
    request.setRequestHeader("Content-type", "application/json");
    request.send(JSON.stringify(body));
    request.onload = function (e) {
        if (request.status === 200) {
            var result = JSON.parse(request.response)
            if ((result != null)) {
                if ((result["error"] != 0)) {
                    alert(result["ErrorCode"])
                } else {
                    var res = JSON.stringify(result["result"])
                    resu(res);
                }
            } else {
                alert("服务器 json 解析错误")
            }
        } else {
            alert('服务器请求失败，请重试！')
        }
    }
}
