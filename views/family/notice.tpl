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
                卡号
            </th>
            <th>
                绑定对象姓名
            </th>
            <th>
                绑定对象姓身份
            </th>
            <th>
                通知对象
            </th>
            <th style="width:15%">
                操作
            </th>
        </tr>


        {{range $key, $value := .cardlink}}
        <tr>
            <td>#</td>
            <td>{{$value.Card}}</td>
            <td>
                {{ $value.Real_name}}
            </td>
            <td>
               学生
            </td>
            <td>

                <ul class="list-group">
                    {{range $rec := $value.Recevie }}
                    <li class="list-group-item">
                        <span class="badge delcard" style="cursor: pointer" onclick="location.href='{{ urlfor "FamilyController.NoticeDeleteBind" }}?card={{$value.Card}}&guid={{$rec.Guid}}'" >删除</span>
                    {{$rec.Name}}
                    ---
                    {{ if eq $rec.Identify 0}}
                        学生
                    {{else}}
                        家长
                    {{end}}
                    {{end}}
                    </li>
                </ul>

        </td>
        <td>
            <button type="button"  data-val="{{$value.Card}}" data-toggle="modal" data-type="1"  data-target=".input_card"  class="btn btn-primary bcard">添加通知对象</button>
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
            <form class="form-inline" name="card_form" id="card_form" method="post" action='{{ urlfor "FamilyController.NoticeBind" }}'>
                <input type="hidden" name="card" id="card" value=""/>
                <input type="hidden" name="family_guid" value="{{.family_guid}}"/>
                <div class="form-group">
                    <label for="notice">选择通知对象</label>
                    <select name="notice" id="notice" class="form-control" >
                        {{range $v := .fms.users}}
                        <option value="{{$v.Guid}}">{{$v.Realname}}</option>
                        {{end}}
                        {{range $v := .fms.stus}}
                        <option value="{{$v.Guid}}">{{$v.Realname}}</option>
                        {{end}}
                    </select>
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
