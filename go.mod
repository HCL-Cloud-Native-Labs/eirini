module code.cloudfoundry.org/eirini

go 1.16

replace (
	k8s.io/api => k8s.io/api v0.20.3
	k8s.io/client-go => k8s.io/client-go v0.20.3
)

require (
	cloud.google.com/go v0.81.0 // indirect
	code.cloudfoundry.org/cfhttp/v2 v2.0.0
	code.cloudfoundry.org/lager v2.0.0+incompatible
	code.cloudfoundry.org/runtimeschema v0.0.0-20180622184205-c38d8be9f68c
	code.cloudfoundry.org/tlsconfig v0.0.0-20200131000646-bbe0f8da39b3
	github.com/Azure/go-autorest/autorest v0.11.18 // indirect
	github.com/Microsoft/go-winio v0.5.0 // indirect
	github.com/cloudfoundry/tps v0.0.0-20210303221408-6295d5712990
	github.com/containers/image v3.0.2+incompatible
	github.com/containers/storage v1.25.0 // indirect
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker v20.10.6+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.6.3 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/go-logr/logr v0.4.0
	github.com/gofrs/flock v0.8.0
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/go-cmp v0.5.5
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/googleapis/gnostic v0.5.5 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/go-retryablehttp v0.7.0
	github.com/hashicorp/go-uuid v1.0.2
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/jessevdk/go-flags v1.5.0
	github.com/jinzhu/copier v0.3.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/maxbrunsfeld/counterfeiter/v6 v6.4.1
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/onsi/ginkgo v1.16.1
	github.com/onsi/gomega v1.11.0
	github.com/opencontainers/image-spec v1.0.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.10.0
	github.com/prometheus/common v0.23.0
	github.com/sirupsen/logrus v1.8.1 // indirect
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
	golang.org/x/mod v0.4.2 // indirect
	golang.org/x/net v0.0.0-20210427231257-85d9c07bbe3a // indirect
	golang.org/x/oauth2 v0.0.0-20210427180440-81ed05c6b58c // indirect
	golang.org/x/sys v0.0.0-20210426230700-d19ff857e887 // indirect
	golang.org/x/term v0.0.0-20210422114643-f5beecf764ed // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	gomodules.xyz/jsonpatch/v2 v2.1.0
	google.golang.org/genproto v0.0.0-20210406143921-e86de6bf7a46 // indirect
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	k8s.io/api v0.21.0
	k8s.io/apiextensions-apiserver v0.21.0 // indirect
	k8s.io/apimachinery v0.21.0
	k8s.io/client-go v1.5.2
	k8s.io/code-generator v0.21.0
	k8s.io/klog v1.0.0
	k8s.io/kube-openapi v0.0.0-20210421082810-95288971da7e // indirect
	k8s.io/utils v0.0.0-20210305010621-2afb4311ab10 // indirect
	sigs.k8s.io/controller-runtime v0.8.3
	sigs.k8s.io/structured-merge-diff/v4 v4.1.1 // indirect
)
