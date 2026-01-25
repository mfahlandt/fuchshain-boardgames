---
title: "Kontakt"
description: "So erreichst du uns und findest zum Spieleabend"
---

## Anfahrt & Kontakt

### Veranstaltungsort

**Fuchshain**  
Seedorf 1  
95706 Schirnding

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

## 📧 Kontakt aufnehmen

Du hast Fragen zum Spieleabend oder möchtest dich anmelden?

**E-Mail:** [spieleabend@example.de](mailto:spieleabend@example.de)

Wir antworten in der Regel innerhalb von 1-2 Tagen.

---

## Anfahrt

### Mit dem Auto
- Von Marktredwitz: B303 Richtung Schirnding (ca. 15 Min)
- Von Eger/Cheb (CZ): Grenzübergang Schirnding (ca. 5 Min)
- Parkplätze direkt am Haus vorhanden (kostenlos)

### Mit öffentlichen Verkehrsmitteln
- Bahnhof Schirnding ca. 2 km entfernt
- Verbindungen über Marktredwitz

---

## Social Media

Folge uns für aktuelle Infos und Spielefotos:

- Facebook: [Spieleabend Fuchshain](#)
- Instagram: [@spieleabend_fuchshain](#)

---

## Häufige Fragen

**Muss ich mich anmelden?**  
Nein, komm einfach vorbei! Eine Anmeldung ist nicht erforderlich.

**Kostet der Eintritt etwas?**  
Nein, die Teilnahme ist kostenlos.

**Kann ich eigene Spiele mitbringen?**  
Ja, sehr gerne! Wir freuen uns über neue Spiele.

**Gibt es Verpflegung?**  
Getränke werden angeboten. Snacks zum Teilen sind willkommen.

**Bin ich als Anfänger willkommen?**  
Auf jeden Fall! Wir erklären gerne Regeln und haben auch einfache Spiele dabei.
