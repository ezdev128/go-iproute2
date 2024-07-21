module github.com/ezdev128/go-iproute2

go 1.19

require (
	github.com/Asphaltt/go-iproute2 v0.0.0-20220404100349-2854774d085e
	github.com/mdlayher/netlink v1.7.2
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/spf13/cobra v1.8.1
	golang.org/x/sys v0.22.0
)

require (
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/josharian/native v1.1.0 // indirect
	github.com/mdlayher/socket v0.4.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/tklauser/numcpus v0.4.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
)

replace github.com/shirou/gopsutil => ./pkg/gopsutil
