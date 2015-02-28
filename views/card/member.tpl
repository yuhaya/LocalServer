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
                        {{range $card := $v}}
                            {{$card}}<br/>
                        {{end}}
                    {{ end }}
            {{ end }}
        {{ else }}
            未绑定卡
        {{ end }}
    </td>
    <td>
        <button type="button" data-val="{{$value.Guid}}" data-toggle="modal" data-type="1"  data-target=".input_card"  class="btn btn-primary bcard">绑定卡号</button>
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
        {{range $card := $v}}
        {{$card}}<br/>
        {{end}}
        {{ end }}
        {{ end }}
        {{ else }}
        未绑定卡
        {{ end }}
    </td>
    <td>
        <button type="button" data-val="{{$value.Guid}}" data-type="0" data-toggle="modal" data-target=".input_card"  class="btn btn-primary bcard">绑定卡号</button>
    </td>
</tr>
{{end}}


</table>

</div>
<div class="col-md-1"></div>
</div>

<div class="modal fade input_card" tabindex="-1" role="dialog" aria-labelledby="myLargeModalLabel" aria-hidden="true">

<div class="modal-dialog">
    <div class="modal-content">
        <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
            <h4 class="modal-title" id="myModalLabel">操作项</h4>
        </div>
        <div class="modal-body">
            <form class="form-inline" name="card_form" id="card_form" method="post" action='{{ urlfor "CardController.Add" }}'>
                <input type="hidden" name="guid" id="memguid" value=""/>
                <input type="hidden" name="type" id="memtype" value=""/>
                <input type="hidden" name="family_guid" value="{{.family_guid}}"/>
                <div class="form-group">
                    <label for="card">输入卡号</label>
                    <input type="email" class="form-control" id="card" name="card" placeholder="卡号">
                </div>
                <div id="error" style="display: inline;color: red"></div>
            </form>
        </div>
        <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
            <button type="button" id="card_form_submit" class="btn btn-primary">确定</button>
        </div>
    </div>
</div>


</div>


