define(function(require, exports, module) {

    exports.init = function(){
    }

    function manager(){
        $(".devices").click(function(){
            var card_guid = $(this).attr("data-id");
            var url = $(this).attr("data-href");
            location.href = url+"?card="+card_guid;
        })
    }
});