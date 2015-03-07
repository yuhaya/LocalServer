<ul class="nav nav-tabs">
    <li role="presentation" class="active"><a href="{{urlfor "FamilyController.Index"}}">家庭列表</a></li>
    <li role="presentation"><a href="{{urlfor "FamilyController.Add"}}">添加家庭</a></li>
</ul>


<table class="table table-striped" style="margin-top: 30px">
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
            <input type="button" onclick="location.href='{{urlfor "FamilyController.EditFamily"}}?guid={{$value.guid}}'" value="编辑" class="btn btn-primary"/>
            <input type="button" onclick="location.href='{{urlfor "FamilyController.ManagerMember"}}?guid={{$value.guid}}'" value="管理家庭成员" class="btn btn-primary"/>
            <input type="button"  onclick="location.href='{{urlfor "FamilyController.EditFamily"}}?guid={{$value.guid}}'"  value="家庭成员关系" class="btn btn-primary"/>
            <input type="button"  onclick="location.href='{{urlfor "FamilyController.MemeberList"}}?family_guid={{$value.guid}}'"  value="家庭成员卡号绑定" class="btn btn-primary"/>
            <input type="button"  onclick="location.href='{{urlfor "FamilyController.Notice"}}?guid={{$value.guid}}'"   value="家庭成员卡号通知绑定" class="btn btn-primary"/>
            <input type="button"  data-href="{{urlfor "FamilyController.Delete"}}?guid={{$value.guid}}"  value="删除" class="btn btn-danger DelFm"/>
        </td>
    </tr>
    {{end}}
{{ end }}
</table>

<div class="row">
<div class="col-md-4"></div>
<div class="col-md-4">
    {{if .paginator.HasPages}}
    <nav style="margin: 10px auto">
        <ul class="pagination">
            {{if .paginator.HasPrev}}
            <li><a href="{{.paginator.PageLinkFirst}}">{{ i18n .Lang "paginator.first_page"}}</a></li>
            <li><a href="{{.paginator.PageLinkPrev}}">&laquo;</a></li>
            {{else}}
            <li class="disabled"><a>{{ i18n .Lang "paginator.first_page"}}</a></li>
            <li class="disabled"><a>&laquo;</a></li>
            {{end}}
            {{range $index, $page := .paginator.Pages}}
            <li{{if $.paginator.IsActive .}} class="active"{{end}}>
            <a href="{{$.paginator.PageLink $page}}">{{$page}}</a>
        </li>
        {{end}}
        {{if .paginator.HasNext}}
        <li><a href="{{.paginator.PageLinkNext}}">&raquo;</a></li>
        <li><a href="{{.paginator.PageLinkLast}}">{{ i18n .Lang "paginator.last_page"}}</a></li>
        {{else}}
        <li class="disabled"><a>&raquo;</a></li>
        <li class="disabled"><a>{{ i18n .Lang "paginator.last_page"}}</a></li>
        {{end}}
    </ul>
</nav>
{{end}}
</div>
<div class="col-md-4"></div>
</div>
