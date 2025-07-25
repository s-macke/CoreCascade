set -e

./CoreCascade -scene penumbra -method vanilla_radiance_cascade -output assets/penumbra
rm assets/penumbra.raw
./CoreCascade -scene center -method vanilla_radiance_cascade -output assets/center
rm assets/center.raw
./CoreCascade -scene pinhole -method vanilla_radiance_cascade -output assets/pinhole
rm assets/pinhole.raw
./CoreCascade -scene shadows -method vanilla_radiance_cascade -output assets/shadows
rm assets/shadows.raw
./CoreCascade -scene beam -method vanilla_radiance_cascade -output assets/beam
rm assets/beam.raw

 #(cd plots && gnuplot plot.gp)