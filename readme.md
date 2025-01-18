# opname

Operation friendly name generator.

Generated names satisfy following:

- Has given prefix
- Easy to copy and past
- Valid as DNS label, i.e.,
  - contain at most 63 characters
  - contain only lowercase alphanumeric characters or '-'
  - start with an alphanumeric character
  - end with an alphanumeric character
- Has nick name
- The order generated follows lexicographical order

## Format

```
<prefix>-<date>-<time>-<pretty name>
```

### Example

```
usr20250118215723red
```

## Usage

### CLI

1. Install: `go install github.com/naoyafurudono/opname:latest`
2. Run: `opname <prefix>`.

### Use as a library

`go install github.com/naoyafurudono/opname:latest`

See godoc for details.
