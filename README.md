# Go Named Formatter

This is like Go fmt package but offers named format. Parameters will be parsed to function as map[string]interface{}.

## Install

```
go get -u github.com/chonla/format
```

## Usage

Put `<name>` in between `%` and format. For example, `%<name>s`.

## Naming rule

Name can be any combination of A-Z, a-z, 0-9 or _ character.

## Example

```
var params = map[string]interface{}{
    "sister": "Susan",
    "brother": "Louis",
}
format.Printf("%<brother> loves %<sister>.", params)
```

## Output

```
Louis love Susan.
```

## Tips

Named variable can be reused.

```
var params = map[string]interface{}{
    "sister": "Susan",
    "brother": "Louis",
}
format.Printf("%<brother>s loves %<sister>s. %<sister>s also loves %<brother>s.", params)
```

## Output

```
Louis loves Susan. Susan also loves Louis.
```

## Available methods

### Printf

Printf formats string with named parameters and writes to standard output.

### Printfln

Printfln is like Printf but a newline is appended.

### Sprintf

Sprintf formats string with named parameters and return the result.

### Sprintfln

Sprintfln is like Sprintf but a newline is appended.

## License

[MIT](./LICENSE.txt)