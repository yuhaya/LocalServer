<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
<div class="container main">
    <div class="row">
        <div class="col-md-12">
            <div class="panel panel-default">
                <div class="panel-heading">
                    <h4>幼儿园总览</h4>
                </div>
                <table id="tb_grade" class="table">
                    <thead>
                    <tr>
                        <th>年级</th>
                        <th>班级</th>
                        <th>班主任</th>
                        <th>协管员</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range $k,$v:=.grade_list}}
                    <tr id="{{$v.Guid}}">
                        <td rowspan="0">{{$v.Name}}</td>
                    </tr>
                    {{end}}
                    </tbody>
                </table>
            </div>
        </div>

    </div>
</div>
</body>
</html>