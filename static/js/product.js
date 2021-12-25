//利用Ajax提交搜索请求
function searchPro() {
    var proInfo = $("#shuru").val()
    console.log(proInfo)
    //ajax发起请求
    $.ajax({
        url:"/searchPro",
        type:"POST",
        dataType:"json",
        data:JSON.stringify(proInfo),
        contentType:"application/json;charset=utf-8",
        success:function(msg){
            //msg就是返回的结果数
            console.log(msg)
            if( msg != null){
                //如果查询到了产品，把id传递回来，然后通过此id再次发起请求，执行显示产品详情的方法
                window.location.href = "/productInfo?id="+msg
            }
        }

    })
}