{{ define "base" }}
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>{{ template "title" . }}</title>
    <link rel="icon" href="data:,">
    <style>
    {{ template "style" . }}
    </style>
  </head>
  <body>

    {{ template "menu" . }}

    <hr/>

    {{ template "action" . }}
    {{ template "content" . }}

    <hr/>

    {{/* template "footer" . */}}

  </body>
</html>
{{ end }}
