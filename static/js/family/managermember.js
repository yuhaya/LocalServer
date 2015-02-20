/**
 * Created by mao on 2/15/15.
 */
/**
 * Created by mao on 2/15/15.
 */
define(function(require, exports, module) {
    require("validate_helper");
    exports.init = function(){
        //var form =  $("#form_edit_family")
        //form.validate(function($form, e){  })
    }

    /**
     * 设置主家长
     */
    function set_main(){
        $(".main_parent").click(function(){
            var url = $(this).attr("data-url")
            $.ajax({
                url:url,
                type:"POST",
                dataType:"json",
                success:function(data){
                    
                }
            })
        })
    }

});