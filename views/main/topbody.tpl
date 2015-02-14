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
        }
        #main{
            width: 100%;
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
            &nbsp;&nbsp;
        </td>
        <td style="width: 100px;text-align: right">
        </td>
    </tr>
</table>
</body>
</html>