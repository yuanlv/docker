package cli

import flag "github.com/docker/docker/pkg/mflag"

// ClientFlags represents flags for the docker client.
// ClientFlags 结构体代表docker client的标记信息
type ClientFlags struct {
	FlagSet   *flag.FlagSet
	Common    *CommonFlags
	PostParse func()

	ConfigDir string
}
