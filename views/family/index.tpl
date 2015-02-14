<ul class="nav nav-tabs">
    <li role="presentation" class="active"><a href="{{urlfor "FamilyController.Index"}}">家庭列表</a></li>
    <li role="presentation"><a href="{{urlfor "FamilyController.Add"}}">添加家庭</a></li>
</ul>


<table class="table" style="margin-top: 30px">
<tr>
    <th>#</th>
    <th>家庭名称</th>
    <th>主家长名称</th>
    <th>主家长手机号</th>
    <th>操作</th>
</tr>
{{ if eq .num 0 }}
<tr>
    <td colspan="5" align="center">
        无家庭数据,点击顶部添加家庭
    </td>
</tr>
{{ else }}
    {{range $key, $value := .list}}
    <tr>
        <td>{{$key}}</td>
        <td>{{$value.name}}</td>
        <td>{{$value.user_realname}}</td>
        <td>{{$value.user_phone}}</td>
        <td style="width: 700px">
            <input type="button" value="编辑" class="btn btn-primary"/>
            <input type="button" value="管理家庭成员" class="btn btn-primary"/>
            <input type="button" value="家庭成员关系" class="btn btn-primary"/>
            <input type="button" value="家庭成员卡号绑定" class="btn btn-primary"/>
            <input type="button" value="家庭成员卡号通知绑定" class="btn btn-primary"/>
            <input type="button" value="删除" class="btn btn-danger"/>
        </td>
    </tr>
    {{end}}
{{ end }}
</table>