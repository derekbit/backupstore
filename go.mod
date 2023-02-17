module github.com/longhorn/backupstore

go 1.13

replace (
	k8s.io/apimachinery => k8s.io/apimachinery v0.23.1
	k8s.io/mount-utils => k8s.io/mount-utils v0.23.1
)

require (
	github.com/aws/aws-sdk-go v1.34.2
	github.com/google/fscrypt v0.3.3
	github.com/google/uuid v1.3.0
	github.com/honestbee/jobq v1.0.2
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.3.0
	github.com/spf13/afero v1.5.1
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli v1.22.5
	go.uber.org/multierr v1.9.0
	golang.org/x/net v0.0.0-20211209124913-491a49abca63
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f
	k8s.io/apimachinery v0.26.3
	k8s.io/mount-utils v0.26.3
)
