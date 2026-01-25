// Fix Umlaute - ersetzt kaputte UTF-8 Sequenzen durch echte Umlaute
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Verwendung: go run fix-umlaute.go <datei>")
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fehler: %v\n", err)
		os.Exit(1)
	}

	content := string(data)

	// Kaputte UTF-8 Sequenzen (doppelt-encodierte oder falsche Codepage)
	replacements := map[string]string{
		// Typische doppelt-encodierte UTF-8 Zeichen
		"\xe2\x94\x9c\xc6\x92": "ß", // ├ƒ -> ß
		"\xe2\x94\x9c\xc2\xbc": "ü", // ├╝ -> ü (guess)
		"\xc3\xa4":             "ä", // ä
		"\xc3\xb6":             "ö", // ö
		"\xc3\xbc":             "ü", // ü
		"\xc3\x84":             "Ä", // Ä
		"\xc3\x96":             "Ö", // Ö
		"\xc3\x9c":             "Ü", // Ü
		"\xc3\x9f":             "ß", // ß
		"\xe2\x80\x93":         "–", // en-dash
		"\xe2\x80\x99":         "'", // right single quote
	}

	for old, newVal := range replacements {
		content = strings.ReplaceAll(content, old, newVal)
	}

	// Zusätzlich ae/oe/ue in deutschen Wörtern ersetzen
	wordReplacements := map[string]string{
		"fuer":           "für",
		"Fuer":           "Für",
		"muessen":        "müssen",
		"muesst":         "müsst",
		"koennt":         "könnt",
		"koennen":        "können",
		"moeglich":       "möglich",
		"grosse":         "große",
		"groessere":      "größere",
		"waehlt":         "wählt",
		"waehrend":       "während",
		"spaeter":        "später",
		"zurueck":        "zurück",
		"ueberarbeitete": "überarbeitete",
		"ueberwinden":    "überwinden",
		"drueber":        "drüber",
		"Schluessel":     "Schlüssel",
		"Zuege":          "Züge",
		"fuehrt":         "führt",
		"Wuerfel":        "Würfel",
		"Voelker":        "Völker",
		"voellig":        "völlig",
		"hoeren":         "hören",
		"gehoert":        "gehört",
		"Hoehlen":        "Höhlen",
		"schoen":         "schön",
		"wunderschoen":   "wunderschön",
		"Koenig":         "König",
		"loesen":         "lösen",
		"boese":          "böse",
		"Goetter":        "Götter",
		"aegyptische":    "ägyptische",
		"staerker":       "stärker",
		"Aera":           "Ära",
		"Lueckentexte":   "Lückentexte",
		"fuellen":        "füllen",
		"natuerlich":     "natürlich",
		"regelmaessig":   "regelmäßig",
		"aehnlich":       "ähnlich",
		"waechst":        "wächst",
		"Maerchen":       "Märchen",
		"erzaehlt":       "erzählt",
		"Saeulen":        "Säulen",
		"Traenke":        "Tränke",
		"Plaettchen":     "Plättchen",
		"Schaetze":       "Schätze",
		"Haendler":       "Händler",
		"kaempft":        "kämpft",
		"Faehigkeiten":   "Fähigkeiten",
		"erklaert":       "erklärt",
		"waere":          "wäre",
		"haette":         "hätte",
		"naechste":       "nächste",
		"naechsten":      "nächsten",
		"toedlich":       "tödlich",
		"toeten":         "töten",
		"wuetende":       "wütende",
		"suess":          "süß",
		"zugaenglich":    "zugänglich",
		"Laeden":         "Läden",
		"Gemueter":       "Gemüter",
		"Glueck":         "Glück",
		"glueckslaunig":  "glückslaunig",
		"aelteren":       "älteren",
		"Spass":          "Spaß",
		"verpruegeln":    "verprügeln",
		"Verraeter":      "Verräter",
		"Raetsel":        "Rätsel",
		"eigenstaendig":  "eigenständig",
		"Kriegsfuehrung": "Kriegsführung",
		"Unterstuetzung": "Unterstützung",
		"atmosphaerisch": "atmosphärisch",
		"zuechten":       "züchten",
		"Parkplaetze":    "Parkplätze",
		"oeffentlichen":  "öffentlichen",
		"ergaenzt":       "ergänzt",
	}

	for old, newVal := range wordReplacements {
		content = strings.ReplaceAll(content, old, newVal)
	}

	if err := os.WriteFile(os.Args[1], []byte(content), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Fehler beim Schreiben: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Umlaute korrigiert!")
}
