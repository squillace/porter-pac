package pac

import (
	"fmt"

	"get.porter.sh/porter/pkg/exec/builder"
	yaml "gopkg.in/yaml.v2"
)

// BuildInput represents stdin passed to the mixin for the build command.
type BuildInput struct {
	Config MixinConfig
}

// MixinConfig represents configuration that can be set on the pac mixin in porter.yaml
// mixins:
// - pac:
//	  clientVersion: "v0.0.0"

type MixinConfig struct {
	ClientVersion string `yaml:"clientVersion,omitempty"`
}

// This is an example. Replace the following with whatever steps are needed to
// install required components into
const dockerfileLines = `RUN apt-get update && \
apt-get install zip unzip gpg tree curl -y && \
mkdir -p cnab/app/ && \
curl -l https://dotnet.microsoft.com/download/dotnet/scripts/v1/dotnet-install.sh -o cnab/app/dotnet-install.sh && \ 
chmod +x cnab/app/dotnet-install.sh && \
./cnab/app/dotnet-install.sh --version 5.0.301`

// Build will generate the necessary Dockerfile lines
// for an invocation image using this mixin
func (m *Mixin) Build() error {

	// Create new Builder.
	var input BuildInput

	err := builder.LoadAction(m.Context, "", func(contents []byte) (interface{}, error) {
		err := yaml.Unmarshal(contents, &input)
		return &input, err
	})
	if err != nil {
		return err
	}

	suppliedClientVersion := input.Config.ClientVersion

	if suppliedClientVersion != "" {
		m.ClientVersion = suppliedClientVersion
	}

	fmt.Fprintf(m.Out, dockerfileLines)
	fmt.Fprintf(m.Out, "\nENV DOTNET_ROOT=/root/.dotnet\n")
	fmt.Fprintf(m.Out, "\nENV DOTNET_SYSTEM_GLOBALIZATION_INVARIANT=1")
	fmt.Fprintf(m.Out, "\nENV PATH=$HOME/cnab/app/tools/:${PATH}")
	// Example of pulling and defining a client version for your mixin
	// fmt.Fprintf(m.Out, "\nRUN curl https://get.helm.sh/helm-%s-linux-amd64.tar.gz --output helm3.tar.gz", m.ClientVersion)

	return nil
}
