set -e

gnuplot penumbra.gp
magick convert penumbra_annotated.png -alpha off -set colorspace Gray -separate -average ../assets/penumbra_annotated.png

gnuplot probe_center.gp
magick probe_center.png -alpha off -resize 50% -colors 64 ../assets/probe_center.png

gnuplot probes.gp
magick convert -delay 100 -loop 0 probes1.png probes2.png probes3.png -resize 50% -colors 64 ../assets/probes.gif

gnuplot probes_bilinear_fix.gp
magick convert -delay 100 -loop 0 probes*_bilinear_fix.png -resize 50% -colors 64 ../assets/probes_bilinear_fix.gif

gnuplot probes_bilinear_fix_simple.gp
magick convert -delay 100 -loop 0 probes1_bilinear_fix_simple2_*.png -resize 50% -colors 64 ../assets/merge_vanilla.gif
magick convert -delay 100 -loop 0 probes1_bilinear_fix_simple1_*.png -resize 50% -colors 64 ../assets/merge_bilinear_fix.gif

gnuplot probe_non_spatial.gp
#magick convert +append path_tracing_cascade0.png cascades_non_spatial.png ../assets/cascade_comparison.png
#cp cascades_non_spatial.png ../assets/

magick convert cascades_non_spatial.png -alpha off -resize 50% -colors 64 ../assets/cascades_non_spatial.png
magick convert path_tracing_cascade0.png -alpha off -resize 50% -colors 64 ../assets/path_tracing_cascade0.png
magick path_tracing_spatial_probes.png -alpha off -resize 50% -colors 64 ../assets/path_tracing_spatial_probes.png

echo "Done"