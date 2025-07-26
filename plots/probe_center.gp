set terminal pngcairo size 2048,2048 enhanced font 'Verdana,40'

set output '../assets/probe_center.png'
set key inside top left vertical maxrows 1 sample 0.1
set key font ",40"
unset colorbox
set size ratio -1
set border lw 4
set xlabel "x"
set ylabel "y"

set title "Probe Center of Cascades for a 32x32 Grid"

unset arrow
set arrow from -1.0,-1.0 to -1.0,1.0 nohead lw 2 lc rgb "black"
set arrow from -1.0,-1.0 to 1.0,-1.0 nohead lw 2 lc rgb "black"
set arrow from 1.0,-1.0 to 1.0,1.0 nohead lw 2 lc rgb "black"
set arrow from -1.0,1.0 to 1.0,1.0 nohead lw 2 lc rgb "black"


plot [-2:2] [-2:2]          \
"probe_center.data" u 1:2:(0.01) every :::0::0 w circles fill solid title "cascade 0", \
"probe_center.data" u 1:2:(0.02) every :::1::1 w circles fill solid title "cascade 1", \
"probe_center.data" u 1:2:(0.03) every :::2::2 w circles fill solid title "cascade 2", \
"probe_center.data" u 1:2:(0.04) every :::3::3 w circles fill solid title "cascade 3", \
"probe_center.data" u 1:2:(0.05) every :::4::4 w circles fill solid title "cascade 4"
