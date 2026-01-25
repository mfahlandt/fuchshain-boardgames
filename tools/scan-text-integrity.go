package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

type finding struct {
	path   string
	reason string
}

func main() {
	root := "."
	extOK := map[string]bool{
		".md":   true,
		".toml": true,
		".yaml": true,
		".yml":  true,
		".json": true,
		".html": true,
		".css":  true,
		".js":   true,
		".txt":  true,
	}

	var findings []finding

	_ = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			findings = append(findings, finding{path: path, reason: fmt.Sprintf("WALK_ERROR: %v", err)})
			return nil
		}
		if d.IsDir() {
			// Skip typical large/vendor dirs that may contain non-utf8 assets
			s := filepath.ToSlash(path)
			if strings.Contains(s, "/themes/PaperMod/assets/") || strings.Contains(s, "/public/") {
				return fs.SkipDir
			}
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		if !extOK[ext] {
			return nil
		}

		b, readErr := os.ReadFile(path)
		if readErr != nil {
			findings = append(findings, finding{path: path, reason: fmt.Sprintf("READ_ERROR: %v", readErr)})
			return nil
		}

		if bytes.IndexByte(b, 0x00) >= 0 {
			findings = append(findings, finding{path: path, reason: "NULLBYTE"})
		}

		// Very common troublemaker for tooling: UTF-8 BOM at start.
		if len(b) >= 3 && b[0] == 0xEF && b[1] == 0xBB && b[2] == 0xBF {
			findings = append(findings, finding{path: path, reason: "UTF8_BOM"})
		}

		if !utf8.Valid(b) {
			findings = append(findings, finding{path: path, reason: "NON_UTF8"})
		}

		return nil
	})

	if len(findings) == 0 {
		fmt.Println("OK: keine Nullbytes/BOM/NON_UTF8 in gescannten Textdateien (public & theme vendor assets übersprungen).")
		return
	}

	fmt.Println("Auffälligkeiten:")
	for _, f := range findings {
		fmt.Printf("- %s: %s\n", f.reason, f.path)
	}
	os.Exit(1)
}
