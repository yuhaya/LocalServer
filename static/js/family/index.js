define(function(require, exports, module) {

    exports.init = function(){

        $(".DelFm").click(function(){
            var res = window.top.confirm("删除家庭将删除该家庭的所有的学生以及家长!")
            if (res){
                location.href = $(this).attr("data-href");
            }
        })
    }

});