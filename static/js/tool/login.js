layui.use(['form','layer','jquery'],function(){
    var form = layui.form,
        layer = parent.layer === undefined ? layui.layer : top.layer
        $ = layui.jquery;

    $(".loginBody .seraph").click(function(){
        layer.msg("这只是做个样式，至于功能，你见过哪个后台能这样登录的？还是老老实实的找管理员去注册吧",{
            time:5000
        });
    })

    //登录按钮
    form.on("submit(login)",function(data){
        $(this).text("登录中...").attr("disabled","disabled").addClass("layui-disabled");
	var	username1 = $("#userName").val(); 
   var password1 = $("#password").val();
	 
	
        setTimeout(function(){
		$.post("/Login/Login", {'UserName':username1, 'UserPwd':password1}, function (data) {
    //alert(JSON.stringify(data));
	//layer.msg(data);
	if(data == "0")
{
	layer.msg("不正确的账号和密码,请检查");
	
	$("#loginbutt").text("登录");
	$("#loginbutt").removeAttr("disabled");
	$("#loginbutt").removeClass("layui-disabled");
}else
{
	layer.msg("登录成功");
	
	window.location.href = "/user";
}
	
});

//layer.msg(ddd);

           // window.location.href = "/layuicms2.0";
        },500);
        return false;
    })

    //表单输入效果
    $(".loginBody .input-item").click(function(e){
        e.stopPropagation();
        $(this).addClass("layui-input-focus").find(".layui-input").focus();
    })
    $(".loginBody .layui-form-item .layui-input").focus(function(){
        $(this).parent().addClass("layui-input-focus");
    })
    $(".loginBody .layui-form-item .layui-input").blur(function(){
        $(this).parent().removeClass("layui-input-focus");
        if($(this).val() != ''){
            $(this).parent().addClass("layui-input-active");
        }else{
            $(this).parent().removeClass("layui-input-active");
        }
    })
})
