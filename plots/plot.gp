
set terminal pngcairo size 2048,2048 enhanced font 'Verdana,40'

###

set output 'plot3.png'
unset colorbox
set size ratio -1
#set samples 17    # x-axis
#set isosamples 15 # y-axis
set border lw 4
set xlabel "x"
set ylabel "y"

plot            \
"plot3.data" u 1:2:3:4 every :::0::0 w vectors lw 2 title "cascade 0", \
"plot3.data" u 1:2:3:4 every :::1::1 w vectors lw 2 title "cascade 1", \
"plot3.data" u 1:2:3:4 every :::2::2 w vectors lw 2 title "cascade 2", \
"plot3.data" u 1:2:3:4 every :::3::3 w vectors lw 2 title "cascade 3"


###

unset arrow
set arrow from -1.0,-1.0 to -1.0,1.0 nohead lw 2 lc rgb "black"
set arrow from -1.0,-1.0 to 1.0,-1.0 nohead lw 2 lc rgb "black"
set arrow from 1.0,-1.0 to 1.0,1.0 nohead lw 2 lc rgb "black"
set arrow from -1.0,1.0 to 1.0,1.0 nohead lw 2 lc rgb "black"

set output 'plot2.png'
unset key
unset colorbox
set size ratio -1

plot [-1.005:-0.98] [-1.005:-0.98]          \
"plot2.data" u 1:2:3:4 every :::0::0 w vectors lw 2, \
"plot2.data" u 1:2:3:4 every :::1::1 w vectors lw 2, \
"plot2.data" u 1:2:3:4 every :::2::2 w vectors lw 2, \


###

set output 'plot4.png'
set key inside bottom left vertical maxrows 1 sample 0.1
set key font ",44"
unset colorbox
set size ratio -1

plot [:] [:]          \
"plot4.data" u 1:2:(0.01) every :::0::0 w circles fill solid title "cascade 0", \
"plot4.data" u 1:2:(0.02) every :::1::1 w circles fill solid title "cascade 1", \
"plot4.data" u 1:2:(0.03) every :::2::2 w circles fill solid title "cascade 2", \
"plot4.data" u 1:2:(0.04) every :::3::3 w circles fill solid title "cascade 3", \
"plot4.data" u 1:2:(0.05) every :::4::4 w circles fill solid title "cascade 4", \
"plot4.data" u 1:2:(0.06) every :::5::5 w circles fill solid title "cascade 5", \

###

unset arrow

set output 'plot5.png'
#unset key
unset colorbox
set size ratio -1

plot [:] [:]          \
"plot5.data" u 1:2:3:4 every :::0::0 w vectors lw 10 title "cascade 0", \
"plot5.data" u 1:2:3:4 every :::1::1 w vectors lw 10 title "cascade 1", \
"plot5.data" u 1:2:3:4 every :::2::2 w vectors lw 10 title "cascade 2"

###

unset arrow
set output 'plotEnergy.png'
unset key
unset colorbox
set size ratio 1
set xlabel "x"
set ylabel "Energy"


plot [0:] [0:] "plotEnergy.data" u 1:($2*$1) w lines lw 10 title "Energy"
