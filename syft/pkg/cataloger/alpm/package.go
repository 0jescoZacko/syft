package alpm

import (
	"strings"

	"github.com/anchore/packageurl-go"
	"github.com/anchore/syft/syft/linux"
	"github.com/anchore/syft/syft/pkg"
	"github.com/anchore/syft/syft/source"
)

func newPackage(m *parsedData, release *linux.Release, dbLocation source.Location) pkg.Package {
	licenseCandidates := strings.Split(m.Licenses, "\n")

	p := pkg.Package{
		Name:         m.Package,
		Version:      m.Version,
		Locations:    source.NewLocationSet(dbLocation),
		Licenses:     pkg.NewLicenseSet(pkg.NewLicensesFromLocation(dbLocation.WithoutAnnotations(), licenseCandidates...)...),
		Type:         pkg.AlpmPkg,
		PURL:         packageURL(m, release),
		MetadataType: pkg.AlpmMetadataType,
		Metadata:     m.AlpmMetadata,
	}
	p.SetID()

	return p
}

func packageURL(m *parsedData, distro *linux.Release) string {
	if distro == nil || distro.ID != "arch" {
		// note: there is no namespace variation (like with debian ID_LIKE for ubuntu ID, for example)
		return ""
	}

	qualifiers := map[string]string{
		pkg.PURLQualifierArch: m.Architecture,
	}

	if m.BasePackage != "" {
		qualifiers[pkg.PURLQualifierUpstream] = m.BasePackage
	}

	return packageurl.NewPackageURL(
		"alpm", // `alpm` for Arch Linux and other users of the libalpm/pacman package manager. (see https://github.com/package-url/purl-spec/pull/164)
		distro.ID,
		m.Package,
		m.Version,
		pkg.PURLQualifiers(
			qualifiers,
			distro,
		),
		"",
	).ToString()
}
