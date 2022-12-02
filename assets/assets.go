package assets

//nolint:gocritic,golint
import _ "embed"

//nolint:gochecknoglobals,gocritic
var (
	//go:embed dark.ico
	IconDark []byte

	//go:embed light.ico
	IconLight []byte

	//go:embed color.ico
	IconColor []byte
)
