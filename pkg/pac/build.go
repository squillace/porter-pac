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
./cnab/app/dotnet-install.sh --version 5.0.301 && \
mkdir -p /cnab/app/tools`

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

	fmt.Fprintln(m.Out, dockerfileLines)
	fmt.Fprintln(m.Out, "ENV DOTNET_ROOT=/root/.dotnet")
	fmt.Fprintln(m.Out, "ENV DOTNET_SYSTEM_GLOBALIZATION_INVARIANT=1")
	fmt.Fprintln(m.Out, "ENV PATH=$HOME/cnab/app/tools/:${PATH}")
	// Example of pulling and defining a client version for your mixin
	fmt.Fprintf(m.Out, "RUN curl -L https://www.nuget.org/api/v2/package/Microsoft.PowerApps.CLI.Core.linux-x64/%s --output /cnab/app/pac.nuget", m.ClientVersion)
	fmt.Fprintln(m.Out, "\nRUN cd /cnab/app/ && unzip /cnab/app/pac.nuget && chmod +x /cnab/app/tools/pac")

	return nil
}
