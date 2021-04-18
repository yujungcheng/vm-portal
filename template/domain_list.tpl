{{ define "content" }}
    <div class="action">
      <a href="domain?action=list">[ List ]</a>
      <a href="domain?action=create">[ Create ]</a>
      <a href="domain?action=delete">[ Delete ]</a>
      <a href="domain?action=modify">[ Modify ]</a>
    </div>
    <hr/>

    <table>
      <tr>
        <th>Name</th>
        <th>ID</th>
        <th>UUID</th>
        <th>State</th>
        <th>Vcpu</th>
        <th>Memory</th>
        <!--<th>Max Memory KB</th>-->
        <th>Disks</th>
        <th>Interfaces</th>
        <!--<th>CpuTime</th>-->
      </tr>
      {{ range .Domains }}
      <tr>
        <td>{{ .Name }}</td>
        <td>{{ .ID }}</td>
        <td>{{ .UUID }}</td>
        <td>{{ .StateStr }}</td>
        <td>{{ .Vcpu }}</td>
        <td>{{ .MemoryStr }}</td>
        <!--<td>{{ .MaxMem }}</td>-->
        <td>
          {{ range $index, $element := .Disks }}
            {{ if $index }}<br/>{{end}}
            {{ $element }}
          {{ end }}
        </td>
        <td>
          {{ range $index, $element := .Interfaces }}
            {{ if $index }}<br/>{{end}}
            {{ $element }}
          {{ end }}
        </td>
        <!--<td>{{ .CpuTime }}</td>-->
      </tr>
      {{ end }}

    </table>
{{ end }}
