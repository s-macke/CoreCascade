# Set the output terminal and file name
set terminal pngcairo size 800,800
set output 'penumbra_annotated.png'

# Remove plot elements like tics, border, and legend
unset tics
unset border
unset key

# Set the plot ranges to match the image dimensions (if known)
set xrange [0:800]
set yrange [0:800]

set arrow 1 from 10,600 to 800,0 nohead filled lw 2 lc rgb "white" front
set arrow 2 from 10,200 to 800,800 nohead filled lw 2 lc rgb "white" front

set arrow 3 from 400,300 to 400,500 heads filled lw 2 lc rgb "white" front

set label 1 "Penunmbra" at 450,400 font "Arial,20" textcolor rgb "white" front
set label 2 "Umbra" at 450,100 font "Arial,20" textcolor rgb "white" front

# Plot the image as the background
# The 'flipy' option is often needed because image and plot y-coordinates can be inverted.
plot '../assets/bilinear_fix_penumbra.png' binary filetype=png with rgbimage