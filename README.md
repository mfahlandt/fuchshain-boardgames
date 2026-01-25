# 🎲 Spieleabend Fuchshain

Website für den Spieleabend in Fuchshain – Dein Treffpunkt für Brettspiele, Kartenspiele und gesellige Abende.

🌐 **Live:** [spieleabend.fuchshain.farm](https://spieleabend.fuchshain.farm/)

## Features

- 📋 **Spielesammlung** – Alle verfügbaren Spiele mit Filtern und Suche
- 🖨️ **Druckversion** – Katalog aller Spiele zum Ausdrucken
- 📅 **Termine** – Übersicht der kommenden Spieleabende
- 🗳️ **Voting** – Abstimmung für Spiele via Nuudel
- 🗺️ **Anfahrt** – Interaktive Karte mit OpenStreetMap

## Tech Stack

- **[Hugo](https://gohugo.io/)** – Static Site Generator
- **GitHub Actions** – Automatisches Deployment
- **GitHub Pages** – Hosting

## Lokale Entwicklung

### Voraussetzungen

- [Hugo Extended](https://gohugo.io/installation/) (v0.120+)
- Git

### Starten

```bash
# Repository klonen
git clone https://github.com/mfahlandt/fuchshain-boardgames.git
cd fuchshain-boardgames

# Entwicklungsserver starten
hugo server -D
```

Die Seite ist dann unter `http://localhost:1313/` erreichbar.

### Build

```bash
hugo --minify
```

Der Output liegt im `/public/` Ordner.

## Projektstruktur

```
├── content/           # Seiteninhalte (Markdown)
│   ├── spiele/        # Einzelne Spieleseiten
│   └── ...
├── data/
│   ├── spiele.yaml    # Spieledatenbank
│   └── termine.yaml   # Termine
├── layouts/
│   └── shortcodes/    # Hugo Shortcodes
├── static/
│   ├── images/        # Bilder
│   └── fonts/         # Schriften
├── themes/            # Hugo Themes
└── tools/             # Hilfs-Scripts
```

## Spiele hinzufügen

1. Neuen Eintrag in `data/spiele.yaml` anlegen:

```yaml
- name: "Spielname"
  slug: "spielname"
  komplexitaet: 3          # 1-5
  spielerzahl_min: 2
  spielerzahl_max: 4
  dauer_min: 60
  bgg_id: 123456           # BoardGameGeek ID
  kategorie: "Strategie"
  bild: "/images/spiele/spielname.jpg"  # Optional
  beschreibung: "Kurze Beschreibung"
  langbeschreibung: |
    Ausführliche Beschreibung
    über mehrere Zeilen.
```

2. Optional: Bild in `static/images/spiele/` ablegen

3. Spieleseite generieren:
```bash
go run tools/generate-spiele-pages.go
```

## Deployment

Das Deployment erfolgt automatisch via GitHub Actions bei Push auf `main`.

### Manuelles Deployment

Im GitHub Repository unter **Actions** → **Deploy Hugo Site** → **Run workflow**

## Lizenz

MIT License – siehe [LICENSE](LICENSE)

---

🦊 Made with ❤️ in Fuchshain
