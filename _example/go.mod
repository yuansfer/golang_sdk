module github.com/yuansfer/golang_sdk/_example

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190403202508-8e1b8d32e692
	golang.org/x/net => github.com/golang/net v0.0.0-20190403144856-b630fd6fe46b
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190403152447-81d4e9dc473e
	golang.org/x/text => github.com/golang/text v0.3.0
)

require (
	github.com/astaxie/beego v1.11.1
	github.com/yuansfer/golang_sdk v0.0.0-20190305150004-15c472ee0f8c
)
