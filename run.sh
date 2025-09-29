set -e
set -o pipefail
set -u

rotate_batch () {

for i in $(seq 0 50);
do
    ID=$(printf '%02d\n' "$i")
    echo "=== Vanilla Shadows ${ID} ==="
    ./CoreCascade -scene shadows -time "${i}" -method vanilla_radiance_cascade -output "anim/vanilla_shadows_${ID}"
    rm "anim/vanilla_shadows_${ID}.raw"
done

}

absorption_anim_batch () {
mkdir -p absorption_anim

for i in $(seq 0 100);
do
    echo "=== Absorption Animation ${i} ==="
    time=$(bc <<< "scale=2; ${i}/100.")
    ID=$(printf '%02d\n' "$i")
    ./CoreCascade -scene absorption -method vanilla_radiance_cascade -output "absorption_anim/vanilla_absorption_${ID}" -time "${time}"
    rm "absorption_anim/vanilla_absorption_${ID}.raw"
done

}

fluid_anim_batch () {

mkdir -p fluid_anim
for i in $(seq 1 200);
do
    echo "=== Fluid Animation ${i} ==="
    ID=$(printf '%04d\n' "$i")
    ./CoreCascade -scene fluid_height -method path_tracing_3d_parallel -output "fluid_anim/fluid_${ID}" -time "${i}"
    #./CoreCascade -scene fluid -method vanilla_radiance_cascade -output "fluid_anim/fluid_${ID}" -time "${i}"
    rm "fluid_anim/fluid_${ID}.raw"
done

}

directional_anim_batch () {

mkdir -p directional_anim
for i in $(seq 0 60);
do
    echo "=== Directional Animation ${i} ==="
    ID=$(printf '%04d\n' "$i")
    ./CoreCascade -scene directional -method vanilla_radiance_cascade -output "directional_anim/directional_${ID}" -time "${i}"
    rm "directional_anim/directional_${ID}.raw"
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

light_propagation_volumes_batch () {

echo "=== Light Propagation Volumes Absorption ==="
./CoreCascade -scene absorption -method light_propagation_volumes -output "assets/lpv_absorption" -time 0
rm assets/lpv_absorption.raw

echo "=== Light Propagation Volumes Penumbra ==="
./CoreCascade -scene penumbra -method light_propagation_volumes -output assets/lpv_penumbra
rm assets/lpv_penumbra.raw

echo "=== Light Propagation Volumes Center ==="
./CoreCascade -scene center -method light_propagation_volumes -output assets/lpv_center
rm assets/lpv_center.raw

echo "=== Light Propagation Volumes Pinhole ==="
./CoreCascade -scene pinhole -method light_propagation_volumes -output assets/lpv_pinhole
rm assets/lpv_pinhole.raw

echo "=== Light Propagation Volumes Shadows ==="
./CoreCascade -scene shadows -method light_propagation_volumes -output assets/lpv_shadows
rm assets/lpv_shadows.raw

echo "=== Light Propagation Volumes Beam ==="
./CoreCascade -scene beam -method light_propagation_volumes -output assets/lpv_beam
rm assets/lpv_beam.raw

}

(cd src/2D && go build -o ../../CoreCascade2D)
(cd src/3D && go build -o ../../CoreCascade3D)

#./CoreCascade2D -scene title -method vanilla_radiance_cascade -output assets/vanilla_title
#./CoreCascade2D -scene title -method bilinear_fix_radiance_cascade -output assets/bilinear_fix_title
#./CoreCascade2D -scene fluid -method vanilla_radiance_cascade -output "fluid" -time 100
#./CoreCascade2D -scene directional -method vanilla_radiance_cascade -output "center" -time 0
#./CoreCascade2D -scene directional -method bilinear_fix_radiance_cascade -output "center" -time 0

#./CoreCascade2D -scene absorption -method vanilla_radiance_cascade -output "absorb" -time 0.5
#./CoreCascade2D -scene absorption -method bilinear_fix_radiance_cascade -output "absorb" -time 0.5
#./CoreCascade2D -scene absorption -method path_tracing_parallel -output "absorb2" -time 0.5

#./CoreCascade2D -scene title -method vanilla_radiance_cascade -output assets/vanilla_title
#./CoreCascade2D -scene shadows -method path_tracing_parallel -output shadows

#./CoreCascade3D -scene height -method vanilla_radiance_cascade -output height
#./CoreCascade3D -scene height -method path_tracing_3d_parallel -output height
#./CoreCascade3D -scene fluid_height -method path_tracing_3d_parallel -output height -time 0.

#./CoreCascade2D -method plot

#light_propagation_volumes_batch
#path_tracing_batch
#vanilla_radiance_cascade_batch
#bilinear_fix_radiance_cascade_batch
#rotate_batch
#absorption_anim_batch
#fluid_anim_batch
#directional_anim_batch

