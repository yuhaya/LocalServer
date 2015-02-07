<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
        "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
<head>
    <meta http-equiv="Content-Type" content="text/html;charset=UTF-8">
    <title>Error</title>
    <style type="text/css">
        body{
            background-color: #dddddd;
        }
        h2{
            font-size: 40px;
            text-align: center;
            margin-top: 50px;
            text-transform: uppercase;
            text-shadow: #fff 0 2px 0;
        }
        a{
            text-align: center;
            display: block;
            margin: 20px auto;
        }
    </style>
</head>
<body>
<h2>{{.content}}</h2>
{{ range $key, $value := .urlmsg }}
<a href="{{ $value }}">{{ $key }}</a>
{{ end }}

</body>
</html>



