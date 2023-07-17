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

Die zwei geschweiften Klammern `{{ }}` bzw. alternativ zwei eckige Klammern `[[ ]]` sind Begrenzungszeichen für Go Template Funktionen und Variablen, wobei jede `{{` mit `}}` bzw. `[[` mit `]]` enden muss.

_(Aus Gründen der Vereinfachung verwende ich standardmäßig die Notation mit den `{{ }}` für diesen Tutorial)_

Der Punkt in `{{ . }}` verweist auf den momentanen Wert während der Ausführung. Wenn man über eine **Slice** iteriert, wird die Ausgabe für jeden Wert in der Schleife separat eingesetzt.

Ein Bindestrich zu Beginn in `{{- }}` entfernt alle nachgestellten Leerzeichen für den vorangehenden Text und wenn ein Template-Befehl mit einem Bindestrich endet `{{ -}}`, dann werden alle führenden Lehrzeichen aus dem unmittelbar folgenden Text entfernt.

Es ist auch möglich ein Befehl mit einem Bindestrich zu beginnen und zu enden `{{- . -}}`, in diesem Fall werden alle nachgestellten Leerzeichen des vorangehenden Textes als auch alle führenden Lehrzeichen aus dem unmittelbar folgenden Text entfernt.

Kommentare werden in Form von `{{/* Dies ist ein Kommentar. Kann auch mehrzeilig sein */}}` geschrieben.

Pipelines werden mit einem Pipe-Symbol `{{ "Eingangswert" | mache_was_damit | mache_noch_was_damit | print }}` dargestellt und arbeiten von Links nach Rechts verschiedene Befehle ab wobei Links als Eingabe und Rechts als Ausgabe verwendet wird.
