set -e

rotate_batch () {

for i in $(seq 0 50);
do
    ID=$(printf '%02d\n' "$i")
    echo "=== Vanilla Shadows ${ID} ==="
    ./CoreCascade -scene shadows -time "${i}" -method vanilla_radiance_cascade -output "anim/vanilla_shadows_${ID}"
    rm "anim/vanilla_shadows_${ID}.raw"
done

}


path_tracing_batch () {

echo "=== Path Tracing Penumbra ==="
./CoreCascade -scene penumbra -method path_tracing_parallel -output assets/path_tracing_penumbra -input assets/path_tracing_penumbra.raw

echo "=== Path Tracing Center ==="
./CoreCascade -scene center -method path_tracing_parallel -output assets/path_tracing_center -input assets/path_tracing_center.raw

echo "=== Path Tracing Pinhole ==="
./CoreCascade -scene pinhole -method path_tracing_parallel -output assets/path_tracing_pinhole -input assets/path_tracing_pinhole.raw

echo "=== Path Tracing Shadows ==="
./CoreCascade -scene shadows -method path_tracing_parallel -output assets/path_tracing_shadows -input assets/path_tracing_shadows.raw

echo "=== Path Tracing Beam ==="
./CoreCascade -scene beam -method path_tracing_parallel -output assets/path_tracing_beam -input assets/path_tracing_beam.raw

}

vanilla_radiance_cascade_batch () {

echo "=== Vanilla Penumbra ==="
./CoreCascade -scene penumbra -method vanilla_radiance_cascade -output assets/vanilla_penumbra
rm assets/vanilla_penumbra.raw

echo "=== Vanilla Center ==="
./CoreCascade -scene center -method vanilla_radiance_cascade -output assets/vanilla_center
rm assets/vanilla_center.raw

echo "=== Vanilla Pinhole ==="
./CoreCascade -scene pinhole -method vanilla_radiance_cascade -output assets/vanilla_pinhole
rm assets/vanilla_pinhole.raw

echo "=== Vanilla Shadows ==="
./CoreCascade -scene shadows -method vanilla_radiance_cascade -output assets/vanilla_shadows
rm assets/vanilla_shadows.raw

echo "=== Vanilla Beam ==="
./CoreCascade -scene beam -method vanilla_radiance_cascade -output assets/vanilla_beam
rm assets/vanilla_beam.raw

}

bilinear_fix_radiance_cascade_batch () {

echo "=== Bilinear Fix Penumbra ==="
./CoreCascade -scene penumbra -method bilinear_fix_radiance_cascade -output assets/bilinear_fix_penumbra
rm assets/bilinear_fix_penumbra.raw

echo "=== Bilinear Fix Center ==="
./CoreCascade -scene center -method bilinear_fix_radiance_cascade -output assets/bilinear_fix_center
rm assets/bilinear_fix_center.raw

echo "=== Bilinear Fix Pinhole ==="
./CoreCascade -scene pinhole -method bilinear_fix_radiance_cascade -output assets/bilinear_fix_pinhole
rm assets/bilinear_fix_pinhole.raw

echo "=== Bilinear Fix Shadows ==="
./CoreCascade -scene shadows -method bilinear_fix_radiance_cascade -output assets/bilinear_fix_shadows
rm assets/bilinear_fix_shadows.raw

echo "=== Bilinear Fix Beam ==="
./CoreCascade -scene beam -method bilinear_fix_radiance_cascade -output assets/bilinear_fix_beam
rm assets/bilinear_fix_beam.raw

}


(cd src && go build -o ../CoreCascade)


#./CoreCascade -scene title -method vanilla_radiance_cascade -output assets/vanilla_title
#./CoreCascade -scene title -method bilinear_fix_radiance_cascade -output assets/bilinear_fix_title

#./CoreCascade -method plot
#path_tracing_batch
#vanilla_radiance_cascade_batch
bilinear_fix_radiance_cascade_batch
#rotate_batch

