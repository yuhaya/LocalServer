define(function(require, exports, module) {

    require("validate_helper");

    exports.init = function(){
        var form =  $("#form_add_family")
        form.validate(function($form, e){  })
    }

});