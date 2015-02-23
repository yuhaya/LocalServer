<ul class="nav nav-tabs">
    <li role="presentation" ><a href="{{urlfor "FamilyController.Index"}}">家庭列表</a></li>
    <li role="presentation" ><a href="{{urlfor "FamilyController.Add"}}">添加家庭</a></li>
    <li role="presentation" class="active"><a >管理家庭成员</a></li>
</ul>


<div class="row" style="margin-top: 30px">
<div class="col-md-1">
</div>
<div class="col-md-10">
    <button type="button" class="btn btn-primary btn-lg" onclick="location.href='{{urlfor "FamilyController.AddMemberShow"}}?memeber_type=user&family_guid={{.family_guid}}'" >添加家长</button>
    <button type="button" class="btn btn-default btn-lg" onclick="location.href='{{urlfor "FamilyController.AddMemberShow"}}?memeber_type=stu&family_guid={{.family_guid}}'" >添加学生</button>
</div>
<div class="col-md-1">
</div>
</div>

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
    <th style="width: 200px">
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
    <td style="width: 300px">
        <input type="button" onclick="location.href='{{urlfor "FamilyController.EditUserShow"}}?guid={{.Guid}}&family_guid={{$.family_guid}}'" value="编辑" class="btn btn-primary"/>
        <input type="button" onclick="location.href='{{urlfor "FamilyController.ShowUser"}}?guid={{.Guid}}&family_guid={{$.family_guid}}'" value="查看" class="btn btn-primary"/>
        {{ if ne .Guid $.main_guid}}
            <input type="button"  data-url="{{urlfor "FamilyController.SetMainUser"}}?guid={{.Guid}}&family_guid={{$.family_guid}}'" value="设为主家长" class="main_parent btn btn-primary"/>
        {{ end }}
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
    <td style="width: 700px">
        <input type="button" onclick="location.href='{{urlfor "FamilyController.EditStuShow"}}?guid={{$value.Guid}}&family_guid={{$.family_guid}}'" value="编辑" class="btn btn-primary"/>
        <input type="button" onclick="location.href='{{urlfor "FamilyController.ShowStu"}}?guid={{$value.Guid}}&family_guid={{$.family_guid}}'" value="查看" class="btn btn-primary"/>
    </td>
</tr>
{{end}}


</table>

</div>
<div class="col-md-1"></div>
</div>