{{/*
type.lookup supports the lookup of complex values in params
the -> `data` key part for list/slice is needed to make helm happy.
all data types which are not map or slice are treated as they are.

variable:
    - test
    - a: test
    - a: test
      b: test
variable:
  a: test
  b: test
*/}}
//  Beginn Template Definitionsblock genannt type.lookup
{{- define "type.lookup" -}}
// Beginn eines If Blocks
{{- if kindIs "map" . -}} // Wenn Datenstruktur eine Map ist, dann wird der aktuelle Datensatz (als . definiert) ohne weiteres Übernommen
  {{- . }} // Daten werden linksseitig ohne Leerzeichen übernommen
{{- else if kindIs "slice" . -}} // Wenn Datenstruktur eine Slice ist, dann wird vor dem aktuellen Datensatz der Text "data: " eingefügt und der Datensatz selbst eingelesen und mit einer neuen Zeile und zwei Leerzeichen versehen und in den Platzhalter eingesetzt
data: {{- . | nindent 2 }}
{{- else }} // Alle anderen Datenstrukturen werden so übernommen wie sie reinkommen
  {{- . }}
{{- end }}
// Ende des If Blocks
{{- end }}
// Ende Template Definitionsblock für type.lookup

{{/*
lookup function for environment or site or location or cluster specific variables
Everything below will always win against default

environment always wins against provider
location always wins against environment and provider
site always wins against location and environment and provider
siteLocation always wins against site and location and environment and provider
providerSiteLocation always wins against siteLocation and site and location and environment and provider
cluster always wins against all other definitions.
Use with caution.
*/}}
{{- define "lookup" -}} // Definiere eine neue Template namens lookup
// Das get in (get (default . .root) "variables") erhält aus der Map (default . .root) den Value zu dem Key "variables"
// Das default in (default . .root) prüft ob .root leer ist, wenn ja, nutze den Default-Wert der hier mit . (aktueller Datensatz) definiert ist
// Setze in $variables inhalt von "v1 ein oder den Wert in (dict)
{{- $root_vs_variables := default . .root }} // wenn .root leer oder nicht existent, dann nutze . 
{{- $variables := default (dict) (default (dict) (get (default (dict) (get (default . .root) "variables")) "v1")) -}}
{{- $params := (get (default . .root) "params") -}}
{{- $params_v1 := get $params "v1" -}} // erhält aus der Map $params den Wert aus dem Key "v1"
{{- $default_variable := default "" (get (default (dict) $variables.default) .variable) -}}
{{- $provider_variable := default "" (get (default (dict) (get (default (dict) $variables.provider) $params_v1.provider)) .variable) -}}
{{- $environment_variable := default "" (get (default (dict) (get (default (dict) $variables.environment) $params_v1.environment)) .variable) -}}
{{- $location_variable := default "" (get (default (dict) (get (default (dict) $variables.location) $params_v1.location)) .variable) -}}
{{- $site_variable := default "" (get (default (dict) (get (default (dict) $variables.site) $params_v1.site.id)) .variable) -}}
{{- $site_location_variable := default "" (get (default (dict) (get (default (dict) $variables.siteLocation) (print $params_v1.site.id "-" $params_v1.location ))) .variable) -}}
{{- $provider_site_location_variable := default "" (get (default (dict) (get (default (dict) $variables.providerSiteLocation) (print $params_v1.provider "-" $params_v1.site.id "-" $params_v1.location ))) .variable) -}}
{{- $cluster_variable := default "" (get (default (dict) $variables.cluster) .variable) -}}
{{- if $cluster_variable -}}
{{- include "type.lookup" $cluster_variable }} 
// füge die Variable $cluster_variable in das "type.lookup" template
{{- else if $provider_site_location_variable -}}
{{- include "type.lookup" $provider_site_location_variable }}
{{- else if $site_location_variable -}}
{{- include "type.lookup" $site_location_variable }}
{{- else if $site_variable -}}
{{- include "type.lookup" $site_variable }}
{{- else if $location_variable -}}
{{- include "type.lookup" $location_variable }}
{{- else if $environment_variable -}}
{{- include "type.lookup" $environment_variable }}
{{- else if $provider_variable -}}
{{- include "type.lookup" $provider_variable }}
{{- else if $default_variable -}}
{{- include "type.lookup" $default_variable }}
{{- end }}
{{- end }}
