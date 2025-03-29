{{/*
Return the name of the chart.
*/}}
{{- define "service-b.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end }}

{{/*
Return the fully qualified name.
*/}}
{{- define "service-b.fullname" -}}
{{- if .Values.fullnameOverride -}}
  {{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
  {{- printf "%s-%s" .Release.Name (include "service-b.name" .) | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end }}
