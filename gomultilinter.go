package lint

import (
	"github.com/liut0/gomultilinter/api"
)

func RunGoLinter(apiFile *api.File) []Problem {
	linterPkg := &pkg{
		fset:  apiFile.FSet,
		files: map[string]*file{},

		typesPkg:  apiFile.PkgInfo.Pkg,
		typesInfo: &apiFile.PkgInfo.Info,

		sortable: map[string]bool{},
		main:     false,

		problems: []Problem{},
	}

	pos := apiFile.FSet.Position(apiFile.ASTFile.Pos())
	file := &file{
		fset:     apiFile.FSet,
		f:        apiFile.ASTFile,
		filename: pos.Filename,
		pkg:      linterPkg,
		src:      nil,
	}
	file.lint()

	return linterPkg.problems
}
