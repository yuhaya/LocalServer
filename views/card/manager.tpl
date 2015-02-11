<div id="con">
    {{ if ne .err "" }}
        {{.err}}
    {{ else }}
    <ul id="device_list">
        {{ if eq .num 0 }}
             <li>未添加设备信息</li>
        {{ else }}
                {{range .device_list}}
                    <li data-id="{{ .Guid }}" class="devices" data-href="{{ urlfor "CardController.Show" }}">
                    <span>序列号/IP:</span><span>{{ .Device }}</span><span>{{ .Description }}</span>
                    {{ if eq .Status 1 }}
                    <span>连接正常</span>
                    {{ else }}
                    <span>连接异常</span>
                    {{ end }}
                    </li>
                {{end}}
        {{ end }}
    </ul>
    {{ end }}
</div>