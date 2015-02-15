<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
<div class="container main">
    <div class="row">
        <div class="col-md-6">
            <div class="panel panel-default">
                <div class="panel-heading">
                    <ul class="panel-head-title">
                        <li><h4>年级</h4></li>
                        <li>
                            <button id="btn_add_grade" class="btn btn-default"><span
                                        class="glyphicon glyphicon-plus"></span>
                                年级管理
                            </button>
                        </li>
                    </ul>
                </div>
                <ul id="display-grade" class="list-group">
                    {{range $k,$v:=.list}}
                    <a href="#" rating="{{$v.Rating}}" pid="{{$v.Id}}" guid="{{$v.Guid}}" sid="{{$v.SchoolGuid}}"
                       class="list-group-item">{{$v.Name}}</a>
                    {{end}}
                </ul>
            </div>
        </div>
        <div class="col-md-6">
            <div class="panel panel-default">
                <div class="panel-heading">
                    <ul class="panel-head-title">
                        <li><h4>班级</h4></li>
                        <li>
                            <button id="btn_add_class" class="btn btn-default pull-right disabled"><span
                                        class="glyphicon glyphicon-plus"></span>
                                添加班级
                            </button>
                        </li>
                    </ul>
                </div>
                <table id="tb_class" class="table class-table">
                    <thead>
                        <tr>
                            <td>班级名称</td>
                            <td>班主任</td>
                            <td>协管员</td>
                            <td>操作</td>
                        </tr>
                    </thead>
                    <tbody>

                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>

<!-- 添加年级 -->
<div class="modal fade" id="addGrade" tabindex="-1" role="dialog" aria-hidden="true" data-backdrop="static">
    <div class="modal-dialog">
        <div class="modal-content ">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span
                            aria-hidden="true">&times;</span></button>
                <h4 class="modal-title">年级管理</h4>
            </div>
            <div class="modal-body">
                <div id="wrong_grade"></div>
                <div class="input-group-name">
                    <h5>年级名称</h5>

                    <div class="input-group">
                        <input id="txt_grade_name" type="text" class="form-control" placeholder=""/>
                      <span class="input-group-btn">
                        <button id="btn_append_grade" class="btn btn-default" type="button">添加</button>
                      </span>
                    </div>
                </div>
                <ul class="panel-head-title">
                    <li><h6>从低年级到高年级排序列表</h6></li>
                    <li>
                        <div class="btn-group" role="group" aria-label="...">
                            <button id="btn_up" type="button" class="btn btn-default" title="上移"><span
                                        class="glyphicon glyphicon-arrow-up"></span></button>
                            <button id="btn_down" type="button" class="btn btn-default" title="下移"><span
                                        class="glyphicon glyphicon-arrow-down"></span></button>
                            <button id="btn_del" type="button" class="btn btn-default" title="删除"><span
                                        class="glyphicon glyphicon-remove"></span></button>
                        </div>
                    </li>
                </ul>

                <ul id="list_grade" class="list-group">
                    {{range $k,$v:=.list}}
                    <a href="#" data="{{$v.Rating}}" pid="{{$v.Id}}" guid="{{$v.Guid}}" sid="{{$v.SchoolGuid}}"
                       class="list-group-item">{{$v.Name}}</a>
                    {{end}}
                </ul>
            </div>
            <div class="modal-footer">
                <button id="btn_save_grade" type="button" class="btn btn-primary grade-save">保存</button>
            </div>
        </div>
    </div>
</div>
<!--添加班级-->
<div class="modal fade" id="addClass" tabindex="-1" role="dialog" data-backdrop="static">
    <div class="modal-dialog modal-sm">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                <h4 class="modal-title">添加班级</h4>
            </div>
            <div class="modal-body">
                <div id="wrong_class"></div>
                <table class="table tb-class">
                    <tr>
                        <td>班级名称</td>
                        <td><input id="txt_classname" type="text" class="form-control"/></td>
                    </tr>
                    <tr>
                        <td>班主任</td>
                        <td><input id="txt_teacher" type="text" class="form-control"/></td>
                    </tr>
                    <tr>
                        <td>协管员</td>
                        <td><input id="txt_assistant" type="text" class="form-control"/></td>
                    </tr>
                </table>
            </div>
            <div class="modal-footer">
                <button id="btn_save_class" type="button" class="btn btn-primary grade-save">保存</button>
            </div>
        </div>
    </div>
</div>
</body>
</html>