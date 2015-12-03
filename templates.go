package main

const (
	indexTemplate = `
<html><head>
	<title>{{ .Title }}</title>
	<link href='app.css' rel='stylesheet' />
	<script src='app.js'></script>
</head><body>
	<div class='main'>
		<h1>Swift Credential Test</h1>
		<form action='{{ .Action }}' method='post' id='testform' target='_self'>
			<fieldset>
				<legend>{{ .Heading }}</legend>
				{{ range .Variables }}<p>
					<label for='{{ . }}'>{{ . }}</label><br />
					<input name='{{ . }}' type='text' required />
				</p>
				{{ end }}

			</fieldset>
			<button>Submit</button>
		</form>
	</div>
</body></html>
`
	swifttestTemplate = `
<html><head>
	<title>{{ if .Success }}It Works!{{ else }}Problem :({{ end }}</title>
	<link href='app.css' rel='stylesheet' />
	<script src='app.js'></script>
</head><body>
	<div class='main'>
	<h1>{{ if .Success }}It Works!{{ else }}There was a problem{{ end }}</h1>
	{{ with .Error }}<span class='error'>{{ . }}</span>{{ end }}
	{{ with .Vars }}
	<table>
		<tr><th>Key</th><th>Value</th></tr>
		{{ range $k, $v := . }}<tr><td>{{ $k }}</td><td>{{ $v }}</td></tr>
		{{ end }}
	</table>
	{{ end }}

	{{ with .Items }}
	<h3>Found</h3>
	<ul>
		{{ range . }}<li>{{ . }}</li>{{ end }}
	</ul>
	{{ end }}

	</div>
</body></html>
`
)
