package ruby

import (
	"github.com/anchore/packageurl-go"
	"github.com/anchore/syft/syft/pkg"
	"github.com/anchore/syft/syft/source"
)

func newGemfileLockPackage(name, version string, locations ...source.Location) pkg.Package {
	p := pkg.Package{
		Name:      name,
		Version:   version,
		PURL:      packageURL(name, version),
		Locations: source.NewLocationSet(locations...),
		Language:  pkg.Ruby,
		Type:      pkg.GemPkg,
	}

	p.SetID()

	return p
}

func newGemspecPackage(m gemData, gemSpecLocation source.Location) pkg.Package {
	p := pkg.Package{
		Name:         m.Name,
		Version:      m.Version,
		Locations:    source.NewLocationSet(gemSpecLocation.WithAnnotation(pkg.EvidenceAnnotationKey, pkg.PrimaryEvidenceAnnotation)),
		Licenses:     pkg.NewLicenseSet(pkg.NewLicensesFromLocation(gemSpecLocation, m.Licenses...)...),
		PURL:         packageURL(m.Name, m.Version),
		Language:     pkg.Ruby,
		Type:         pkg.GemPkg,
		MetadataType: pkg.GemMetadataType,
		Metadata:     m.GemMetadata,
	}

	p.SetID()

	return p
}

func packageURL(name, version string) string {
	var qualifiers packageurl.Qualifiers

	return packageurl.NewPackageURL(
		packageurl.TypeGem,
		"",
		name,
		version,
		qualifiers,
		"",
	).ToString()
}
