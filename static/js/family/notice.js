/**
 * Created by mao on 3/7/15.
 */
/**
 * Created by mao on 2/28/15.
 */
define(function(require, exports, module) {

    var FormValidator = require("validate");
    exports.init = function(){

        //location.href = location.href
        var card = "";

        $(".bcard").click(function(){
            card = $(this).attr("data-val");
        });

        $("#card_form_submit").click(function(){
            $("#card").val(card);
            $("#card_form").submit()
        });

    }

});