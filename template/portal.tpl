{{template "base" .}}

{{ define "title" }}VM Portal{{ end }}

{{ define "style" }}
    table {
      font-family: arial, sans-serif;
      border-collapse: collapse;
      width: 100%;
    }
    td, th {
      border: 1px solid #dddddd;
      text-align: left;
      padding: 4px;
      vertical-align: top;
    }
    /*
    tr:nth-child(even) {
      background-color: #dddddd;
    }
    */
    th {
      background-color: #dddddd;
    }
    div.menu {
      font-weight: bolder;
      /* normal, bold, bolder, lighter, 100~900 */
      /* font-weight: 900; */
    }
    div.action {
      padding-left: 140px;
    }
{{ end }}
