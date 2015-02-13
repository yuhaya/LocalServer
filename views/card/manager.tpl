<div id="con">
    <div class="page-header">
        <h1>卡片管理 <small>选择刷卡机</small></h1>
    </div>
    {{ if ne .err "" }}
        {{.err}}
    {{ else }}
    <div id="device_list">

        {{ if eq .num 0 }}
            <div class="col-md-12">系统还未添加刷卡机设备</div>
        {{ else }}
                {{range .device_list}}
                    <div class="row devices" data-id="{{ .Guid }}" data-href="{{ urlfor "CardController.Show" }}">
                        <div class="col-md-7">
                                <span>序列号/IP:</span><span>{{ .Device }}</span>
                        </div>
                        <div class="col-md-3">{{ .Description }}</div>
                        <div class="col-md-2">
                            {{ if eq .Status 1 }}
                                连接正常
                            {{ else }}
                                连接异常
                            {{ end }}
                        </div>
                    </div>
                {{end}}
        {{ end }}
    </div>
    {{ end }}
</div>