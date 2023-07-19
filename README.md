# Go Template

## Einführung

Das Go Template Tool ermöglicht den Entwicklern eine Vorlage für die Nutzer zu definieren, die ihre Eingaben in eine gewünschte Ausgabeform validieren.

Hierfür stellt Golang standardmäßig eine Bibliothek namens "text/template" bereit, mit der Text-Vorlagen verarbeitet werden können.

Zusätzlich wurde die Standard-Bibliothek um die "html/template" Bibliothek erweitert, die zu allen Funktionen der Text-Verarbeitung um für HTML-Vorlagen nützliche Funktionen austattet.

Gängige Beispiele sind z.B. HTML-formatierte Ausgaben, die mithilfe des html/template Standard-Bibliotheks, die Formatierung in HTML auch für den Entwickler stark vereinfacht und diese aus dem Code auslagert.

Konkretes Beispiel:

<li>Entwickler implementiert eine Funktion, die Eingaben aus der Kommandozeile in eine vordefinierte HTML-Vorlage einsetzt, und diese als HTML-Datei speichert.
<br><br>
<li>Nutzer verwendet diese Anwendung und gibt den Text als Argument in der Kommandozeile ein und erhält eine sauber definierte HTML-Ausgabe
<br><br>
Siehe Code-Beispiel html_generator

## Infos

Die zwei geschweiften Klammern `{{ }}` sind Begrenzungszeichen für Go Template Funktionen und Variablen, wobei jede `{{` mit `}}` enden muss.

Der Punkt in `{{ . }}` verweist auf den momentanen Wert während der Ausführung. Wenn man über eine **Slice** iteriert, wird die Ausgabe für jeden Wert in der Schleife separat eingesetzt und das Template für jeden einzelnen Wert nacheinander ausgegeben.

Ein Bindestrich zu Beginn `{{- }}` entfernt alle Leerzeichen und leere Zeilen unmittelbar **vor** einem `{{ }}` Block, bzw. wenn ein Block mit einem Bindestrich endet `{{ -}}`, dann werden alle Leerzeichen und leere Zeilen unmittelbar **nach** dem `{{ }}` Block entfernt.

Wichtig: Leerzeichen und neue Zeilen der eingesetzten Daten innerhalb des `{{ }}` Blocks bleiben davon unberührt und werden weiterhin so übernommen wie sie eingehen.

Es ist auch möglich ein Befehl mit einem Bindestrich zu beginnen und zu enden `{{- . -}}`, damit werden alle Leerzeichen und leere Zeilen unmittelbar vor und unmittelbar nach den geschweiften Klammern entfernt.

Beispiel:

```
# Sei . ein Text mit Leerzeichen wie
# "   Hallo Welt!  " und würde das Template wie folgt aussehen:
{{- range . }}
...
metadata:
  namespace: {{ . }}
...
{{- end }}
# Würde dies folgenden Output generieren:
...
metadata:
  name: aid_namespace_test
  namespace:    Hallo Welt!
...
# Würde aber der Platzhalter im Namespace
# wie folgt aussehen:
...
metadata:
  name: aid_namespace_test
  namespace: {{- . }}
...
# Wäre der Output zwar um Leerzeichen und
# leere Zeilen gekürzt, aber nicht für die
# Variable innerhalb der geschweiften Klammern:
...
metadata:
  name: aid_namespace_test
  namespace:   Hallo Welt!
...
```

Kommentare werden in Form von `{{/* Dies ist ein Kommentar. Kann auch mehrzeilig sein */}}` geschrieben.

Pipelines werden mit einem Pipe-Symbol `{{ "Eingangswert" | mache_was_damit | mache_noch_was_damit | print }}` dargestellt und arbeiten von Links nach Rechts verschiedene Befehle ab wobei Links als Eingabe und Rechts als Ausgabe verwendet wird.
