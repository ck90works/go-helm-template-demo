## Allgemeine Beispiele für die Verwendung von Funktionen in Go Templates

Zu beachten ist, dass hier gezeigte Funktionen hauptsächlich aus der sprig v3 Bibliothek entstammen, welche für Helm Templates bereits integriert ist und voll funktionsfähig ist.
  
-> [Der Link zur offiziellen Dokumentation aller Funktionen des sprig v3 Bibliotheks](https://masterminds.github.io/sprig/)

```yaml
# Eine gewöhnliche Funktion mit einem Argument ist die trim-Funktion.
{{ trim . }}
# Sei . ein Repräsentant für "  hallo-welt   ",
# so würde der trim-Befehl daraus ein "hallo-welt" machen
# Template Definition mit:
...
  namespace: {{ trim . }}
...
# Führt zu:
...
  namespace: hallo-welt
...

# Man kann auch Funktionen mit einem | symbol verketten.
# So würde z.B. die Template Definition von:
...
  namespace: {{ trim . | nindent 2 }}
...
# Den Datensatz . (welcher den Text "  hallo-welt   " repräsentiert)
# zuerst mit trim auf "hallo-welt" gekürzt und anschließend mit nindent 2
# zuerst eine neue Zeile beginnt (daher `n`indent anstelle indent) und anschließend
# um zwei Leerzeichen nach rechts verschiebt.
# (Analog existiert auch die indent-Funktion, die keine neue Zeile beginnt)
# Somit sähe das Ergebnis zu "  hallo-welt   " als Datensatz wie folgt aus:
...
  namespace:
    hallo-welt
...
```

Zu den Funktionen wie `get` und `default` gibt es hierzu ein vereinfachtes Beispiel:
```yaml
# Sowohl get als auch default sind Funktionen, die zwei Argumente benötigen, also
# get .X .Y bzw. default .X .Y
#
# Im folgenden wird eine Variable namens $datensatz neu definiert und dieser nimmt
# den Wert aus einem Schlüssel von einer Map an:
{{ $datensatz := get aus_einer_map "der_schluessel" }}
# Würde die Map namens aus_einer_map wie folgt aussehen
{
  "der_schluessel": "ein_wert"
  "noch_ein_schluessel": "noch_ein_wert"
}
# Dann würde die Variable $datensatz, denn Wert "ein_wert" annehmen

# Die default Funktion definiert eine Ausweichwert, falls der angesprochene Datensatz
# leer ist bzw. nicht existiert.
# Im folgenden wird der bereits oben definierte $datensatz mit dem Wert "ein_wert" als
# Ausweichwert definiert, sollte der aktuelle Datensatz . leer sein bzw. nicht existieren:
{{ $neuer_wert := default $datensatz . }}
# Angenommen . ist leer, dann ist $neuer_wert = "ein_wert"
# Angenommen . ist nicht leer und hält den Wert "42", dann ist $neuer_wert = "42"
```
Reminder: Der Punkt in `{{ . }}` verweist auf den momentanen Wert während der Ausführung.

## Spezialfälle, die nur in Helm Templates funktionieren

Die `include`-Funktion ist eine von Helm entwickelte Funktion um ein für Helm Templating relevantes Problem zu lösen.
Im folgenden ein Beispiel zu einer solchen `include`-Verwendung:
```yaml
# Wir definieren zwei Templates, bei der wir im zweiten Template temporär
# die Funktionen des ersten Templates für einen bestimmten Datensatz verwenden werden:
{{/*
  Im formatierer-Template werden alle Eingangswerte zuerst danach geprüft
  ob dieser einen Unterstrich im Text beinhaltet, wenn ja, wird dieser mit
  einem Bindestrich ersetzt.
  Als nächstes wird geprüft, ob der Text Großbuchstaben beinhaltet, wenn ja,
  werden alle Großbuchstaben zu Kleinbuchstaben konviertiert. */}}
{{ define "formatierer" }}
{{ if contains "_" . }}
  {{ regexReplaceAll "_" . "-" }}
{{ end }}
{{ if regexMatch "^.+[A-Z]+" . }}
  {{ lower . }}
{{ end }}
{{ end }}
{{/*
  namespace_konvertierer nutzt die include-Funktion für den aktuellen Datensatz, bei der
  dieser zuerst den formatierer-Template durchläuft, dabei eventuell Änderungen vornimmt
  und anschließend in $variable speichert. */}}
{{ define "namespace_konvertierer" }}
{{ $variable := include "formatierer" . }}
namespace: {{ $variable }}
{{ end }}
...
# Wir können jetzt diese definierten Templates in einem weiteren Go Template - der mit der
# selben Go-Anwendung interagiert - nutzen und müssen diese Definitionen nicht ständig neu
# runterschreiben. Die Nutzung eines anderen Templates in einem Template wird mit der template-Funktion ermöglicht:
...
metadata:
  name: aid-test
  {{ template "namespace_konvertierer" . }}
...
# Sei . Repräsentant eines Textes "Aid_Namespace", würde der obige Template folgendes generieren:
...
metadata:
  name: aid-test
  namespace: aid-namespace
...
```
