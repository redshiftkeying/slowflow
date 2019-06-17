module github.com/redshiftkeying/slowflow

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
	golang.org/x/net => github.com/golang/net v0.0.0-20180811021610-c39426892332
	golang.org/x/sys => github.com/golang/sys v0.0.0-20180810173357-98c5dad5d1a0
	golang.org/x/text => github.com/golang/text v0.3.0
)

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/graph-gophers/graphql-go v0.0.0-20190513003547-158e7b876106
)
