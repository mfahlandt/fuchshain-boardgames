// CSV zu spiele.yaml Konverter
//
// Verwendung:
//   1. Google Sheets -> Datei -> Herunterladen -> CSV
//   2. go run csv-to-spiele-yaml.go input.csv > ../data/spiele.yaml
//
// Erwartete CSV-Spalten (aus Google Sheets):
//   BGG ID, Spielname, Kategorie, Kurzbeschreibung (Pitch), Komplexität (Sterne), Spielerzahl, Dauer (Min)

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Verwendung: go run csv-to-spiele-yaml.go <input.csv>")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Erwartete Spalten im CSV (Header-Zeile erforderlich):")
		fmt.Fprintln(os.Stderr, "  BGG ID, Spielname, Kategorie, Kurzbeschreibung (Pitch),")
		fmt.Fprintln(os.Stderr, "  Komplexität (Sterne), Spielerzahl, Dauer (Min)")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fehler beim Oeffnen: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Lese die ersten Bytes um UTF-8 BOM zu erkennen und ueberspringen
	bom := make([]byte, 3)
	n, _ := file.Read(bom)
	if n >= 3 && bom[0] == 0xEF && bom[1] == 0xBB && bom[2] == 0xBF {
		// UTF-8 BOM erkannt - bleibe an Position 3
	} else {
		// Kein BOM - zurueck zum Anfang
		file.Seek(0, 0)
	}

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fehler beim Lesen: %v\n", err)
		os.Exit(1)
	}

	if len(records) < 2 {
		fmt.Fprintln(os.Stderr, "CSV muss mindestens Header + 1 Datenzeile haben")
		os.Exit(1)
	}

	// Header analysieren - normalisiere Spaltennamen
	header := records[0]
	colIndex := make(map[string]int)
	for i, col := range header {
		normalized := strings.ToLower(strings.TrimSpace(col))
		colIndex[normalized] = i
	}

	// Mapping der Spaltennamen (deine Sheets-Namen -> interne Namen)
	columnMappings := map[string][]string{
		"name":         {"spielname", "name"},
		"bgg_id":       {"bgg id", "bgg_id", "bggid"},
		"kategorie":    {"kategorie", "category"},
		"beschreibung": {"kurzbeschreibung (pitch)", "kurzbeschreibung", "beschreibung", "pitch"},
		"komplexitaet": {"komplexität (sterne)", "komplexität", "komplexitaet", "sterne"},
		"spielerzahl":  {"spielerzahl", "spieler", "players"},
		"dauer":        {"dauer (min)", "dauer", "minuten", "duration"},
	}

	// Finde die tatsaechlichen Spaltenindizes
	findColumn := func(key string) int {
		for _, variant := range columnMappings[key] {
			if idx, ok := colIndex[variant]; ok {
				return idx
			}
		}
		return -1
	}

	// Pflichtfelder pruefen
	nameIdx := findColumn("name")
	if nameIdx == -1 {
		fmt.Fprintln(os.Stderr, "Pflichtfeld 'Spielname' fehlt im CSV-Header")
		fmt.Fprintf(os.Stderr, "Gefundene Spalten: %v\n", header)
		os.Exit(1)
	}

	// YAML Header
	fmt.Println("# Spielesammlung fuer Spieleabend Fuchshain")
	fmt.Println("#")
	fmt.Println("# Felder:")
	fmt.Println("#   name: Spielname")
	fmt.Println("#   komplexitaet: 1-5 (1=sehr einfach/Einsteiger, 5=sehr komplex)")
	fmt.Println("#   spielerzahl_min: Minimale Spielerzahl")
	fmt.Println("#   spielerzahl_max: Maximale Spielerzahl")
	fmt.Println("#   spielerzahl_beste: Beste Spielerzahl (kann auch Range sein, z.B. \"3-4\")")
	fmt.Println("#   dauer_min: Spieldauer in Minuten")
	fmt.Println("#   bgg_id: BoardGameGeek ID (optional, fuer Link zu BGG)")
	fmt.Println("#   kategorie: Spielkategorie (optional)")
	fmt.Println("#   beschreibung: Kurzbeschreibung (optional)")
	fmt.Println("")

	// Daten konvertieren
	bggIdx := findColumn("bgg_id")
	katIdx := findColumn("kategorie")
	beschrIdx := findColumn("beschreibung")
	komplexIdx := findColumn("komplexitaet")
	spielerzahlIdx := findColumn("spielerzahl")
	dauerIdx := findColumn("dauer")

	for _, row := range records[1:] {
		if len(row) == 0 || len(row) <= nameIdx || strings.TrimSpace(row[nameIdx]) == "" {
			continue // Leere Zeilen ueberspringen
		}

		getVal := func(idx int) string {
			if idx >= 0 && idx < len(row) {
				return strings.TrimSpace(row[idx])
			}
			return ""
		}

		name := getVal(nameIdx)
		if name == "" {
			continue
		}

		// Spielerzahl parsen (z.B. "2-4" oder "3-6" oder "2")
		spielerzahl := getVal(spielerzahlIdx)
		minPlayer, maxPlayer := parseSpielerzahl(spielerzahl)

		// Dauer parsen (z.B. "60" oder "45-90")
		dauer := getVal(dauerIdx)
		dauerMin := parseDauer(dauer)

		// Komplexitaet parsen
		komplex := getVal(komplexIdx)
		if komplex == "" {
			komplex = "2" // Default
		}

		fmt.Printf("- name: \"%s\"\n", escapeYaml(name))
		fmt.Printf("  komplexitaet: %s\n", komplex)
		fmt.Printf("  spielerzahl_min: %s\n", minPlayer)
		fmt.Printf("  spielerzahl_max: %s\n", maxPlayer)
		fmt.Printf("  dauer_min: %s\n", dauerMin)

		if bggID := getVal(bggIdx); bggID != "" && bggID != "0" {
			fmt.Printf("  bgg_id: %s\n", bggID)
		}

		if kategorie := getVal(katIdx); kategorie != "" {
			fmt.Printf("  kategorie: \"%s\"\n", escapeYaml(kategorie))
		}

		if beschreibung := getVal(beschrIdx); beschreibung != "" {
			fmt.Printf("  beschreibung: \"%s\"\n", escapeYaml(beschreibung))
		}

		fmt.Println("")
	}
}

// parseSpielerzahl parst "2-4" zu min="2", max="4" oder "3" zu min="3", max="3"
func parseSpielerzahl(s string) (string, string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "1", "4" // Default
	}

	// Regex fuer "X-Y" Pattern
	re := regexp.MustCompile(`(\d+)\s*[-–]\s*(\d+)`)
	if matches := re.FindStringSubmatch(s); len(matches) == 3 {
		return matches[1], matches[2]
	}

	// Nur eine Zahl
	re = regexp.MustCompile(`(\d+)`)
	if matches := re.FindStringSubmatch(s); len(matches) == 2 {
		return matches[1], matches[1]
	}

	return "1", "4"
}

// parseDauer parst "45-90" zu "45" oder "60" zu "60"
func parseDauer(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return "30" // Default
	}

	// Nimm den ersten Wert bei Range
	re := regexp.MustCompile(`(\d+)`)
	if matches := re.FindStringSubmatch(s); len(matches) >= 2 {
		return matches[1]
	}

	return "30"
}

func escapeYaml(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	return s
}
