<script type="application/javascript" src="/static/js/lib/bootstrap-3.3.2/js/bootstrap.timepicker.min.js"></script>
<script type="application/javascript" src="/static/js/lib/jquery-ui-timepicker-addon-i18n.js"></script>
<link rel="stylesheet" href="/static/js/lib/bootstrap-3.3.2/css/bootstrap-datepicker.css">

<div class="row" style="margin-top: 20px">
    <div class="col-md-2"></div>
    <div class="col-md-8">
        <form action="{{urlfor "PointController.Create"}}" name="att_form" method="post" id="att_form">
            <div class="form-group control-group">
                <label for="card">卡号</label>
                <input type="text" name="card" class="form-control validate" data-display="卡号" data-rules="required|max_length[50]"  id="card" placeholder="输入卡号">
                <p class="help-block"></p>
            </div>
            <div class="form-group control-group">
                <label for="time">刷卡时间</label>
                <input type="datetime" name="time" class="form-control validate" data-display="刷卡时间" data-rules="required"  id="time" placeholder="输入时间">
                <p class="help-block"></p>
            </div>

            <div class="radio">
                <label>
                    <input type="radio" name="type" id="type" value="0" checked>
                    入校
                </label>
            </div>
            <div class="radio">
                <label>
                    <input type="radio" name="type" id="type2" value="1">
                    出校
                </label>
            </div>
            <input type="hidden" name="auto" value="0"/>
            <div class="form-group control-group">
                <label for="device">刷卡机</label>
                <input type="text" class="form-control" name="device validate" id="device" placeholder="刷卡机标识">
                <p class="help-block"></p>
            </div>

            <div class="radio">
                <label>
                    <input type="radio" name="kind" id="kind" value="0" checked>
                    入校刷卡机
                </label>
            </div>
            <div class="radio">
                <label>
                    <input type="radio" name="kind" id="kind2" value="1">
                    出校刷卡机
                </label>
            </div>

            <input type="hidden" value="" name="vmp"/>

            <button type="submit" class="btn btn-default">Submit</button>
        </form>
    </div>
    <div class="col-md-2"></div>
</div>