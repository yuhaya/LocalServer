<ul class="nav nav-tabs">
  <li role="presentation" class="active"><a href="#">分配叮当卡</a></li>
  <li role="presentation"><a href="#">管理叮当卡</a></li>
  <li role="presentation"><a href="#">创建家庭</a></li>
</ul>

<div class="container" id="search_box">
	<form class="form-horizontal" action="{{urlfor "CardController.Search"}}" method="post" name="search_form" id="search_form">
	  <div class="form-group">
	    <div class="col-sm-10">
	    	<div class="control-group">
	      		<input type="text" class="form-control validate" value="{{ .search_condition }}" name="search_condition" id="search_condition" data-rules="required"  placeholder="家庭成员姓名/成员手机号/家庭名称" data-display="搜索条件">
	      		<p class="help-block"></p>
			</div>
	    </div>
	    <input type="submit" id="search" class="col-sm-2 btn btn-primary" value="搜索">
	  </div>
	</form>

{{ if gt .search_by_familiy_num 0 }}
<h2>家庭名称检索结果</h2>
<table class="table table-striped">
<tr>
    <th>#</th>
    <th>家庭名称</th>
    <th>主家长名称</th>
    <th>主家长手机号</th>
    <th style="width: 20%">操作</th>
</tr>
{{range .search_by_familiy}}
<tr>
    <td>#</td>
    <td>
        {{.name}}
    </td>
    <td>
        {{.user_realname}}
    </td>
    <td>
        {{.user_phone}}
    </td>
    <td>

        <input type="button" class="btn btn-default" value="分配卡">
    </td>
</tr>
{{end}}
</table>
{{ end }}

        {{ if gt .search_by_user_num 0 }}
<h2>成员名称检索结果</h2>
<table class="table table-striped">
<tr>
    <th>#</th>
    <th>家庭名称</th>
    <th>成员名称</th>
    <th style="width: 20%">操作</th>
</tr>
{{range .search_by_user}}
<tr>
    <td>#</td>
    <td>
        {{.name}}
    </td>
    <td>
        {{.user_realname}}
    </td>
    <td>
        <input type="button" class="btn btn-default" value="分配卡">
    </td>
</tr>
{{end}}
</table>
        {{ end }}


        {{ if gt .search_by_phone_num 0 }}
<h2>家长手机号检索结果</h2>
<table class="table table-striped">
<tr>
    <th>#</th>
    <th>家庭名称</th>
    <th>成员名称</th>
    <th>家长手机号</th>
    <th style="width: 20%">操作</th>
</tr>
{{range .search_by_phone}}
<tr>
    <td>#</td>
    <td>
        {{.name}}
    </td>
    <td>
        {{.user_realname}}
    </td>
    <td>
        {{.user_phone}}
    </td>
    <td>
        <input type="button" class="btn btn-default" value="分配卡">

    </td>
</tr>
{{end}}
</table>
        {{ end }}



</div>