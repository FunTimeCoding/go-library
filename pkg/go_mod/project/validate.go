package project

import "github.com/coreos/go-semver/semver"

func (p *Project) Validate() {
	defer func() {
		if r := recover(); r != nil {
			p.concern = append(p.concern, PanicOccurred)
		}
	}()

	versionSemantic := semver.New(p.Version)
	runtimeSemantic := semver.New(p.runtimeVersion)

	if versionSemantic.LessThan(*runtimeSemantic) {
		p.concern = append(p.concern, RuntimeOld)
	}

	if runtimeSemantic.LessThan(*versionSemantic) {
		p.concern = append(p.concern, RuntimeNewer)
	}
}
