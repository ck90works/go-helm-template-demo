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
