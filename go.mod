module github.com/choria-io/go-choria

go 1.16

// shuts up vulnerability alerts that did not affect this project
replace github.com/opencontainers/runc v0.0.0-20161107232042-8779fa57eb4a => github.com/opencontainers/runc v1.0.3

require (
	github.com/AlecAivazis/survey/v2 v2.3.2
	github.com/Freman/eventloghook v0.0.0-20191003051739-e4d803b6b48b
	github.com/aelsabbahy/goss v0.3.16
	github.com/antonmedv/expr v1.9.0
	github.com/awesome-gocui/gocui v1.1.0
	github.com/brutella/hc v1.2.5
	github.com/choria-io/go-updater v0.0.4-0.20211231123842-da243cbc508c
	github.com/cloudevents/sdk-go/v2 v2.8.0
	github.com/fatih/color v1.13.0
	github.com/ghodss/yaml v1.0.0
	github.com/gofrs/uuid v4.2.0+incompatible
	github.com/golang-jwt/jwt/v4 v4.2.0
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.5.7
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510
	github.com/gosuri/uiprogress v0.0.1
	github.com/guptarohit/asciigraph v0.5.2
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51
	github.com/looplab/fsm v0.3.0
	github.com/miekg/pkcs11 v1.1.1
	github.com/mitchellh/mapstructure v1.4.3
	github.com/nats-io/jsm.go v0.0.28-0.20220204133248-560e881d29a4
	github.com/nats-io/nats-server/v2 v2.7.2
	github.com/nats-io/nats.go v1.13.1-0.20220206022646-bf1b00513eab
	github.com/nats-io/natscli v0.0.29-0.20220204133629-ccbbdfae9919
	github.com/olekukonko/tablewriter v0.0.5
	github.com/onsi/ginkgo/v2 v2.1.1
	github.com/onsi/gomega v1.18.1
	github.com/open-policy-agent/opa v0.37.1
	github.com/prometheus/client_golang v1.12.1
	github.com/prometheus/client_model v0.2.0
	github.com/robfig/cron v1.2.0
	github.com/sirupsen/logrus v1.8.1
	github.com/tidwall/gjson v1.14.0
	github.com/tidwall/pretty v1.2.0
	github.com/xeipuuv/gojsonschema v1.2.0
	go.uber.org/atomic v1.9.0
	golang.org/x/crypto v0.0.0-20220131195533-30dcbda58838
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220204135822-1c1b9b1eba6a
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211
	golang.org/x/tools v0.1.9
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	rsc.io/goversion v1.2.0
)
