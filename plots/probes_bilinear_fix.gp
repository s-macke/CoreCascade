set terminal pngcairo size 1024,1024 enhanced font 'Verdana,40'

unset key
unset colorbox
set size ratio -1
set border lw 2
#set xlabel 'x'
#set ylabel 'y'

###

set output 'probes1_bilinear_fix.png'
plot [-2:2] [-2:2]           \
"probes_bilinear_fix.data" u 1:2:3:4 every :::0::0 w vectors lw 2 lc rgb "0x80AF0000"

###

set output 'probes2_bilinear_fix.png'
plot [-2:2] [-2:2]           \
"probes_bilinear_fix.data" u 1:2:3:4 every :::0::0 w vectors lw 2 lc rgb "0x80AF0000", \
"probes_bilinear_fix.data" u 1:2:3:4 every :::1::1 w vectors lw 2 lc rgb "0x6000AF00"

set output 'probes3_bilinear_fix.png'
plot [-2:2] [-2:2]            \
"probes_bilinear_fix.data" u 1:2:3:4 every :::0::0 w vectors lw 2 lc rgb "0x80AF0000", \
"probes_bilinear_fix.data" u 1:2:3:4 every :::1::1 w vectors lw 2 lc rgb "0x6000AF00",  \
"probes_bilinear_fix.data" u 1:2:3:4 every :::2::2 w vectors lw 2 lc rgb "0x400000AF"
