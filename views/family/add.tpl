<ul class="nav nav-tabs">
    <li role="presentation" ><a href="{{urlfor "FamilyController.Index"}}">家庭列表</a></li>
    <li role="presentation" class="active"><a href="{{urlfor "FamilyController.Add"}}">添加家庭</a></li>
</ul>

<div class="row" style="margin-top: 30px">
<div class="col-md-3"></div>
<div class="col-md-6">
<form action="{{urlfor "FamilyController.AddSubmit"}}" method="POST" name="form_add_family" id="form_add_family">
<div class="form-group control-group">
    <label for="family_name">家庭名称</label>
    <input type="text" class="form-control validate" data-display="家庭名称" data-rules="required|max_length[10]" data-rules="family_name" name="family_name" id="family_name" placeholder="Family Name">
    <p class="help-block"></p>
</div>
<div class="form-group control-group">
    <label for="name">主家长名称</label>
    <input type="text" class="form-control validate" data-display="主家长名称" data-rules="required|max_length[10]"  name="name" id="name" placeholder="Name">
    <p class="help-block"></p>
</div>
<div class="form-group control-group">
    <label for="phone">主家长手机号</label>
    <input type="text" class="form-control validate" data-display="主家长手机号" data-rules="required|numeric|max_length[13]"   name="phone" id="phone" placeholder="Phone">
</div>

<div class="radio">
    <label>
        <input type="radio" name="gender" id="sex" value="0" checked>
        女
    </label>
</div>
<div class="radio">
    <label>
        <input type="radio" name="gender" id="sex2" value="1">
        男
    </label>
</div>

<div class="form-group">
    <label for="id_card">身份证号码</label>
    <input type="text" class="form-control" name="id_card" id="id_card" placeholder="Id Card">
</div>

<div class="form-group">
    <label for="password">App登陆初始密码(123456)</label>
    <input type="password" class="form-control" value="123456" disabled name="password" id="password" placeholder="Start Password">
</div>


<button type="submit" class="btn btn-primary">确定</button>
<button type="button" class="btn btn-default">重置</button>
</form>
</div>
<div class="col-md-3"></div>
</div>