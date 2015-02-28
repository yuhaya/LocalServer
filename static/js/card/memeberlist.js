/**
 * Created by mao on 2/28/15.
 */
define(function(require, exports, module) {

    var FormValidator = require("validate");
    exports.init = function(){

        //location.href = location.href
        var guid = "";
        var type = 0;

        $(".bcard").click(function(){
            guid = $(this).attr("data-val");
            type = $(this).attr("data-type");
        });

        $("#card_form_submit").click(function(){
            $("#memguid").val(guid);
            $("#memtype").val(type);
            $("#card_form").submit()
        });

        var validator = new FormValidator('card_form', [{
            name: 'card',
            display: '卡号',
            rules: 'required|min_length[3]|max_length[40]'
        }], function(errors, event) {
            if (errors.length > 0) {
                var errorString = '';

                for (var i = 0, errorLength = errors.length; i < errorLength; i++) {
                    errorString += errors[i].message;
                }

                $("#error").text(errorString);
            }else{

            }
        });
    }

});