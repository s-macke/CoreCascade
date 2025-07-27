set -e

gnuplot penumbra.gp
gnuplot probe_center.gp

gnuplot probes.gp
magick convert -delay 100 -loop 0 probes*.png ../assets/probes.gif

gnuplot probe_non_spatial.gp
magick convert +append path_tracing_cascade0.png cascades_non_spatial.png ../assets/cascade_comparison.png



