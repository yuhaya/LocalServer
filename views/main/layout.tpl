<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
        "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{.AppName}}</title>
    <!-- Bootstrap -->
    <link href="/static/css/lib/bootstrap.min.css" rel="stylesheet">

    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
    <script src="/static/js/lib/html5shiv.min.js"></script>
    <script src="/static/js/lib/respond.min.js"></script>
    <![endif]-->
    <link rel="stylesheet" href="/static/css/lib/common.css"/>
    <link rel="stylesheet" href="/static/css/{{.ControllerName}}/{{.MethodName}}.css"/>
</head>
<body>
<div id="title_name">
{{ if ne .TitleName ""}}
    >> {{.TitleName}}
{{ end }}
</div>
<div id="content">{{.LayoutContent}}</div>

<script type="application/javascript" src="/static/js/lib/jquery-1.9.1.min.js"></script>
<script type="application/javascript" src="/static/js/lib/bootstrap.min.js"></script>
<script type="application/javascript" src="/static/js/lib/sea.js"></script>
<script type="application/javascript">
    seajs.config({

        // 别名配置
        alias: {
            'es5-safe': '/static/js/lib/es5-safe'
        },

        // 路径配置
        paths: {
            'js_lib_root': '/static/js/lib'
        },

        // 变量配置
        vars: {
            'controller': '{{.ControllerName}}',
            'method' : '{{.MethodName}}',
            'module_path' : '/static/js/{{.ControllerName}}/{{.MethodName}}'
        },

        // 映射配置
        map: [
//            ['http://example.com/js/app/', 'http://localhost/js/app/']
        ],

        // 预加载项
        preload: [
            Function.prototype.bind ? '' : 'es5-safe',
        ],

        // 调试模式
        debug: true,

        // Sea.js 的基础路径
        base: '/static/js/',

        // 文件编码
        charset: 'utf-8'
    });
    seajs.use(['lib/commen', '{module_path}'], function(commen, main) {
        commen.init();
        main.init();
    });
</script>
</body>
</html>