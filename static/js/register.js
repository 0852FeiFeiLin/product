//检验输入是否合法
$(document).ready(function () {
    //1、获取表单元素   --> 通过name操作表单里面的具体元素
    $("#form").validate({ //这个就是验证合法性
        //2、指定要验证的内容
        rules: {
            //验证用户名和密码
            username: {
                required: true, //非空验证
                rangelength: [4, 12] //长度范围验证
            },
            password: {
                required: true,
                rangelength: [6, 20]
            },
            repassword: {
                required: true,
                rangelength: [6, 20],
                //判断两次密码是否一致，通过id
                equalTo: "#password"
            },
            email:{
                required: true,
                rangelength: [8, 20],
            }
        },
        //3、指定验证未通过时显示的错误信息
        messages: {  //注意：这个messages不加就是显示的英文
            username: {
                required: "用户名不能为空",
                rangelength: "用户名长度必须在4~12之间"
            },
            password: {
                required: "密码不能为空",
                rangelength: "密码长度必须在6~20之间"
            },
            repassword: {
                required: "确认密码不能为空",
                rangelength: "确认密码长度必须在6~20之间",
                equalTo: "输入的密码两次不一致"
            },
            email:{
                required: "邮箱不能为空",
                rangelength: "邮箱长度必须在8~20之间",
            }
        },
    })
});
function checkName() {
    /*获取 = Username 的值*/
    var username = document.getElementById("userName").value;
    /*定义一个json对象，因为怕存在跨域问题*/
    var userName = {
        "username":username,
    };
    $.ajax({
        url:"/isRegister",//请求的地址
        type:"POST",
        /*dataType:数据格式 有xml,json,html...*/
        dataType:"json",
        /*把数据序列化发送*/
        data:JSON.stringify(userName),
        /*内容类型*/
        contentType:"application/json;charset=utf8",
        success:function (msg) {
            /*回传的参数*/
            console.log(msg);
            if(msg === "true"){//已经存在
                $("#msg").css("color","red");
                $("#msg").html("用户名已经存在")
                $("#submit").attr("disabled",true)
                $("#submit").css("color","#ccc")

            }else if(msg === "false"){//不存在
                $("#msg").css("color","green");
                $("#msg").html(" ")
            }
        },
    })
}