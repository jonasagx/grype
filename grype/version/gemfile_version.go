package version

import (
	"fmt"

	hashiVer "github.com/anchore/go-version"
)

type gemfileVersion struct {
	raw    string
	verObj *hashiVer.Version
}

func newGemfileVersion(raw string) (*gemfileVersion, error) {
	verObj, err := hashiVer.NewVersion(gemfileNormalizer.Replace(raw))
	if err != nil {
		return nil, fmt.Errorf("unable to crate gemfile version obj: %w", err)
	}
	return &gemfileVersion{
		raw:    raw,
		verObj: verObj,
	}, nil
}

func (g *gemfileVersion) Compare(other *Version) (int, error) {
	if other.Format != GemfileFormat {
		return -1, fmt.Errorf("unable to compare Gemfile.lock version to given format: %s", other.Format)
	}
	if other.rich.gemfileVer == nil {
		return -1, fmt.Errorf("given empty gemfileVersion object")
	}

	return other.rich.gemfileVer.verObj.Compare(g.verObj), nil
}
