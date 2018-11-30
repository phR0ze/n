# tmpl
tmpl does simple template substitutions fast

## Usage
```Golang
data := `labels:
  chart: {{ name }}:{{ version }}
  release: {{ Release.Name }}
  heritage: {{ Release.Service }}`
t, _ := tmpl.New(data, "{{", "}}")
result, err := t.Process(map[string]string{
	"name":            "foo",
	"version":         "1.0.2",
	"Release.Name":    "babble",
	"Release.Service": "fish",
})

// (Results)
//-----------------------
// labels:
//   chart: foo:1.0.2
//   release: babble
//   heritage: fish
```