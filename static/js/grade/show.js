
define(function(require, exports, module) {

    exports.init = function(){
        $(document).ready(function(){
           $("#tb_grade tbody tr").each(function(){
               var tr = $(this);
               var guid = tr.attr("id");
               if(guid&&guid!=""){
                   getClassInfo(guid);
               }
           })
        })
    }
    function getClassInfo(guid){
        var url = "/class/list"
        var data = {
            "guid":guid
        }
        $.post(url,data,function(data){
            if(data.Code=="0"){
                //找到对应的年级元素
                var el = $("#"+guid);
                if(data.Data!=null&&data.Data.length>0) {

                    for (var i = 0; i < data.Data.length; i++) {
                        var html = "<tr><td>" + data.Data[i].Classname + "</td><td>" + data.Data[i].Teacher + "</td><td>" + data.Data[i].Assist + "</td></tr>";
                        el.after(html)
                    }
                    el.find("td").first().attr("rowspan",data.Data.length+1)
                }
            }
        })
    }
});
