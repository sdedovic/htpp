# HTML/Template (Plus Plus)
## htpp - opinionated improvements to Go HTML templating
Really, this is just so I can
1. Have a single executable to render a template, which
2. Using `extends` syntax to define dependencies, to
3. Allow me to make a ParcelJS plugin, so
4. I can develop frontends with the Parcel build system and development server that use Go templates

---
## Current Status - **Work In Progress**
### Syntax
#### `extends` syntax
```
extends <relative-path-to-template>
```
See: [examples/simple](./examples/simple)
### CLI
```
htpp

Usage:
  htpp [--std-in] <filename>
  htpp [--print-dependencies] <filename>

Options:
  --std-in              Read JSON applied to template from stdin
  --print-dependencies  Print all dependencies of the supplied template. Does not execute the template.
````

## Future Thoughts / ToDo
- [ ] public api to render template
  - [ ] support non-file (in-memory, `string`) templates
  - [ ] support directory of templates
- [ ] cache template compilation for future runs
- [ ] other relative path resolvers? See below.
  
### File resolvers
- Each template has an identifier
- if the identifier is a filename, template can be resolved relatively, absolutely, by name with or without ext
- how are conflicts handled? e.g. `foo` and `foo.tmpl`?

Relative File - **DONE**
```gotemplate
extends ./base.htpp

{{block "body"}}
<div>
    <p>This is a test</p>
</div>
{{end}}
```

Relative File - **DONE**
```gotemplate
extends ../base.htpp

{{block "body"}}
<div>
    <p>This is a test</p>
</div>
{{end}}
```

Absolute File - **PROPOSED**
```gotemplate
extends /base.htpp

{{block "body"}}
<div>
    <p>This is a test</p>
</div>
{{end}}
```

Identifier or Adjacent File - **PROPOSED**
```gotemplate
extends base.htpp

{{block "body"}}
<div>
    <p>This is a test</p>
</div>
{{end}}
```

Identifier - **PROPOSED**
```gotemplate
extends base

{{block "body"}}
<div>
    <p>This is a test</p>
</div>
{{end}}
```