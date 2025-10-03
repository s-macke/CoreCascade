set terminal pngcairo size 2048,2048 enhanced font 'Verdana,40'

set output 'angular_distribution_3d.png'
set key inside top left vertical maxrows 1 sample 0.1
set key font ",40"
unset colorbox
set size ratio -1
set border lw 4
set xlabel "u"
set ylabel "v"

set title "Angular Distributeion of Cascades in 3D"

plot [0:1] [0:1]          \
"probes3D.data" u 1:2:(0.01) every :::0::0 w circles fill solid title "cascade 0", \
"probes3D.data" u 1:2:(0.01) every :::1::1 w circles fill solid title "cascade 0"

