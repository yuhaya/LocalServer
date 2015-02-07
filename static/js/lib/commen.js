define(function(require, exports, module) {

    exports.init = function(){
        init_style();
    }

    function init_style(){
        var win_h = $(window).height();
        $("body").height(win_h);
    }

});