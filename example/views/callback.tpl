{{template "header"}}
<title>pay - Yuansfer</title>
</head>

<body onload="startTime()">
    <div class="container" style="width: 500px;">
        {{template "navbar" .}}
        <form class="form-horizontal" method="POST" action="/">
            <div class="form-group">
                <label class="col-lg-4 control-label">金额(USD)：</label>
                <div class="col-lg-4">
                    {{.Amt}}
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">人民币金额(元)：</label>
                <div class="col-lg-4">
                    {{.RmbAmt}}
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">交易方式：</label>
                <div class="col-lg-4">
                    {{.Vendor}}
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">流水号：</label>
                <div class="col-lg-4">
                    {{.Reference}}
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">附言：</label>
                <div class="col-lg-4">
                    {{.Note}}
                </div>
            </div>
            <div class="form-group">
                <div class="col-lg-offset-2 col-lg-10">
                    <a href="/">
                        <button class="btn btn-default" onclick="return backToHome();">返回</button>
                    </a>
                    <script type="text/javascript">
                    function backToHome() {
                        window.location.href = "/";
                        return false;
                    }
                    </script>
                </div>
            </div>
        </form>
    </div>
</body>

</html>