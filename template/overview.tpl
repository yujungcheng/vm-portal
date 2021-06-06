{{ define "menu" }}
    <div class="menu">
      <a href="overview">**Overview**</a> |
      <a href="cluster">Cluster</a> |
      <a href="domain">Domain</a> |
      <a href="image">Image</a> |
      <a href="volume">Volume</a> |
      <a href="network">Network</a> |
      <a href="template">Template</a> |
      <a href="system">System</a> |
      <a href="backup">Backup</a> |
      <a href="restore">Restore</a> |
      <a href="event">Event</a>
    </div>
{{ end }}

{{ define "action" }}
{{ end }}

{{ define "content" }}
    <!-- System Status -->
    <table>
      <tr>
        <td>
          <b>Socket Number:</b> {{ .CpuSocketsNum }}
        </td>
        <td>
          <b>Core(s) per socket:</b> {{ .CpuCorePerSocket }}
        </td>
        <td>
          <b>Thread(s) per core:</b> {{ .CpuThreadPerCore }}
        </td>
        <td>
          <b>Core Number:</b> {{ .CpuCoreNum }}
        </td>
        <td>
          <b>CPU Model:</b> {{ .CpuModel }}
        </td>
        <td>
          <b>CPU MHz:</b> {{ .CpuMHz }}
        </td>
        <td>
          <b>Numa Cell Number:</b> {{ .NumaCellNum }}
        </td>
        <td>
          <b>Memory Size:</b> {{ .MemorySize }}
        </td>

      </tr>
    </table>
    <br/>

    <!-- Resource Status -->
    <table>
      <thead>
        <tr>
          <th>Domains</th>
          <th>Storage Pools</th>
          <th>Networks</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td class="">
            {{ range .DomainNames }}
              {{ . }}<br/>
            {{ end }}
          </td>
          <td class="">
            {{ range .StoragePools }}
              {{ . }}</br>
            {{ end }}
          </td>
          <td class="">
            {{ range .Networks }}
              {{ . }}<br/>
            {{ end }}
          </td>
        </tr>
        <!--
        <tr>
          <td class="">Clusters</td>
          <td class="">Templates</td>
        </tr>
        -->
      </tbody>
    </table>
    <br/>

    <!-- Portal Status -->
    <table>
      <tr>
        <td>
          <b>Status:</b> OK
        </td>
        <td>
          <b>Version:</b> Developing
        </td>
        <td>
          <b>Uptime:</b> {{ .PortalUptime }}
        </td>
        <td>
          <b>PID:</b> {{ .PortalPID }}
        </td>
      </tr>
    </table>

{{ end }}

{{ define "footer" }}
    <!-- System Version -->
    <b>Hostname:</b> {{ .Hostname }}
    &nbsp;&nbsp;&nbsp;&nbsp;
    <b>Libvirt Version:</b> {{ .LibvirtVersion }}
    &nbsp;&nbsp;&nbsp;&nbsp;
    <b>Kernel Version:</b>
    &nbsp;&nbsp;&nbsp;&nbsp;
    <b>OS Distribution:</b>

{{ end }}
