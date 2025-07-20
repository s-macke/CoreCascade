# CoreCascade

A 2D implementation of the Radiance Cascades global illumination technique in Go. This project serves as an experimental platform for understanding and exploring the Radiance Cascades algorithm in a simplified 2D environment.

## Overview

Radiance Cascades is a hierarchical probe-based global illumination technique that efficiently computes indirect lighting by using multiple resolution levels (cascades). This implementation provides both traditional path tracing and the Radiance Cascades method for comparison.

## Features

- **Multiple Rendering Methods**:
  - Standard path tracing
  - Parallel path tracing
  - Radiance Cascades rendering
- **SDF-based Scene Representation**: Uses Signed Distance Fields for flexible object representation
- **Configurable Cascade Parameters**: Adjust cascade levels, angular resolution, and spatial resolution
- **Visualization Tools**: Automatic generation of plots showing cascade behavior

## Getting Started

### Prerequisites

- Go 1.23.9 or later
- Standard Go toolchain

### Building

```bash
./build.sh
```

Or build directly:
```bash
cd go && go build -o ../cascade
```

### Running

```bash
./cascade
```

This will:
- Generate `output.png` with the rendered scene
- Create visualization plots in the `plots/` directory

### Switching Rendering Methods

Edit `go/main.go` and uncomment your desired rendering method:

```go
// Choose one:
img := RenderPathTracing(scene)           // Standard path tracing (default)
//img := RenderPathTracingParallel(scene)  // Parallel path tracing
//img := RenderCascade(scene)             // Radiance Cascades
```

## Algorithm Overview

### Path Tracing
- Traditional Monte Carlo approach
- Random ray directions in 2π circle
- Simple, unbiased, but slow convergence

### Radiance Cascades
- Hierarchical grid of "probes" at multiple resolutions
- Each cascade level has:
  - 4x more angular directions
  - 2x finer spatial resolution
  - 4x longer ray intervals
- Bilinear interpolation merges information between levels
- Efficient capture of both local and global illumination

## Scene Definition

Scenes use Signed Distance Fields (SDFs) with the following primitives:
- **Circle**: Defined by center, radius, and color
- **Box**: Axis-aligned box with center, dimensions, and color

## Visualization

The project generates several plots to visualize the cascade behavior:
- Cascade probe positions at different levels
- Ray interval demonstrations
- Angular resolution comparisons

## References

See `links.md` for an extensive collection of:
- Original papers and research
- Implementation references in various languages
- Tutorial articles and blog posts
- Community resources and discussions

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

Based on the Radiance Cascades technique by Alexander Sannikov. See `links.md` for the original paper and related resources.