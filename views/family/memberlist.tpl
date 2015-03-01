<ul class="nav nav-tabs">
    <li role="presentation" class="active"><a href="{{urlfor "FamilyController.Index"}}">家庭列表</a></li>
<li role="presentation"><a href="{{urlfor "FamilyController.Add"}}">添加家庭</a></li>
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
</tr>

<!--显示家长-->
{{range $key, $value := .members.users}}
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
    <td>
        {{ if eq $value.Guid $.main_guid}}
            主家长
        {{ else }}
           家长
        {{ end }}
    </td>
    <td>
        {{ if (map_exist $.cards $value.Guid) }}
            {{ range $k,$v := $.cards }}
                    {{ if eq $k $value.Guid}}

        <ul class="list-group">
            {{range $card := $v}}
            <li class="list-group-item">
                <span class="badge delcard" style="cursor: pointer" onclick="location.href='{{ urlfor "FamilyController.Del" }}?card={{$card}}&guid={{$value.Guid}}&family_guid={{$.family_guid}}'" >删除</span>
                {{$card}}
            </li>
            {{end}}
        </ul>


                    {{ end }}
            {{ end }}
        {{ else }}
            未绑定卡
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
    <td>
        {{ if (map_exist $.cards $value.Guid) }}
        {{ range $k,$v := $.cards }}
        {{ if eq $k $value.Guid}}
        <ul class="list-group">
            {{range $card := $v}}
            <li class="list-group-item">
                <span class="badge delcard" style="cursor: pointer" onclick="location.href='{{ urlfor "FamilyController.Del" }}?card={{$card}}&guid={{$value.Guid}}&family_guid={{$.family_guid}}'">删除</span>
                {{$card}}
            </li>
            {{end}}
        </ul>
        {{ end }}
        {{ end }}
        {{ else }}
        未绑定卡
        {{ end }}
    </td>
</tr>
{{end}}


</table>

</div>
<div class="col-md-1"></div>
</div>




