{{template "header"}}
<title>pay - Yuansfer</title>
</head>

<body onload="startTime()">
    <div class="container" style="width: 500px;">
        {{template "navbar" .}}
        <form class="form-horizontal" method="POST" action="/">
            <div class="form-group">
                <label class="col-lg-4 control-label">商户号：</label>
                <div class="col-lg-6">
                    <input id="merchantNo" class="form-control" name="merchantNo" placeholder="merchantNo">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">店铺号：</label>
                <div class="col-lg-6">
                    <input id="storeNo" class="form-control" name="storeNo" placeholder="storeNo">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">金额：</label>
                <div class="col-lg-6">
                    <input id="amt" class="form-control" name="amt" placeholder="amount">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">人民币金额：</label>
                <div class="col-lg-6">
                    <input id="rmbAmt" class="form-control" name="rmbAmt" placeholder="rmbAmt">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">支付方式：</label>
                <div class="col-lg-6">
                    <!-- <input id="vendor" class="form-control" name="vendor" placeholder="vendor"> -->
                    <tr>
                        <td>
                            <input type="hidden" class="form-control" name="vendor1" value="vendor">
                        </td>
                    </tr>
                    <tr>
                        <td>
                            <select name="vendor">
                                <option value="alipay">alipay</option>
                                <option value="wechatpay">wechatpay</option>
                                <option value="unionpay">unionpay</option>
                                <option value="creditcard">credit card</option>
                            </select>
                        </td>
                    </tr>
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">Reference：</label>
                <div class="col-lg-6">
                    <input id="reference" class="form-control" name="reference" placeholder="reference">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">Description：</label>
                <div class="col-lg-6">
                    <input id="description" class="form-control" name="description" placeholder="description">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">Note：</label>
                <div class="col-lg-6">
                    <input id="note" class="form-control" name="note" placeholder="note">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">Terminal：</label>
                <div class="col-lg-6">
                    <tr>
                        <td>
                            <input type="hidden" class="form-control" name="terminal1" value="terminal">
                        </td>
                    </tr>
                    <tr>
                        <td>
                            <select name="terminal">
                                <option value="ONLINE">ONLINE</option>
                                <option value="WAP">WAP</option>
                                <option value="APP">APP</option>
                            </select>
                        </td>
                    </tr>
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">IpnUrl：</label>
                <div class="col-lg-6">
                    <input id="ipnUrl" class="form-control" name="ipnUrl" placeholder="ipnUrl">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">CallbackUrl：</label>
                <div class="col-lg-6">
                    <input id="callbackUrl" class="form-control" name="callbackUrl" placeholder="callbackUrl">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">Goods：</label>
                <div class="col-lg-6">
                    <input id="goods" class="form-control" name="goods" placeholder="Goods">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">Quantity：</label>
                <div class="col-lg-6">
                    <input id="quantity" class="form-control" name="quantity" placeholder="Quantity">
                </div>
            </div>
            <div class="form-group">
                <div class="col-lg-offset-2 col-lg-10">
                    <a href="/" target="_blank">
                        <button type="submit" class="btn btn-default" onclick="return checkInput();">提交</button>
                    </a>
                    <button class="btn btn-default" onclick="return backToHome();">返回</button>
                    <script type="text/javascript">
                    function backToHome() {
                        window.location.href = "/";
                        return false;
                    }

                    function checkInput() {
                        // var amt = document.getElementById("amt");
                        // if (amt.value.length == 0) {
                        //  alert("请输入金额");
                        //  return false;
                        // }

                        // var rmbAmt = document.getElementById("rmbAmt");
                        // if (rmbAmt.value.length == 0) {
                        //  alert("请输入金额");
                        //  return false;
                        // }

                        var vendor = document.getElementById("vendor");
                        if (vendor.value.length == 0) {
                            alert("请输入支付方式");
                            return false;
                        }
                    }
                    </script>
                </div>
            </div>
        </form>
    </div>
</body>

</html>