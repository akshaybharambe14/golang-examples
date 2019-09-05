# GO 1.13 xerrors examples

[Evolution of error checking](https://crawshaw.io/blog/xerrors)

## Important

```go
xerrors.Errorf("%q: %w", s, errorFromLevelTwo)
```

Here the ':' after '%q' is important. If not provided, xerrors fails to identify the error by `xerrors.Is()`
