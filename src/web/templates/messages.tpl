{{ define "flashes" }}
  {{ if index . "errors" }}
    <div class="errors" style="background: darksalmon; padding: 10px;">
        {{ range index . "errors" }}
        <p>{{ . }}</p>
        {{ end }}
    </div>
  {{ end }}
  {{ if index . "flashes" }}
    <ul class="flashes">
        {{ range index . "flashes" }}
        <li>{{ . }}
        {{ end }}
    </ul>
  {{ end }}

{{ end }}