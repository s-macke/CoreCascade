set -e

gnuplot probe_center.gp
gnuplot probes.gp

convert -delay 100 -loop 0 probes*.png ../assets/probes.gif