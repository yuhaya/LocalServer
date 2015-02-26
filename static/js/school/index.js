/**
 * Created by zhengping on 15/2/25.
 */
define(function (require, exports, module) {

    exports.init = function () {
        $(document).ready(function () {
            $("#door").keydown(onlyNum)
        })
        save();
    }

    function onlyNum() {
        if (!(event.keyCode == 46) && !(event.keyCode == 8) && !(event.keyCode == 37) && !(event.keyCode == 39))
            if (!((event.keyCode >= 48 && event.keyCode <= 57) || (event.keyCode >= 96 && event.keyCode <= 105)))
                event.returnValue = false;
    }

    function save(){
        $("#save").click(function(){
            var guid = $("#guid").val();
            var name = $("#name").val();
            var province = $("#province").val();
            var city = $("#city").val();
            var county = $("#county").val();
            var address = $("#address").val();
            var door=$("#door").val();

            var data = {
                "guid":guid,
                "name":name,
                "province":province,
                "city":city,
                "county":county,
                "address":address,
                "door":door
            }
            var url = "/school/create"
            $.post(url,data,function(data){
                if(data.Code=='0'){
                    alert("保存成功")
                }else{
                    alert("保存失败")
                }
            });
        });
    }

});