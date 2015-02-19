<div class="page-header">
  <h1>添加家庭成员 <small>家长</small></h1>
</div>

<div class="row" style="margin-top: 30px">
<div class="col-md-3"></div>
<div class="col-md-6">
<form action="{{urlfor "FamilyController.AddMember"}}" method="POST" name="form_add_memeber" id="form_add_memeber">

<input type="hidden" value="{{.family_guid}}" name="family_guid" />
<input type="hidden" value="{{.memeber_type}}" name="memeber_type" />

<div class="form-group control-group">
    <label for="Realname">家长姓名</label>
    <input type="text" class="form-control validate" data-display="主家长名称" data-rules="required|max_length[10]"  name="Realname" id="Realname" placeholder="Realname">
    <p class="help-block"></p>
</div>
<div class="form-group control-group">
    <label for="Phone">家长手机号</label>
    <input type="text" class="form-control validate" data-display="家长手机号" data-rules="required|numeric|max_length[13]"   name="Phone" id="Phone" placeholder="Phone">
</div>
<div class="form-group control-group">
    <label for="Spell">家长姓名拼音</label>
    <input type="text" class="form-control validate" data-display="姓名拼音" data-rules="required|max_length[10]"  name="Spell" id="Spell" placeholder="Spell">
    <p class="help-block"></p>
</div>

<div class="radio">
    <label>
        <input type="radio" name="Gender" id="sex" value="0" checked>
        女
    </label>
</div>
<div class="radio">
    <label>
        <input type="radio" name="Gender" id="sex2" value="1">
        男
    </label>
</div>

<div class="form-group">
    <label for="IdCard">身份证号码</label>
    <input type="text" class="form-control" name="IdCard" id="IdCard" placeholder="Id Card">
</div>

<div class="form-group">
    <label for="Password">App登陆初始密码(123456)</label>
    <input type="Password" class="form-control" value="123456" disabled name="Password" id="Password" placeholder="Start Password">
</div>


<button type="submit" class="btn btn-primary">确定</button>
<button type="button" class="btn btn-default">重置</button>
</form>
</div>
<div class="col-md-3"></div>
</div>
<script>
var data = {
    type:"{{.memeber_type}}"
}
</script>