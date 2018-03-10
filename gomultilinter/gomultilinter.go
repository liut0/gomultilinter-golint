package main

import (
	"context"

	"github.com/liut0/gomultilinter/api"
	"golang.org/x/lint"
)

type GoLint struct {
	MinConfidence float64
}

var LinterFactory api.LinterFactory = &GoLint{}

func (l *GoLint) NewLinterConfig() api.LinterConfig {
	return &GoLint{
		MinConfidence: 0.8,
	}
}

func (l *GoLint) NewLinter() (api.Linter, error) {
	return l, nil
}

func (*GoLint) Name() string {
	return "golint"
}

func (l *GoLint) LintFile(ctx context.Context, file *api.File, reporter api.IssueReporter) error {
	for _, p := range lint.RunGoLinter(file) {
		if p.Confidence > l.MinConfidence {
			reporter.Report(&api.Issue{
				Position: p.Position,
				Severity: api.SeverityWarning,
				Category: p.Category,
				Message:  p.Text,
			})
		}
	}

	return nil
}
