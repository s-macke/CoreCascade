set terminal pngcairo size 1024,1024 enhanced font 'Verdana,30'

#unset key
unset colorbox
set size ratio -1
set border lw 2
set xlabel 'x'
set ylabel 'y'

set format x ""
set format y ""

set key opaque box bottom

###


set output 'probes1_bilinear_fix_simple1_0.png'
plot [:] [:]           \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every :::0::0 w vectors lw 2 lc rgb "0xA0AF0000" title "Bilinear Fix Cascade 0", \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every :::1::1 w vectors lw 2 lc rgb "0xA000AF00" title "Vanilla Cascade 1",  \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every ::0:0:3:0 w vectors lw 8 lc rgb "0x00AF0000" title "", \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every ::0:1:15:1 w vectors lw 8 lc rgb "0x0000AF00" title ""

set output 'probes1_bilinear_fix_simple1_1.png'
plot [:] [:]           \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every :::0::0 w vectors lw 2 lc rgb "0xA0AF0000" title "Bilinear Fix Cascade 0", \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every :::1::1 w vectors lw 2 lc rgb "0xA000AF00" title "Vanilla Cascade 1",  \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every ::4:0:7:0 w vectors lw 8 lc rgb "0x00AF0000" title "", \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every ::16:1:31:1 w vectors lw 8 lc rgb "0x0000AF00" title ""

set output 'probes1_bilinear_fix_simple1_2.png'
plot [:] [:]           \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every :::0::0 w vectors lw 2 lc rgb "0xA0AF0000" title "Bilinear Fix Cascade 0", \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every :::1::1 w vectors lw 2 lc rgb "0xA000AF00" title "Vanilla Cascade 1",  \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every ::8:0:11:0 w vectors lw 8 lc rgb "0x00AF0000" title "", \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every ::32:1:47:1 w vectors lw 8 lc rgb "0x0000AF00" title ""

set output 'probes1_bilinear_fix_simple1_3.png'
plot [:] [:]           \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every :::0::0 w vectors lw 2 lc rgb "0xA0AF0000" title "Bilinear Fix Cascade 0", \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every :::1::1 w vectors lw 2 lc rgb "0xA000AF00" title "Vanilla Cascade 1",  \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every ::12:0:15:0 w vectors lw 8 lc rgb "0x00AF0000" title "", \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every ::48:1:63:1 w vectors lw 8 lc rgb "0x0000AF00" title ""

###

set output 'probes1_bilinear_fix_simple2_0.png'
plot [:] [:]           \
"probes_bilinear_fix_simple.data" u 1:2:5:6 every :::0::0 w vectors lw 2 lc rgb "0x80AF0000" title "Vanilla Cascade 0", \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every :::1::1 w vectors lw 2 lc rgb "0x6000AF00"  title "Vanilla Cascade 1", \
"probes_bilinear_fix_simple.data" u 1:2:5:6 every ::0:0:3:0 w vectors lw 8 lc rgb "0x00AF0000" title "", \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every ::0:1:15:1 w vectors lw 8 lc rgb "0x0000AF00" title ""


set output 'probes1_bilinear_fix_simple2_1.png'
plot [:] [:]           \
"probes_bilinear_fix_simple.data" u 1:2:5:6 every :::0::0 w vectors lw 2 lc rgb "0x80AF0000" title "Vanilla Cascade 0", \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every :::1::1 w vectors lw 2 lc rgb "0x6000AF00"  title "Vanilla Cascade 1", \
"probes_bilinear_fix_simple.data" u 1:2:5:6 every ::4:0:7:0 w vectors lw 8 lc rgb "0x00AF0000" title "", \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every ::16:1:31:1 w vectors lw 8 lc rgb "0x0000AF00" title ""


set output 'probes1_bilinear_fix_simple2_2.png'
plot [:] [:]           \
"probes_bilinear_fix_simple.data" u 1:2:5:6 every :::0::0 w vectors lw 2 lc rgb "0x80AF0000" title "Vanilla Cascade 0", \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every :::1::1 w vectors lw 2 lc rgb "0x6000AF00"  title "Vanilla Cascade 1", \
"probes_bilinear_fix_simple.data" u 1:2:5:6 every ::8:0:11:0 w vectors lw 8 lc rgb "0x00AF0000" title "", \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every ::32:1:47:1 w vectors lw 8 lc rgb "0x0000AF00" title ""

set output 'probes1_bilinear_fix_simple2_3.png'
plot [:] [:]           \
"probes_bilinear_fix_simple.data" u 1:2:5:6 every :::0::0 w vectors lw 2 lc rgb "0x80AF0000" title "Vanilla Cascade 0", \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every :::1::1 w vectors lw 2 lc rgb "0x6000AF00"  title "Vanilla Cascade 1", \
"probes_bilinear_fix_simple.data" u 1:2:5:6 every ::12:0:15:0 w vectors lw 8 lc rgb "0x00AF0000" title "", \
"probes_bilinear_fix_simple.data" u 1:2:3:4 every ::48:1:63:1 w vectors lw 8 lc rgb "0x0000AF00" title ""

