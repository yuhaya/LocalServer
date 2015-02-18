/**
 * Created by mao on 2/15/15.
 */
define(function(require, exports, module) {
    require("validate_helper");
    exports.init = function(){
        var form =  $("#form_add_memeber")
        form.validate(function($form, e){  })
    }

});