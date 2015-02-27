<ul class="nav nav-tabs">
<li role="presentation" ><a href="javascript:history.go(-1)">分配叮当卡</a></li>
    <li role="presentation" class="active"><a href="#">家庭成员绑定卡号</a></li>
</ul>


<div class="row">
<div class="col-md-1"></div>
<div class="col-md-10">


<table class="table table-striped" style="margin-top: 30px">
<tr>
    <th>
        #
    </th>
    <th>
        成员姓名
    </th>
    <th>
        性别
    </th>
    <th>
        成员身份
    </th>
    <th>
        已绑定卡号
    </th>
    <th style="width:15%">
        操作
    </th>
</tr>

<!--显示家长-->
{{range .users}}
<tr>
    <td>#</td>
    <td>{{.Realname}}</td>
    <td>
        {{ if eq .Gender 0}}
            女
        {{ else }}
            男
        {{ end }}
    </td>
    <td>
        {{ if eq .Guid $.main_guid}}
            主家长
        {{ else }}
           家长
        {{ end }}
    </td>
    <td></td>
    <td>
        <input type="button"  value="绑定卡号" class="btn btn-primary"/>
    </td>
</tr>
{{end}}

<!--显示学生-->
{{range $key, $value := .members.stus}}
<tr>
    <td>#</td>
    <td>{{$value.Realname}}</td>
    <td>
        {{ if eq $value.Gender 0}}
            女
        {{ else }}
            男
        {{ end }}
    </td>
    <td>学生</td>
    <td></td>
    <td>
        <input type="button"  value="绑定卡号" class="btn btn-primary"/>
    </td>
</tr>
{{end}}


</table>

</div>
<div class="col-md-1"></div>
</div>