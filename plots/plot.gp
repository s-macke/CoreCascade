
set terminal pngcairo size 800,800 enhanced font 'Verdana,10'
set output 'plot3.png'
#unset key
unset colorbox
set size ratio -1
#set samples 17    # x-axis
#set isosamples 15 # y-axis

plot            \
"plot3.data" u 1:2:3:4 every :::0::0 w vectors lw 2 title "cascade 0", \
"plot3.data" u 1:2:3:4 every :::1::1 w vectors lw 2 title "cascade 1", \
"plot3.data" u 1:2:3:4 every :::2::2 w vectors lw 2 title "cascade 2", \
"plot3.data" u 1:2:3:4 every :::3::3 w vectors lw 2 title "cascade 3"


###

set output 'plot.png'
unset key
unset colorbox
set size ratio -1
#set samples 17    # x-axis
#set isosamples 15 # y-axis

plot            \
"plot.data" u 1:2:3:4 every :::0::0 w vectors lw 2, \
"plot.data" u 1:2:3:4 every :::1::1 w vectors lw 2,  \
"plot.data" u 1:2:3:4 every :::2::2 w vectors lw 2

###

set arrow from -1.0,-2.0 to -1.0,1.0 nohead lw 2 lc rgb "black"
set arrow from -2.0,-1.0 to 1.0,-1.0 nohead lw 2 lc rgb "black"

set output 'plot2.png'
unset key
unset colorbox
set size ratio -1

plot [-1.005:-0.98] [-1.005:-0.98]          \
"plot2.data" u 1:2:3:4 every :::0::0 w vectors lw 2, \
"plot2.data" u 1:2:3:4 every :::1::1 w vectors lw 2, \
"plot2.data" u 1:2:3:4 every :::2::2 w vectors lw 2, \



