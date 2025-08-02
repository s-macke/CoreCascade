![Title](/assets/title.webp)

# CoreCascade

A 2D implementation of the Radiance Cascades global illumination technique in Go. This project serves as an experimental platform for understanding and exploring the Radiance Cascades algorithm in a simplified 2D environment.

Its benefits include:

* No noise
* Good convergence against satisfying results
* Scales independent of number of light sources 
* Efficiently captures global illumination
* Does run in real-time if implemented on a GPU

## Overview

Radiance Cascades is a hierarchical probe-based global illumination technique that efficiently computes indirect lighting by using multiple resolution levels (cascades). 
A good tutorial on the technique can be found [here](https://m4xc.dev/articles/fundamental-rc/).
Here is a short overview of the technique.

## Path Tracing in 2D

Path tracing is a Monte Carlo ray-tracing technique that fires many random rays from a probe position into a scene 
to simulate the way light interacts with surfaces. Each ray traces a path through the scene, bouncing off surfaces and 
gathering color information. 

![path_tracing_cascade0.png](plots/path_tracing_cascade0.png)

The final probe color is produced by averaging the contributions of all these rays.
In 2D those probes are just points in the scene ordered on a grid and the rays are casted from these points into the scene.

![path_tracing_spatial_probes.png](plots/path_tracing_spatial_probes.png)

In the trivial case, each probe sits in the center of a pixel. So, having several million probes per image. 
This method is simple, but due to the stochastic nature noisy and can be slow to converge. 

## Penumbra Condition (Soft Shadows Condition)

The centerpiece of Radiance Cascade is to exploit certain geometric properties of light tracing.

![penunmbra_annotated.png](assets/penunmbra_annotated.png)

* We need **less** angular resolution next to the light source
* We need **more** spatial resolution next to the light source.

# Angular Component
We discretize the angular space dependent on the distance to the light probe.

- In our example each cascade level has:
  - 4x more angular directions
  - 2x longer ray intervals

![cascades_non_spatial.png](assets/cascades_non_spatial.png)

# Spatial Component

Instead of probing all cascades at the same position, we use a hierarchical grid of probes that are placed at different spatial resolutions.

- Each cascade level has around:
  - 2x less probes and hence less spatial resolution
  - is displaced to the center of the previous cascade level

This image shows the center of each probe of each cascade level
- ![probe_center.png](assets/probe_center.png)

Cascade 0 depicts the probes with the highest spatial resolution and usually are placed at the center of each pixel.

Some of the spatial probes of the higher cascades are outside of the image and must be handled as well.
The reason for this is, that we perform a bilinear interpolation between the probes of the different cascades.
E. g. the result of the probe tracing of cascade i+1 is bilinearly interpolated to the positions of the probes of cascade i.
However, most of the implementation set the probe trace result to zero (black color), which seems to be sufficient for most cases.

# Angular and Spatial Probes

We end up with a hierarchical grid of "probes" at multiple resolutions.
![probes.gif](assets/probes.gif)

## Merging

With bilinear interpolation the trace information of cascade i+1 is estimated at the probe of cascade i.
![marge_vanilla.gif](assets/merge_vanilla.gif)

# Bilinear Fix

The vanilla version has several issues such as ringing and light leaking. 
So, currently several ideas are being tested to improve the overall quality of the rendering.
One of the most prominent is called bilinear fix, but requires four times more rays to be casted.
If traces the rays from their usual start position in cascade i to the start position of the child rays 
in cascade i+1 for each of the four probes. 

To make it easier to visualize, only the rays from cascade 0 contain the bilinear fix, while the rays from cascade 1 are the same as in the vanilla version.
![merge_bilinear_fix.gif](assets/merge_bilinear_fix.gif)

While I am not sure, how the bilinear fix works, the results are much better than the vanilla version.

# Results
Comparison between different rendering methods in the same scene.

![Center Circle](/assets/center.webp)
![Penunbra](/assets/penumbra.webp)
![Color Pinhole](/assets/pinhole.webp)
![Shadows](/assets/shadows.webp)
![Long Pinhole](/assets/beam.webp)

## Features

- **Multiple Rendering Methods**:
  - Path tracing
  - Radiance Cascades rendering
  - Light Propagation Volumes
  - Light Propagation (Under development)
- **SDF-based Scene Representation**: Uses Signed Distance Fields for flexible object representation

## Getting Started

### Building

Build directly:
```bash
cd src && go build -o ../cascade
```

### Running

To see the options use
```bash
./CoreCascade --help
```


## Scene Definition

Scenes use Signed Distance Fields (SDFs) with the following primitives:
- **Circle**: Defined by center, radius, and color
- **Box**: Axis-aligned box with center, dimensions, and color

## References

See `links.md` for an extensive collection of:
- Original papers and research
- Implementation references in various languages
- Tutorial articles and blog posts
- Community resources and discussions

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
