{{ define "messages" }}
  {{ if index . "errors" }}
    <div class="errors" style="background: darksalmon; padding: 10px;">
        {{ range index . "errors" }}
        <p>{{ . }}</p>
        {{ end }}
    </div>
  {{ end }}
  {{ if index . "messages" }}
    <ul class="flashes">
        {{ range index . "messages" }}
        <li>{{ . }}
        {{ end }}
    </ul>
  {{ end }}

{{ end }}