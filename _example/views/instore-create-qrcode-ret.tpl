{{template "header"}}
<title>inquire - Yuansfer</title>
</head>

<body onload="startTime()">
    <div class="container" style="width: 500px;">
        {{template "navbar" .}}
        <form class="form-horizontal" method="POST" action="/">
            <div class="form-group">
                <label class="col-lg-4 control-label">Reference</label>
                <label class="col-lg-4">{{.Reference}}</label>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">RectCode</label>
                <label class="col-lg-4">{{.RectCode}}</label>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">DeepLink</label>
                <label class="col-lg-4">{{.DeepLink}}</label>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">TransactionNo:</label>
                <label class="col-lg-4">{{.TransactionNo}}</label>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">TimeOut</label>
                <label class="col-lg-4">{{.TimeOut}}</label>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">QrcodeUrl</label>
                <label class="col-lg-4">{{.QrcodeUrl}}</label>
            </div>
            <div class="form-group">
                <div class="col-lg-offset-2 col-lg-10">
                    <button class="btn btn-default" onclick="return backToHome();">返回</button>
                    <script type="text/javascript">
                    function backToHome() {
                        window.location.href = "/instore-qrcode";
                        return false;
                    }
                    </script>
                </div>
            </div>
        </form>
    </div>
</body>

</html>