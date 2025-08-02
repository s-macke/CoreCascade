set -e


combine_images () {
magick convert vanilla_${1}.png -gravity North -pointsize 30 -font helvetica -fill white -annotate +0+10 'Vanilla Radiance Cascade' -background black -gravity East -splice 20x0 vanilla_${1}2.png
magick convert bilinear_fix_${1}.png -gravity North -pointsize 30 -font helvetica -fill white -annotate +0+10 'Bilinear Fix Radiance Cascade' -background black -gravity East -splice 20x0 bilinear_fix_${1}2.png
magick convert path_tracing_${1}.png -gravity North -pointsize 30 -font helvetica -fill white -annotate +0+10 'Path Tracing (Reference)' path_tracing_${1}2.png
magick convert +append vanilla_${1}2.png bilinear_fix_${1}2.png path_tracing_${1}2.png -alpha off -quality 70 -define webp:alpha-compression=0 -define webp:method=6  ${1}.webp
rm vanilla_${1}2.png bilinear_fix_${1}2.png path_tracing_${1}2.png
}

#rm -f bilinear_fix_title.raw
#rm -f vanilla_title.raw
#magick convert vanilla_title.png title.webp
magick convert bilinear_fix_title.png -alpha off -quality 80 -define webp:alpha-compression=0 title.webp

combine_images "beam"
combine_images "center"
combine_images "penumbra"
combine_images "pinhole"
combine_images "shadows"


