<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
        "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
<head>
    <meta http-equiv="Content-Type" content="text/html;charset=UTF-8">
    <title>{{.AppName}}</title>
    <style>
        *{
            margin: 0px;
            padding: 0px;
        }
        body{
            background-color: #fff;
            font: 15px/1.4 Helvetica, arial, freesans, clean, sans-serif, "Segoe UI Emoji", "Segoe UI Symbol";
        }
        #list{
            color: #000;
        }
        #list td{
            cursor: pointer;
            padding: 5px 10px;
        }
        #list td:hover{
            color: blue;
        }

    </style>
</head>
<body>

<table id="list">
    <tr>
        <td data-url="{{urlfor "ExampleController.Index"}}">
            示例模块
        </td>
    </tr>
    <tr>
        <td data-url="{{urlfor "FamilyController.Index"}}">
            家庭管理
        </td>
    </tr>
    <tr>
        <td data-url="{{urlfor "CardController.Index"}}">
            卡片管理
        </td>
    </tr>
</table>
<script type="application/javascript" src="/static/js/lib/jquery-1.9.1.min.js"></script>
<script>
    $("#list td").click(function(){
        var url = $(this).attr("data-url");
        window.top.frames["right"].location.href = url;
    })
</script>
</body>
</html>