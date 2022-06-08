package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGemfileVersionSemantic(t *testing.T) {
	tests := []testCase{
		// empty values
		{version: "2.3.1", constraint: "", satisfied: true},
		// typical cases
		{version: "1.2.0", constraint: ">1.0, <2.0", satisfied: true},
		{version: "1.2.0-x86-linux", constraint: ">1.0, <2.0", satisfied: true},
		{version: "1.2.0-x86", constraint: ">1.0, <2.0", satisfied: true},
		{version: "1.2.0-x86-linux", constraint: "= 1.2.0", satisfied: true},
		{version: "1.2.0-x86_64-linux", constraint: "= 1.2.0", satisfied: true},
		{version: "1.2.0-x86_64-linux", constraint: "< 1.2.1", satisfied: true},
		// https://semver.org/#spec-item-11
		{version: "1.2.0-alpha-x86-linux", constraint: "<1.2.0", satisfied: true},
		{version: "1.2.0-alpha-1-x86-linux", constraint: "<1.2.0", satisfied: true},
		{version: "1.2.0-alpha-1-x86-linux+meta", constraint: "<1.2.0", satisfied: true},
		{version: "1.2.0-alpha-1-x86-linux+meta", constraint: ">1.1.0", satisfied: true},
		{version: "1.2.0-alpha-1-arm-linux+meta", constraint: ">1.1.0", satisfied: true},
		{version: "1.0.0-alpha-a.b-c-somethinglong+build.1-aef.1-its-okay", constraint: "<1.0.0", satisfied: true},
		{version: "1.2.3----RC-SNAPSHOT.12.9.1--.12+788", constraint: "> 1.0.0", satisfied: true},
		{version: "1.2.3----RC-SNAPSHOT.12.9.1--.12+788-armv7-darwin", constraint: "< 1.2.3", satisfied: true},
	}

	for _, test := range tests {
		t.Run(test.tName(), func(t *testing.T) {
			constraint, err := newGemfileConstraint(test.constraint)
			assert.NoError(t, err, "unexpected error from newSemanticConstraint: %v", err)

			test.assertVersionConstraint(t, GemfileFormat, constraint)
		})
	}

}
