# hexmap

A small Go library for hex-grid geometry, with SVG rendering.

## API

- **`Point`** — integer x/y with `Plus`, `Sub`, `Equals`, `InBounds`, `String`.
- **`Hexagon`** — `Center Point`, `Radius float64`, `ID string`. Provides
  `Vertices()` and pointy-top neighbour helpers: `CopyTop`, `CopyBottom`,
  `CopyTopLeft`, `CopyTopRight`, `CopyBottomLeft`, `CopyBottomRight`.
- **`Map`** — `Width`, `Height`, `HexRadius`, `Units`. `Hexes(populatePercent)`
  recursively fills the map with hexes (sorted, de-duplicated by centre) and can
  emit an SVG.

Rendering uses [`github.com/ajstarks/svgo`](https://github.com/ajstarks/svgo).

## Demo

```bash
go run ./cmd > hexes.svg      # a hex plus its six neighbours, labelled
```

## Test

```bash
go test ./...
```
