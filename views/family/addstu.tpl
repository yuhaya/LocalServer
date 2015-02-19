<div class="page-header">
    <h1>添加家庭成员 <small>学生</small></h1>
</div>

<div class="row" style="margin-top: 30px">
<div class="col-md-3"></div>
<div class="col-md-6">
    <form action="{{urlfor "FamilyController.AddMember"}}" method="POST" name="form_add_memeber" id="form_add_memeber">

    <input type="hidden" value="{{.family_guid}}" name="family_guid" />
    <input type="hidden" value="{{.memeber_type}}" name="memeber_type" />

    <div class="form-group control-group">
        <label for="Realname">学生姓名</label>
        <input type="text" class="form-control validate" data-display="学生名称" data-rules="required|max_length[10]"  name="Realname" id="Realname" placeholder="Realname">
        <p class="help-block"></p>
    </div>
    <div class="form-group control-group">
        <label for="Sid">学号</label>
        <input type="text" class="form-control validate" data-display="学号" data-rules="required|numeric|max_length[30]"   name="Sid" id="Sid" placeholder="学号">
    </div>

    <div class="form-group control-group">
        <label for="Spell">姓名拼音</label>
        <input type="text" class="form-control validate" data-display="姓名拼音" data-rules="required|max_length[10]"  name="Spell" id="Spell" placeholder="Spell">
        <p class="help-block"></p>
    </div>


    <div class="form-group control-group">
        <label for="Grade_guid">年级</label>
        <select name="Grade_guid" class="form-control validate"  id="Grade_guid" data-display="年级" data-rules="required" >
            {{ range .grades}}
            <option value="{{.guid}}">{{.name}}</option>
            {{ end }}
        </select>
        <p class="help-block"></p>
    </div>

    <div class="form-group control-group">
        <label for="Class_guid">班级</label>
        <select name="Class_guid" class="form-control validate"  id="Class_guid" data-display="班级" data-rules="required" >
            <option value="">选择班级</option>
        </select>
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
        <label for="Birthday">生日</label>
        <input type="date" class="form-control validate" data-display="生日" data-rules="required"  name="Birthday" id="Birthday" placeholder="Birthday">
    </div>

    <div class="form-group">
        <label for="Enrol_time">入学时间</label>
        <input type="date" class="form-control validate" data-display="入学时间" data-rules="required"  name="Enrol_time" id="Enrol_time" placeholder="Enrol_time">
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
    grade_class : '{{.grade_class_json}}',
    type:"{{.memeber_type}}"
}
</script>