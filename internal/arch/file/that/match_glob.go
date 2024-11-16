package that

import (
	"path/filepath"
	"strings"

	"github.com/omissis/goarkitect/internal/arch/file"
	"github.com/omissis/goarkitect/internal/arch/rule"
)

func MatchGlob(s string) *EndWithExpression {
	return &MatchGlobExpression{
		suffix: s,
	}
}

type MatchGlobExpression struct {
	suffix string

	errors []error
}

func (e *MatchGlobExpression) GetErrors() []error {
	return e.errors
}

func (e *MatchGlobExpression) Evaluate(rb rule.Builder) {
	frb, ok := rb.(*file.RuleBuilder)
	if !ok {
		e.errors = append(e.errors, file.ErrInvalidRuleBuilder)

		return
	}

	files := make([]string, 0)

	for _, f := range frb.GetFiles() {
		abs1, err := filepath.Abs(filepath.Join(e.basePath, e.glob))
		if err != nil {
			e.errors = append(e.errors, file.ErrInvalidRuleBuilder)

			return
		}

		abs2, err := filepath.Abs(filePath)
		if err != nil {
			e.errors = append(e.errors, file.ErrInvalidRuleBuilder)

			return
		}

		match, err := filepath.Match(abs1, abs2)
		if err != nil {
			rb.AddError(err)
		}

		if match {
			files = append(files, f)
		}
	}

	frb.SetFiles(files)

}

