set -e


combine_images () {
magick convert vanilla_${1}.png -gravity North -pointsize 30 -font helvetica -fill white -annotate +0+10 'Vanilla Radiance Cascade' -background black -gravity East -splice 20x0 vanilla_${1}2.png
magick convert path_tracing_${1}.png -gravity North -pointsize 30 -font helvetica -fill white -annotate +0+10 'Reference' path_tracing_${1}2.png
magick convert +append vanilla_${1}2.png path_tracing_${1}2.png ${1}.jpg
}

combine_images "beam"
combine_images "center"
combine_images "penumbra"
combine_images "pinhole"
combine_images "shadows"


