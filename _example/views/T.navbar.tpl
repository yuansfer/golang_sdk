{{define "navbar"}}
<!-- <a class="navbar-brand" href="/">Yuansfer</a> -->
<div class="navbar navbar-inverse navbar-fixed-top">
    <div class="container">
        <a class="navbar-brand" href="/">
            <img alt="Brand" src="/static/img/7372_logo-invoice.jpg" height="20">
		</a>
        <div>
            <ul class="nav navbar-nav">
                <li {{if .IsPay}}class="active" {{end}}><a href="/">Online Pay</a></li>
                <li {{if .IsInquire}}class="active" {{end}}><a href="/inquire">Online Query</a></li>
                <li {{if .IsRefund}}class="active" {{end}}><a href="/refund">Refund</a></li>
                <li {{if .IsInstoreAdd}}class="active" {{end}}><a href="/instore-add">In-store Add</a></li>
                <li {{if .IsInstorePay}}class="active" {{end}}><a href="/instore-pay" target="_blank">In-store Pay</a></li>
                <li {{if .IsCreateQrcode}}class="active" {{end}}><a href="/instore-qrcode">In-store Qrcode</a></li>
                <li {{if .IsReverse}}class="active" {{end}}><a href="/instore-reverse">In-store Reverse</a></li>
                <li {{if .IsMicroPay}}class="active" {{end}}><a href="/micropay">micropay</a></li>
            </ul>
        </div>
        <!-- <div class="pull-right">
            <ul class="nav navbar-nav">
                <li>
                    <a class="navbar-brand" href="/">
                        <img alt="Brand" src="/static/img/7372_logo-invoice.png" height="20">
                    </a>
                </li>
            </ul>
            <a id="txt"></a>
        </div> -->
    </div>
</div>
<br/>
<br/>
<br/>
<br/>
<br/>
{{end}}