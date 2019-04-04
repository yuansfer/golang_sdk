{{template "header"}}
<title>inquire - Yuansfer</title>
</head>

<body onload="startTime()">
    <div class="container" style="width: 500px;">
        {{template "navbar" .}}
        <form class="form-horizontal" method="POST" action="/">
            <div class="form-group">
                <label class="col-lg-4 control-label">金额：</label>
                <label class="col-lg-4">{{.Amount}} {{.Currency}}</label>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">交易结果：</label>
                <label class="col-lg-4">{{.Status}}</label>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">流水号：</label>
                <label class="col-lg-4">{{.Reference}}</label>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">YuansferId:</label>
                <label class="col-lg-4">{{.YuansferId}}</label>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">------退款信息------</label>
            </div>
             <table border="10" cellpadding="0" cellspacing="0" style="width:100%;margin:15px 0px;">
                <tbody>
                    {{range .RefundInfo}}
                        <div class="form-group">
                            <label class="col-lg-4 control-label">退款ID：</label>
                            <label class="col-lg-4">{{.RefundYuansferId}}</label>
                        </div>
                        <div class="form-group">
                            <label class="col-lg-4 control-label">退款金额：</label>
                            <label class="col-lg-4">{{.RefundAmount}} {{$.Currency}}</label>
                        </div>
                    {{end}}
                </tbody>
            </table>
            <div class="form-group">
                <div class="col-lg-offset-2 col-lg-10">
                    <button class="btn btn-default" onclick="return backToHome();">返回</button>
                    <script type="text/javascript">
                    function backToHome() {
                        window.location.href = "/inquire";
                        return false;
                    }
                    </script>
                </div>
            </div>
        </form>
    </div>
</body>

</html>