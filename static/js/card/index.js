define(function(require, exports, module) {

    exports.init = function(){
        Manager()
    }

    function Manager(){
        $("#manager").click(function(){
            location.href = $(this).attr("data-href");
        });
    }

});