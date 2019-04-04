{{template "header"}}
<title>inquire - Yuansfer</title>
</head>

<body onload="startTime()">
    <div class="container" style="width: 500px;">
        {{template "navbar" .}}
        <form class="form-horizontal" method="POST" action="/">
            <div class="form-group">
                <label class="col-lg-4 control-label">TransactionNo</label>
                <label class="col-lg-4">{{.TransactionNo}}</label>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">OriginalTransactionNo</label>
                <label class="col-lg-4">{{.OriginalTransactionNo}}</label>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">Amount</label>
                <label class="col-lg-4">{{.Amount}}</label>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">Vendor:</label>
                <label class="col-lg-4">{{.Vendor}}</label>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">MerchantNo</label>
                <label class="col-lg-4">{{.MerchantNo}}</label>
            </div>
            <div class="form-group">
                <div class="col-lg-offset-2 col-lg-10">
                    <button class="btn btn-default" onclick="return backToHome();">返回</button>
                    <script type="text/javascript">
                    function backToHome() {
                        window.location.href = "/instore-add";
                        return false;
                    }
                    </script>
                </div>
            </div>
        </form>
    </div>
</body>

</html>