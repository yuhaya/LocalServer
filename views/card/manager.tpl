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
        <div class="list-group">
                {{range .device_list}}
            <a href="#" class="list-group-item devices" data-id="{{ .Guid }}" data-href="{{ urlfor "CardController.Show" }}">
                        <h4 class="list-group-item-heading"><span>序列号/IP:</span>{{ .Device }}</h4>
                        <p>{{ .Description }}</p>
                        <p>
                            {{ if eq .Status 1 }}
                                连接正常
                            {{ else }}
                                连接异常
                            {{ end }}
                        </p>
            </a>
                {{end}}
        </div>
        {{ end }}
    </div>
    {{ end }}
</div>