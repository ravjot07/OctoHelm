{{/*
Return the name of the chart.
*/}}
{{- define "service-a.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end }}

{{/*
Return the fully qualified name.
*/}}
{{- define "service-a.fullname" -}}
{{- if .Values.fullnameOverride -}}
  {{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
  {{- printf "%s-%s" .Release.Name (include "service-a.name" .) | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end }}
