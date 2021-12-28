function input_disabled() {
    /*    $("#goto_cart").hide();*/
}

$(function () {
    //点击加入购物车，先判断是否登录，没登录跳转登录，登录了显示小弹窗加入成功
    $("#jiaru").unbind("click").click(function (dangqian) {
        //unbind解释：就是接触此按钮的绑定事件，防止访问两次js代码，
        var isLogin = $("#isLogin").val()
        console.log("是否登录", isLogin)
        if (isLogin === "false") {
            alert("请先登录")
            /*跳转到登录页面*/
            window.location.href = "/login.html"
            return
        }
        $("#pop").show()
        $("#goto_cart").show();

    })
    //点击x，小弹窗关闭
    $("#header-close").click(function () {
        $("#goto_cart").hide();
        $("#pop").hide()
    })
})


/*

/!*这个是设置详情图片的方法*!/
$(function () {
    //获取到当前id
    var id = document.getElementById("proId").value
    /!*转为数字类型*!/
    var idd = Number(id)
    console.log(idd)
    //遍历li，填充背景
    $(".xiaoduimg li").each(function (index, ele) {
        console.log(index)
        /!*
                    $(this).css("background-image","url(手风琴img/img"+(index+1)+".jpg)")
        *!/
        //通过css设置背景图
        $(this).css("background-image", "url(./static/img/xiaoduimg/pro" + id + "/bg" + (index + 1) + ".png)")
    })
})*/
