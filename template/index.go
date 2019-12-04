package template

import "fmt"

var indexHTML = `
<!-- HTML for dev server -->
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Swagger UI</title>
  <link rel="stylesheet" type="text/css" href="./swagger-ui.css">
  <link rel="icon" type="image/png" href="./favicon-32x32.png" sizes="32x32"/>
  <link rel="icon" type="image/png" href="./favicon-16x16.png" sizes="16x16"/>
  <style>
    html {
      box-sizing: border-box;
      overflow: -moz-scrollbars-vertical;
      overflow-y: scroll;
    }

    *,
    *:before,
    *:after {
      box-sizing: inherit;
    }

    body {
      margin: 0;
      background: #fafafa;
    }
  </style>
</head>

<body>
<div id="swagger-ui"></div>

<script src="./swagger-ui-bundle.js"></script>
<script src="./swagger-ui-standalone-preset.js"></script>
<script>
    window.onload = function () {
       // window["SwaggerUIBundle"] = window["swagger-ui-bundle"]
       // window["SwaggerUIStandalonePreset"] = window["swagger-ui-standalone-preset"]
        protocol = location.protocol
        domain = location.hostname
        port = location.port
        baseURL = protocol + "//" + domain + (port == "" ? "" : ":" + port) + "%s"
        // Build a system
        const ui = SwaggerUIBundle({
            url: baseURL,
            dom_id: "#swagger-ui",
            presets: [
                SwaggerUIBundle.presets.apis,
                SwaggerUIStandalonePreset
            ],
            plugins: [
                SwaggerUIBundle.plugins.DownloadUrl
            ],
            layout: "StandaloneLayout"
        })

        window.ui = ui

        ui.initOAuth({
            clientId: "your-client-id",
            clientSecret: "your-client-secret-if-required",
            realm: "your-realms",
            appName: "your-app-name",
            scopeSeparator: " ",
            additionalQueryStringParams: {},
            usePkceWithAuthorizationCodeGrant: false
        })
    }
</script>
</body>

</html>

`

func RendCustomIndexHtml(openAPIDocPath string) string {
	return fmt.Sprintf(indexHTML, openAPIDocPath)
}
