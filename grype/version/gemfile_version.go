package version

import (
	"fmt"
	"strings"
)

type gemfileVersion struct {
	// keeping the raw version for transparency,
	// but its value might change to fit semVer standards: https://semver.org/
	raw    string
	semVer *semanticVersion
}

// Gemfile.lock versions may have "{cpu}-{os}" or "{cpu}-{os}-{version}"
// after the semvVer, for example, 12.2.1-alpha-x86_64-darwin-8, where `2.2.1-alpha`
// is a valid and comparable semVer, and `x86_64-darwin-8` is not a semVer due to
// the underscore. Also, we can't sort based on arch and OS in a way that make sense
// for versions. SemVer is a characteristic of the code, not which arch OS it runs on.
//
// Bunlder's code: https://github.com/rubygems/rubygems/blob/2070231bf0c7c4654bbc2e4c08882bf414840360/bundler/spec/install/gemfile/platform_spec.rb offers more info on possible architecture values, for example `mswin32` may appead without arch.
//
// Spec for pre-release info: https://github.com/rubygems/rubygems/blob/2070231bf0c7c4654bbc2e4c08882bf414840360/bundler/spec/install/gemfile/path_spec.rb#L186
//
// CPU/arch is the most structured value present in gemfile.lock versions, we use it
// to split the version info in half, the first half has semVer, and
// the second half has arch and OS which we ignore.
// When there is no arch we split the version string with: {java, delvik, mswin32}
func extractSemVer(raw string) string {
	platforms := []string{"x86", "x86_64", "universal", "arm", "java", "dalvik", "x64", "powerpc", "sparc", "mswin32"}
	dash := "-"
	for _, p := range platforms {
		vals := strings.SplitN(raw, dash+p, 2)
		if len(vals) == 2 {
			return vals[0]
		}
	}

	return raw
}

func newGemfileVersion(raw string) (*gemfileVersion, error) {
	cleaned := extractSemVer(raw)
	semVer, err := newSemanticVersion(cleaned)
	if err != nil {
		return nil, fmt.Errorf("unable to crate gemfile version obj: %w", err)
	}
	return &gemfileVersion{
		raw:    raw,
		semVer: semVer,
	}, nil
}

func (g *gemfileVersion) Compare(other *Version) (int, error) {
	if other.Format != GemfileFormat && other.Format != SemanticFormat {
		return -1, fmt.Errorf("unable to compare Gemfile version to given format: %s", other.Format)
	}
	if other.rich.gemfileVer == nil || other.rich.gemfileVer.semVer == nil {
		if other.rich.semVer != nil {
			return other.rich.semVer.verObj.Compare(g.semVer.verObj), nil
		}
		return -1, fmt.Errorf("given empty gemfileVersion object")
	}

	return other.rich.gemfileVer.semVer.verObj.Compare(g.semVer.verObj), nil
}
