# Paper
https://drive.google.com/file/d/1L6v1_7HY2X-LV3Ofb6oyTIxgEaP4LOI6/view
https://github.com/Raikiri/RadianceCascadesPaper/blob/main/out_latexmk2/RadianceCascades.pdf
https://github.com/Raikiri/RadianceCascadesPaper

# Video
Video: https://www.youtube.com/watch?v=3so7xdZHKxw
Website: https://simondev.io/projects
GitHub: https://github.com/simondevyoutube/Shaders_RadianceCascades/

# Discord
https://discord.gg/WSW7d2wrps

# Blog article
https://jason.today/gi
https://jason.today/rc
https://github.com/jasonjmcghee

# Demo Website
https://radiance-cascades.com/
https://github.com/radiance-cascades/radiance-cascades.com

# Blog articles
https://mini.gmshaders.com/
https://mini.gmshaders.com/p/yaazarai-gi
https://mini.gmshaders.com/p/radiance-cascades
https://mini.gmshaders.com/p/radiance-cascades2
Code: https://github.com/Yaazarai/GMShadersGI
Code: https://github.com/Yaazarai/RadianceCascades
Code: https://github.com/Yaazarai/GMShaders-Radiance-Cascades

# Blog article
https://m4xc.dev/articles/fundamental-rc/
https://m4xc.dev/blog/surfel-maintenance/
Code: https://github.com/mxcop/src-dgi

# Radiance Cascadee fully written in Rust
https://github.com/entropylost/amida/

# Holographic Radiance Cascades
https://arxiv.org/abs/2505.02041
https://github.com/Yaazarai/Holographic-Radiance-Cascades
https://www.shadertoy.com/view/4ctczl

# Guy tries to write a game engine with Radiance Cascades
https://x.com/MytinoDev

# Another paper explaining Radiance Cascades and Bilinear Fix
https://arxiv.org/abs/2408.14425
https://academic.oup.com/rasti/article/doi/10.1093/rasti/rzae062/7929002

# Solis 2D using rust and bevy-engine
https://github.com/Lommix/solis_2d

# Unity Radiance Cascade implementation
https://github.com/alexmalyutindev/unity-urp-radiance-cascades

# Game Maker Radiance Cascade implementation
https://github.com/Yaazarai/RadianceCascades

# Rust/ WGSL Radiance Cascade implementation
https://github.com/kornelski/bevy_flatland_radiance_cascades

# 2D signed distance functions example:
https://www.shadertoy.com/view/4dfXDn

# volumetric lighting example:
https://www.shadertoy.com/view/tdjBR1

# Clouds 2D with shadowing
https://www.shadertoy.com/view/WddSDr

# Turbulent Flame
https://www.shadertoy.com/view/wffXDr

# Vanilla Radiance Cascade 
https://www.shadertoy.com/view/mtlBzX

# Radiance Cascade Diagram
https://www.shadertoy.com/view/4clcWn

# Radiance Cacade Gear fix
https://www.shadertoy.com/view/XffcD7

# Radiance Cacade Forking fix
https://www.shadertoy.com/view/4clcWn

# RC Experimental Testbed 
https://www.shadertoy.com/view/4ctXD8

# Fast FBM Sphere Tracing
https://www.shadertoy.com/view/MXyBDW

# Nove depth aware upscaling
https://www.shadertoy.com/view/4XXSWS

# Another Blog. Tries is with WGSL
https://tmpvar.com/poc/radiance-cascades/
https://tmpvar.com/poc/radiance-cascades/flatland-2d/radiance-intervals-2d.js

# Radiance Cascade calculator
https://kornel.ski/radiance

# dynamic diffuse global illumination with ray-traced irradiance fields
https://www.jcgt.org/published/0008/02/01/paper-lowres.pdf
https://www.youtube.com/watch?v=0YfO5mFXPos
https://github.com/DaOnlyOwner/hym
https://github.com/Toocanzs/TooD?tab=readme-ov-file
https://developer.nvidia.com/blog/rtx-global-illumination-part-i/
https://github.com/NVIDIAGameWorks/RTXGI-DDGI

# DDA Algorithm Interactive
https://aaaa.sh/creatures/dda-algorithm-interactive/
https://www.shadertoy.com/view/4dX3zl

# Glossar
RC - Radiance Cascade
HRC - Holographic Radiance Cascade
GI - Global Illumination
RMSE - Root mean Square Error
PT - Path Tracing
NEE - Next Event Estimation
DDA - Digital Differential Analyzer

# Improvements
- Bilinear fix
- Holographic Radiance Cascades
- Mip-mapping
- Pre-averaging optimization
- Gear fix
- Nearest fix
- Forking fix
- Parallax fix

# RMSE
sqrt( (1/N) * sum( (I - I')^2 ) )
I: intensiry of color value of pixel i in the rendered image
I': intensity of color value of pixel i in the reference image
The difference is usually computed per channel (RGB), and then averaged or handled depending on implementation.

