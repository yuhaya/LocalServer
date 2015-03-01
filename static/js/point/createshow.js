/**
 * Created by mao on 3/1/15.
 */
define(function(require, exports, module) {

    require("timepicker_addon")
    require("validate_helper");

    exports.init = function(){
        $('#time').datetimepicker({
            timeFormat: 'HH:mm:ss',
            dateFormat: 'yy-mm-dd'
        });
        var form =  $("#att_form")
        form.validate(function($form, e){  })
    }

});