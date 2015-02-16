define(function (require, exports, module) {

    exports.init = function () {
        manageGrade();
        orderbyGrade();
        manageClass();
    }
    //年级管理**************
    function manageGrade() {
        var repeat = false;//重名标识
        $("#btn_add_grade").click(function () {
            $("#addGrade").modal().on("hide.bs.modal", function () {
                location.reload();//关闭后刷新页面
            });
        });
        $("#btn_append_grade").click(function () {
            var txt_grade = $("#txt_grade_name");
            var text = txt_grade.val().trim();
            if (text == "") {//名称为空提示错误
                showWarning($("#wrong_grade"), "年级名称不能为空");
                txt_grade.focus();
                return;
            }
            //判断是否有重名
            $("#list_grade").find("a").each(function () {
                var a = $(this);
                if (a.text() != "" && a.text() == text) {
                    showWarning($("#wrong_grade"), "该年级名称已存在");
                    txt_grade.focus();
                    repeat = true;
                    return;
                }
            });
            if (!repeat) {
                addGrade(function (data) {
                    if (data.Code == "0") {
                        $("#list_grade").append("<a href=\"#\" data=\"" + data.Data.Rating + "\" pid=\"" + data.Data.Id + "\" guid=\"" + data.Data.Guid + "\" sid=\"" + data.Data.SchoolGuid + "\" class=\"list-group-item\">" + text + "</a>");
                        orderbyGrade();
                    }
                    else
                        showWarning($("#wrong_grade"), data.Msg);
                });
            }
        });
        $("#btn_save_grade").click(function () {
            updateGarde(function (data) {
                if (data.Code != "0") {
                    showWarning($("#wrong_grade"), data.Msg);
                } else {
                    $("#addGrade").modal("hide")
                }
            })
        });
    }

    //年级排序
    var oa;//当前被选中的元素
    function orderbyGrade() {
        //切换显示状态
        $("#list_grade").find("a").each(function () {
            var a = $(this);
            a.unbind("click").click(function () {
                $("#list_grade").find("a").removeClass("active");
                a.addClass("active");
                oa = a;
            });
        });
        //删除元素
        $("#btn_del").unbind("click").click(function () {
            if (oa) {
                deleteGrade(oa, function (data) {
                    if (data.Code == "0") {
                        oa.remove();
                        oa = undefined;
                    } else {
                        return showWarning($("#wrong_grade"), data.Msg);
                    }
                });
            }
        });
        //上移元素
        $("#btn_up").unbind("click").click(function () {
            if (oa) {
                //找到上一个相邻元素
                var pre = oa.prev();
                if (!pre)
                    return;
                pre.before(oa);

            }
        });
        //下移元素
        $("#btn_down").unbind("click").click(function () {
            if (oa) {
                var after = oa.next();
                if (!after)
                    return;
                if (after)
                    after.after(oa);
            }
        });
    }

    //添加年级信息
    function addGrade(callback) {
        var name = $("#txt_grade_name").val();
        if (name && name != "") {
            var url = "/grade/create"
            var data = {"name": name}
            $.post(url, data, callback)
        }
    }

    //删除所选年级
    function deleteGrade(el, callback) {
        var guid = el.attr("guid");
        if (guid && guid != "") {
            var url = "/grade/delete";
            var data = {
                "guid": guid
            }
            $.post(url, data, callback);
        }
    }

    //更新排序信息
    function updateGarde(callback) {
        var data = "{\"grades\":[";
        var num = 1;
        //获取所有信息
        $("#list_grade").find("a").each(function () {
            var a = $(this);
            data += "{";
            data += "\"id\":" + a.attr("pid") + ",";
            data += "\"guid\":\"" + a.attr("guid") + "\",";
            data += "\"name\":\"" + a.text() + "\",";
            data += "\"rating\":" + (num++);
            data += "},";
        })
        data = data.substring(0, data.length - 1);
        data += "]}";
        $.post("/grade/update", {"data": data}, callback);
    }

    //显示错误信息
    function showWarning(obj, msg) {
        $(obj).empty().append("<div class=\"alert alert-danger\" role=\"alert\"><button type=\"button\" class=\"close\" data-dismiss=\"alert\" aria-label=\"Close\"><span aria-hidden=\"true\">&times;</span></button>" + msg + "</div>");
        setTimeout(function () {
            $(obj).empty();
        }, 5000)
    }

    //班级管理***************
    function manageClass() {
        $("#display-grade").find("a").each(function () {
            var a = $(this);
            a.unbind("click").click(function () {
                //将添加班级按钮启用
                $("#btn_add_class").removeClass("disabled");
                //切换当前显示状态
                $("#display-grade").find("a").removeClass("active");
                a.addClass("active");
                oa = a;
                //根据年级guid获取班级信息
                getClassList()
            })
        });

        $("#btn_add_class").click(function () {
            $("#addClass").modal()
        })

        $("#btn_save_class").click(function () {
            addClass()
        })
    }

    //根据年级guid获取班级信息
    function getClassList() {
        var guid = oa.attr("guid");

        if (guid && guid != "") {
            var url = "/class/list"
            var data = {
                "guid": guid
            }
            $.post(url, data, function (data) {
                if (data.Code == "0") {
                    $("#tb_class tbody").empty();
                    for (var i = 0; i < data.Data.length; i++) {
                        var html = "<tr>" +
                            "<td>" + data.Data[i].Classname + "</td>" +
                            "<td>" + data.Data[i].Teacher + "</td>" +
                            "<td>" + data.Data[i].Assist + "</td>" +
                            "<td><span guid=\"" + data.Data[i].Guid + "\" grade=\"" + data.Data[i].GradeGuid + "\" tid=\"" + data.Data[i].TeacherGuid + "\" aid=\"" + data.Data[i].AssistantGuid + "\" class=\"glyphicon glyphicon-remove pointer\" title=\"删除\"></span></td>" +
                            "</tr>";
                        $("#tb_class tbody").append(html)
                    }
                    delClass();
                }
            })
        }
    }

    //删除班级
    function delClass() {

        $("#tb_class tbody").find("span").each(function () {
            var s = $(this);
            s.unbind("click").click(function () {
                if (window.confirm("您确定要删除选中的班级？")) {
                    var url = "/class/delete"
                    var data = {
                        "grade_guid": s.attr("grade"),
                        "class_guid": s.attr("guid"),
                        "teacher_guid": s.attr("tid"),
                        "assistant_guid": s.attr("aid")
                    }
                    $.post(url, data, function (data) {
                        if (data.Code == "0") {
                            getClassList();
                        }
                    });
                }
            });
        });

    }

    //添加班级
    function addClass() {
        //获取年级guid
        var guid, classname, teacher, assistant;
        if (oa) {
            guid = oa.attr("guid")
        } else {
            return showWarning($("#wrong_class"), "请先选择年级");
        }

        classname = $("#txt_classname").val().trim();
        if (classname == "") {
            showWarning($("#wrong_class"), "请输入班级名称");
            return $("#txt_classname").focus()
        }
        teacher = $("#txt_teacher").val().trim();
        assistant = $("#txt_assistant").val().trim();
        var url = "/class/create"
        var data = {
            "guid": guid,
            "classname": classname,
            "teacher": teacher,
            "assistant": assistant
        }
        $.post(url, data, function (data) {
            if (data.Code == "0") {
                $("#addClass").modal('hide');
                $("#txt_classname").val("");
                $("#txt_teacher").val("");
                $("#txt_assistant").val("");
                getClassList()
            }
        })
    }
});

