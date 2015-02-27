<div class="container top-20 ">
    <div class="row">
        <div class="col-md-6">
            <div class="form-group">
                <label>学校名称</label>
                <input id="name" type="text" class="form-control" value="{{.school.Name}}"/>
            </div>
            <div class="form-group">
                <label>省份</label>
                <input id="province" type="text" class="form-control" value="{{.school.Province}}"/>
            </div>
            <div class="form-group">
                <label>城市</label>
                <input id="city" type="text" class="form-control" value="{{.school.City}}"/>
            </div>
            <div class="form-group">
                <label>区县</label>
                <input id="county" type="text" class="form-control" value="{{.school.County}}"/>
            </div>
            <div class="form-group">
                <label>详细地址</label>
                <input id="address" type="text" class="form-control" value="{{.school.Address}}"/>
            </div>
            <div class="form-group">
                <label>大门数量</label>
                <input id="door" type="text" class="form-control" style="ime-mode:Disabled" value="{{.school.DoorNum}}"/>
            </div>
            <button id="save" class="btn btn-default">保存</button>
            <input id="guid" type="hidden" value="{{.school.Guid}}"/>
        </div>
    </div>
</div>