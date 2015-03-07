<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
        "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
<head>
    <meta http-equiv="Content-Type" content="text/html;charset=UTF-8">
    <title>{{.AppName}}</title>
    <style type="text/css">
        *{
            padding: 0px;
            margin: 0px;
        }
        body{
            background-color: #f5f5f5;
            font: 13px/1.4 Helvetica, arial, freesans, clean, sans-serif, "Segoe UI Emoji", "Segoe UI Symbol";
            color: #333;
            border-bottom: 2px solid #ccc;
        }
        #main{
            width: 100%;
            float: left;
        }
        #login_out{
            width: 80px;
            height: 25px;
            line-height: 25px;
            border: 0px;
            color: #333;
            background-color: #000000;
            cursor: pointer;
        }
        #login{
            float: right;
            background-color: #2892BC;
            color:#f5f5f5;
            text-align: center;
            vertical-align: middle;
            padding: 30px 20px 10px 20px;
            cursor: pointer;
        }
    </style>
</head>
<body>
<table id="main">
    <tr>
        <td style="width:180px;text-align: center">
            <h2>小叮当管理系统</h2>
            <p>ID : <span>{{.Name}}</span></p>
        </td>
        <td>
            当前模式:
            <select name="" id="mode">
                <option value="0">
                    安检模式
                </option>
                <option value="1">
                    注册模式
                </option>
            </select>

        </td>
        <td style="width: 100px;text-align: right">
            <div id="login">
                登陆
            </div>
        </td>
    </tr>
</table>

<script type="application/javascript" src="/static/js/lib/jquery-1.9.1.min.js"></script>
<script>
    $("#login").click(function () {
        var url = $(this).attr("data-url");
        window.top.frames["right"].location.href = "/user/index";
    })

    var mode_get_url = '{{urlfor "MainController.Mode"}}';
    var mode_set_url = '{{urlfor "MainController.SetMode"}}';

    var m = setInterval(GetMode,3000);

    function GetMode(){
        $.ajax({
            url:mode_get_url,
            type:"post",
            dataType:"json",
            success:function(data){
                if(data.Code){
                    $("#mode").val(data.Data)
                }else{
                    window.top.alert("系统异常! 联系管理员!")
                }
            }
        })
    }

    $("#mode").change(function(){
        clearInterval(m)
        var mode = $(this).val();
        var data = {
            mode:mode
        }
        $.ajax({
            url:mode_set_url,
            data:data,
            type:"post",
            dataType:"json",
            success:function(data){
                if(!data.Code){
                    window.top.alert("设置失败！请联系管理员！")
                }else{
//                    alert("设置成功")
                }
                m = setInterval(GetMode,1000);
            }
        });
    });
</script>
</body>
</html>