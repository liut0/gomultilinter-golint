package main

import (
	"context"

	"github.com/liut0/gomultilinter/api"
	"golang.org/x/lint"
)

type GoLint struct{}

var LinterFactory api.LinterFactory = &GoLint{}

func (l *GoLint) NewLinterConfig() api.LinterConfig {
	return l
}

func (l *GoLint) NewLinter() (api.Linter, error) {
	return l, nil
}

func (*GoLint) Name() string {
	return "golint"
}

func (l *GoLint) LintPackage(ctx context.Context, pkg *api.Package, reporter api.IssueReporter) error {
	for _, p := range lint.RunGoLinter(pkg) {
		reporter.Report(&api.Issue{
			Position: p.Position,
			Severity: api.SeverityWarning,
			Category: p.Category,
			Message:  p.Text,
		})
	}

	return nil
}
