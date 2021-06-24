module github.com/haproxytech/kubernetes-ingress

go 1.16

require (
	github.com/go-test/deep v1.0.7
	github.com/google/renameio v1.0.1
	github.com/haproxytech/client-native/v2 v2.5.1-0.20210902093307-1f696b917f86
	github.com/haproxytech/config-parser/v4 v4.0.0-rc1.0.20210902180329-0171d9c29239
	github.com/jessevdk/go-flags v1.4.0
	github.com/stretchr/testify v1.6.1
	k8s.io/api v0.21.3
	k8s.io/apimachinery v0.21.3
	k8s.io/client-go v0.21.3
)

replace github.com/haproxytech/client-native/v2 => gitlab.int.haproxy.com/mmhedhbi/client-native/v2 v2.0.0-20210919215436-d3a96d518a82
