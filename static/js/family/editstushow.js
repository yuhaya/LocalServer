/**
 * Created by mao on 2/24/15.
 */
define(function(require, exports, module) {
    require("validate_helper");
    exports.init = function(){

        var form =  $("#form_mod_memeber")
        form.validate(function($form, e){  })


        if (data.type == 'stu'){
            change_class()
            $("#Grade_guid").change(change_class)
        }
    }

    function change_class(){

        grade_class = $.parseJSON(data.grade_class)

        var grade = $("#Grade_guid").val()
        var class_val = grade_class[grade]
        var str = "<option>选择班级</option>"
        for ( var i = 0 ;i < class_val.length;i++){
            console.log(class_val[i])
            str+= "<option value="+class_val[i].guid+">"+class_val[i].name+"</option>"
        }
        $("#Class_guid").html(str)
    }

});