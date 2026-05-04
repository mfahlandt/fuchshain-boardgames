---
title: "Kontakt & Anfahrt"
description: "Anfahrt zum Spieleabend in Schirnding im Fichtelgebirge – gut erreichbar aus Marktredwitz, Selb, Wunsiedel, Waldsassen und Tirschenreuth"
---

## Anfahrt & Kontakt

<div style="display: grid; grid-template-columns: 1fr 1fr; gap: 1.5rem; margin-bottom: 1.5rem; align-items: start;">
<div>

### Veranstaltungsort

**Fuchshain**  
Seedorf 1  
95706 Schirnding

</div>
<div>

<img src="/images/PXL_20250802_192815240.jpg" alt="Spieleabend im Fuchshain" style="width: 100%; height: 250px; object-fit: cover; border-radius: 12px;" />

</div>
</div>

<style>
@media (max-width: 768px) {
  div[style*="grid-template-columns: 1fr 1fr"] {
    grid-template-columns: 1fr !important;
  }
}
</style>

---

## Karte

> Hinweis: Die Karte wird erst nach Klick geladen. Dabei werden externe Inhalte von **unpkg.com** (Leaflet) und **OpenStreetMap** (Kartenkacheln) nachgeladen.

<button id="load-map" type="button" style="padding: .6rem 1rem; border-radius: 10px; border: 1px solid var(--border); background: var(--theme); cursor: pointer;">
  Karte laden
</button>

<div id="map" style="height: 400px; border-radius: 12px; margin: 1rem 0; display:none;"></div>

<script>
(function () {
  var btn = document.getElementById('load-map');
  var mapDiv = document.getElementById('map');

  function loadCss(href) {
    return new Promise(function (resolve, reject) {
      var link = document.createElement('link');
      link.rel = 'stylesheet';
      link.href = href;
      link.onload = resolve;
      link.onerror = reject;
      document.head.appendChild(link);
    });
  }

  function loadScript(src) {
    return new Promise(function (resolve, reject) {
      var script = document.createElement('script');
      script.src = src;
      script.async = true;
      script.onload = resolve;
      script.onerror = reject;
      document.head.appendChild(script);
    });
  }

  function initMap() {
    // Koordinaten für Seedorf 1, 95706 Schirnding
    var map = L.map('map').setView([50.04682589639413, 12.24626267772196], 15);

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      maxZoom: 19,
      attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> Mitwirkende'
    }).addTo(map);

    var marker = L.marker([50.04682589639413, 12.24626267772196]).addTo(map);
    marker.bindPopup("<b>Spieleabend Fuchshain</b><br>Seedorf 1, 95706 Schirnding");
  }

  btn.addEventListener('click', function () {
    btn.disabled = true;
    btn.textContent = 'Lade Karte…';

    Promise.all([
      loadCss('https://unpkg.com/leaflet@1.9.4/dist/leaflet.css'),
      loadScript('https://unpkg.com/leaflet@1.9.4/dist/leaflet.js')
    ]).then(function () {
      mapDiv.style.display = 'block';
      initMap();
      btn.style.display = 'none';
    }).catch(function () {
      btn.disabled = false;
      btn.textContent = 'Karte laden';
      alert('Die Karte konnte nicht geladen werden. Bitte später erneut versuchen.');
    });
  });
})();
</script>

---

## 📧 Anmeldung & Kontakt

**Anmeldung erforderlich!** Da nur ca. 18 Plätze verfügbar sind, bitten wir um vorherige Anmeldung.

**So kannst du dich anmelden:**
- **E-Mail:** [spieleabend@fuchshain.farm](mailto:spieleabend@fuchshain.farm)
- **Persönlich:** Nikki oder Mario direkt ansprechen
- **Facebook-Gruppe:** <a href="https://www.facebook.com/groups/fuchshainspieleabend" target="_blank" rel="noopener">Fuchshain Spieleabend</a>

Wir antworten in der Regel innerhalb von 1-2 Tagen.

---

## Anfahrt

### Mit dem Auto
- Von Marktredwitz: B303 Richtung Schirnding (ca. 15 Min)
- Von Selb: über B15 und B303 (ca. 20 Min)
- Von Wunsiedel: über B303 (ca. 20 Min)
- Von Waldsassen: über Mitterteich/B299 (ca. 15 Min)
- Von Tirschenreuth: über Mitterteich (ca. 25 Min)
- Von Eger/Cheb (CZ): Grenzübergang Schirnding (ca. 5 Min)

### Parken
**Kostenlose Parkplätze** hinterm Hof vorhanden – bitte die geschotterten Flächen benutzen.

### Mit öffentlichen Verkehrsmitteln
- Bahnhof Schirnding ca. 2 km entfernt
- Verbindungen über Marktredwitz

---

## Social Media

Folge uns für aktuelle Infos und Spielefotos:

- **Facebook-Gruppe:** <a href="https://www.facebook.com/groups/fuchshainspieleabend" target="_blank" rel="noopener">Fuchshain Spieleabend</a>

---

## Häufige Fragen

**Muss ich mich anmelden?**  
Ja, da nur ca. 18 Plätze verfügbar sind, bitten wir um Anmeldung per E-Mail, persönlich bei Nikki/Mario oder über die Facebook-Gruppe.

**Kostet der Eintritt etwas?**  
Nein, die Teilnahme ist kostenlos.

**Kann ich eigene Spiele mitbringen?**  
Ja, sehr gerne! Wir freuen uns über neue Spiele. Bitte davor bei uns anmelden – du bist dann für die Spielrunde "Gamemaster" und erklärst den Mitspielenden das Spiel.

**Gibt es Verpflegung?**  
Getränke werden vor Ort zum Selbstkostenpreis angeboten. Snacks zum Teilen sind willkommen.

**Bin ich als Anfänger willkommen?**  
Auf jeden Fall! Wir erklären gerne Regeln und haben auch einfache Spiele dabei.

<script type="application/ld+json">
{
  "@context": "https://schema.org",
  "@type": "FAQPage",
  "mainEntity": [
    {
      "@type": "Question",
      "name": "Muss ich mich zum Spieleabend anmelden?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "Ja, da nur ca. 18 Plätze verfügbar sind, bitten wir um Anmeldung per E-Mail an spieleabend@fuchshain.farm, persönlich bei Nikki/Mario oder über die Facebook-Gruppe."
      }
    },
    {
      "@type": "Question",
      "name": "Kostet der Spieleabend Eintritt?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "Nein, die Teilnahme am Spieleabend im Fuchshain ist kostenlos."
      }
    },
    {
      "@type": "Question",
      "name": "Kann ich eigene Brettspiele mitbringen?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "Ja, sehr gerne! Wir freuen uns über neue Spiele. Bitte davor bei uns anmelden – du bist dann für die Spielrunde Gamemaster und erklärst den Mitspielenden das Spiel."
      }
    },
    {
      "@type": "Question",
      "name": "Wo findet der Spieleabend statt?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "Der Spieleabend findet in der Taverne zum Fuchshain statt: Seedorf 1, 95706 Schirnding im Fichtelgebirge. Gut erreichbar aus Marktredwitz (15 Min), Selb (20 Min), Wunsiedel (20 Min), Waldsassen (15 Min) und Tirschenreuth (25 Min)."
      }
    },
    {
      "@type": "Question",
      "name": "Bin ich als Anfänger beim Spieleabend willkommen?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "Auf jeden Fall! Wir erklären gerne Regeln und haben auch einfache Spiele dabei. Ob Einsteiger oder Profi – bei uns ist jeder willkommen."
      }
    }
  ]
}
</script>

