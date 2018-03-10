package lint

import (
	"github.com/liut0/gomultilinter/api"
)

func RunGoLinter(apiPkg *api.Package) []Problem {

	linterPkg := &pkg{
		fset:  apiPkg.FSet,
		files: map[string]*file{},

		typesPkg:  apiPkg.PkgInfo.Pkg,
		typesInfo: &apiPkg.PkgInfo.Info,

		sortable: map[string]bool{},
		main:     false,

		problems: []Problem{},
	}

	for _, f := range apiPkg.PkgInfo.Files {
		pos := apiPkg.FSet.Position(f.Pos())
		file := &file{
			fset:     apiPkg.FSet,
			f:        f,
			filename: pos.Filename,
			pkg:      linterPkg,
			src:      nil,
		}
		file.lint()
	}

	return linterPkg.problems
}
