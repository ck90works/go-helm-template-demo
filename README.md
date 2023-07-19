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

Der Punkt in `{{ . }}` verweist auf den momentanen Wert während der Ausführung. Wenn man über eine **Slice** mithilfe der `range`-Funktion iteriert, wird die Ausgabe für jeden Wert in der Schleife separat eingesetzt und das Template für jeden einzelnen Wert nacheinander ausgegeben.

Ein Bindestrich zu Beginn `{{- }}` entfernt alle Leerzeichen und leere Zeilen unmittelbar **vor** einem `{{ }}` Block, bzw. wenn ein Block mit einem Bindestrich endet `{{ -}}`, dann werden alle Leerzeichen und leere Zeilen unmittelbar **nach** dem `{{ }}` Block entfernt.

Wichtig: Leerzeichen und neue Zeilen der eingesetzten Daten innerhalb des `{{ }}` Blocks bleiben davon unberührt und werden weiterhin so übernommen wie sie eingehen.

Es ist auch möglich ein Befehl mit einem Bindestrich zu beginnen und zu enden `{{- . -}}`, damit werden alle Leerzeichen und leere Zeilen unmittelbar vor und unmittelbar nach den geschweiften Klammern entfernt.

Beispiel: siehe [Bindestriche in Go Template](/beispiele/Bindestriche.md)

Funktionen in Templates werden innerhalb den Begrenzungen unmittelbar vor dem Datensatz, für den die Funktion gilt, aufgerufen, dabei ist zu beachten, dass die Anzahl der Argumente einer Funktion abhängig von seiner Programmierung abhängt. So ist `trim` eine gewöhnliche Funktion mit einem Argument, der führende und nachfolgende Leerzeichen eines Datensatzes entfernt.

Beispiel: siehe [Funktionen in Go Template](/beispiele/Funktionen.md)

Kommentare werden in Form von `{{/* Dies ist ein Kommentar. Kann auch mehrzeilig sein */}}` geschrieben.

Pipelines werden mit einem Pipe-Symbol `{{ "Eingangswert" | mache_was_damit | mache_noch_was_damit | print }}` dargestellt und arbeiten von Links nach Rechts verschiedene Befehle ab wobei Links als Eingabe und Rechts als Ausgabe verwendet wird.

Einfache runde Klammern `(` und `)` dienen dazu, bestimmte Funktionen zu Gliedern und zu verschachteln. 
<br>So würde `{{ print (addiere_folgende_strings "abc" "xyz") }}`
dazu führen, dass zuerst die Texte "abc" und "xyz" zu einem "abcxyz" addiert werden und anschließend als "abcxyz" ausgegeben werden.
