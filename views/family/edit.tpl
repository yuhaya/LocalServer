<ul class="nav nav-tabs">
    <li role="presentation" ><a href="{{urlfor "FamilyController.Index"}}">家庭列表</a></li>
    <li role="presentation" ><a href="{{urlfor "FamilyController.Add"}}">添加家庭</a></li>
<li role="presentation" class="active"><a >编辑家庭</a></li>
</ul>

<div class="row" style="margin-top: 30px">
<div class="col-md-3"></div>
<div class="col-md-6">
    <form action="{{urlfor "FamilyController.EditSubmit"}}" method="POST" name="form_edit_family" id="form_edit_family">
    <input type="hidden" value="{{.fm.Guid}}" name="guid"/>
    <div class="form-group control-group">
        <label for="family_name">家庭名称</label>
        <input type="text" class="form-control validate" value="{{.fm.Name}}" data-display="家庭名称" data-rules="required|max_length[10]" data-rules="family_name" name="family_name" id="family_name" placeholder="Family Name">
        <p class="help-block"></p>
    </div>

    <button type="submit" class="btn btn-primary">确定</button>
    <button type="button" class="btn btn-default">重置</button>
</form>
</div>
<div class="col-md-3"></div>
        </div>