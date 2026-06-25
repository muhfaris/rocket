package generate

import (
	"fmt"
	"os"
	"strings"
)

// AppendToFile appends content to the end of a file, creating it if missing.
func AppendToFile(filepath, content string) error {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("open %s: %w", filepath, err)
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		return fmt.Errorf("write %s: %w", filepath, err)
	}
	return nil
}

// InsertBeforeLastClosingBrace reads a file and inserts content before the final `}`.
// SIMPLIFICATION: naive last-`}`-in-file search. If the file has a top-level closing
// brace that isn't what we want (e.g. a struct literal at end), this will misfire.
// Upgrade path: use go/ast with go/parser + go/printer.
func InsertBeforeLastClosingBrace(filepath, content string) error {
	raw, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("read %s: %w", filepath, err)
	}

	txt := string(raw)
	idx := strings.LastIndex(txt, "}")
	if idx == -1 {
		return fmt.Errorf("no closing brace found in %s", filepath)
	}

	// Insert content before the closing brace (preserve the brace on its own line)
	before := txt[:idx]
	after := txt[idx:]
	out := before + strings.TrimRight(content, "\n") + "\n" + after

	if err := os.WriteFile(filepath, []byte(out), 0644); err != nil {
		return fmt.Errorf("write %s: %w", filepath, err)
	}
	return nil
}

// InsertBeforePackageCloser inserts content before the closing brace of a specific
// top-level declaration identified by its leading text (e.g. "type Service interface {").
// SIMPLIFICATION: finds the first matching declaration header, then its matching `}`
// using brace counting. Works for simple generated code; may misfire on complex hand-edited code.
func InsertBeforePackageCloser(filepath, declarationHeader, content string) error {
	raw, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("read %s: %w", filepath, err)
	}

	txt := string(raw)
	startIdx := strings.Index(txt, declarationHeader)
	if startIdx == -1 {
		return fmt.Errorf("declaration %q not found in %s", declarationHeader, filepath)
	}

	// Find the matching closing brace by counting braces
	braceCount := 0
	inDecl := false
	insertIdx := -1

	for i := startIdx; i < len(txt); i++ {
		ch := txt[i]
		if ch == '{' {
			braceCount++
			inDecl = true
		} else if ch == '}' {
			braceCount--
			if inDecl && braceCount == 0 {
				insertIdx = i
				break
			}
		}
	}

	if insertIdx == -1 {
		return fmt.Errorf("could not find closing brace for %q in %s", declarationHeader, filepath)
	}

	before := txt[:insertIdx]
	after := txt[insertIdx:]
	out := before + strings.TrimRight(content, "\n") + "\n" + after

	if err := os.WriteFile(filepath, []byte(out), 0644); err != nil {
		return fmt.Errorf("write %s: %w", filepath, err)
	}
	return nil
}
