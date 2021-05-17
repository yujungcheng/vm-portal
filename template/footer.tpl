{{ define "footer" }}
    <div class="footer">
    Checked at: {{ .CheckedAt }} | Total CPUs: | Total Memory: | Total Domains:
    </div>

    <table>
      <tr>
        <td>
          Hostname
        </td>
        <td>
          Kernel Version
        </td>
        <td>
          OS Distribution
        </td>
        <td>
          Libvirt Version
        </td>
      </tr>
    </table>

{{ end }}
