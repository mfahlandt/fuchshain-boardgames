// Generiert Content-Dateien fuer alle Spiele aus data/spiele.yaml
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Spiel struct {
	Name         string `yaml:"name"`
	Slug         string `yaml:"slug"`
	Kategorie    string `yaml:"kategorie"`
	Beschreibung string `yaml:"beschreibung"`
}

func main() {
	// Lese spiele.yaml
	data, err := os.ReadFile("../data/spiele.yaml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fehler beim Lesen: %v\n", err)
		os.Exit(1)
	}

	var spiele []Spiel
	if err := yaml.Unmarshal(data, &spiele); err != nil {
		fmt.Fprintf(os.Stderr, "Fehler beim Parsen: %v\n", err)
		os.Exit(1)
	}

	// Erstelle content/spiele/ Verzeichnis
	outDir := "../content/spiele"
	os.MkdirAll(outDir, 0755)

	count := 0
	for _, s := range spiele {
		if s.Slug == "" {
			continue
		}

		filename := filepath.Join(outDir, s.Slug+".md")
		content := fmt.Sprintf(`---
title: "%s"
description: "%s"
layout: "single"
spiel: "%s"
type: "spiele"
---
`, s.Name, s.Beschreibung, s.Slug)

		if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Fehler beim Schreiben von %s: %v\n", filename, err)
			continue
		}
		count++
		fmt.Printf("Erstellt: %s\n", filename)
	}

	fmt.Printf("\n%d Spiel-Seiten erstellt.\n", count)
}
