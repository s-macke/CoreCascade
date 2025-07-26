set -e

gnuplot plot.gp
gnuplot probes.gp

convert -delay 100 -loop 0 probes*.png movie.gif