set terminal pngcairo size 2048,2048 enhanced font 'Verdana,40'

set border lw 4

set output 'cascades_non_spatial.png'
#set title "Radiance Cascades (non-spatial shift)"
#unset key
unset colorbox
set size ratio -1

set format x ""
set format y ""

set key opaque

plot [-0.1:0.1] [-0.1:0.1]          \
"probe_cascades_non_spatial.data" u 1:2:3:4 every :::0::0 w vectors lt 1 lw 10 title "cascade 0", \
"probe_cascades_non_spatial.data" u 1:2:3:4 every :::1::1 w vectors lt 2 lw 10 title "cascade 1", \
"probe_cascades_non_spatial.data" u 1:2:3:4 every :::2::2 w vectors lt 3 lw 10 title "cascade 2", \
"probe_cascades_non_spatial.data" u 1:2:3:4 every :::3::3 w vectors lt 4 lw 10 title "cascade 3"

#"probe_cascades_non_spatial.data" u 1:2:3:4 every :::0:0:0 w vectors lt 1 lw 20 title "", \
#"probe_cascades_non_spatial.data" u 1:2:3:4 every :::1:3:1 w vectors lt 2 lw 20 title "", \
#"probe_cascades_non_spatial.data" u 1:2:3:4 every :::2:15:2 w vectors lt 3 lw 20 title "", \
#"probe_cascades_non_spatial.data" u 1:2:3:4 every :::3:63:3 w vectors lt 4 lw 20 title ""

###

set output 'path_tracing_cascade0.png'
#set title "Monte Carlo Path Tracing"
unset xtics
unset ytics
unset border


do for [i=1:75] {
  # Generate a random angle and radius
  angle = rand(0) * 2 * pi
  radius = 2.0

  # Convert polar coordinates to Cartesian for the endpoint
  x_end = radius * cos(angle)
  y_end = radius * sin(angle)

  # Define an arrow from the center (0,0) to the random endpoint
  # 'head' styles the arrowhead
  # 'lc palette' assigns a random color from the defined palette
  set arrow i from 0,0 to x_end,y_end head filled size screen 0.02,15,45 lw 10 lt 1
}

plot [-2:2] [-2:2] NaN notitle

###

set output 'path_tracing_spatial_probes.png'
#set title "Monte Carlo Path Tracing"
set style fill solid
unset title
unset arrow
unset xtics
unset ytics
unset border

idx = 1
do for [y=-2:2] {
do for [x=-2:2] {
do for [i=1:75] {
  # Generate a random angle and radius
  angle = rand(0) * 2 * pi
  radius = 10.0

  # Convert polar coordinates to Cartesian for the endpoint
  x_start = x
  y_start = y
  x_end = x + radius * cos(angle)
  y_end = y + radius * sin(angle)

  # Define an arrow from the center (0,0) to the random endpoint
  # 'head' styles the arrowhead
  # 'lc palette' assigns a random color from the defined palette
  set arrow idx from x_start,y_start to x_end,y_end head filled size screen 0.02,15,45 lw 10 lt 1 lc rgb "0xD09400d3"
  set obj idx circle at x_start,y_start size scr 0.01 fc rgb "red" front
  idx = idx + 1
}
}
}

plot [-3:3] [-3:3] NaN notitle

