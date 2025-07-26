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
./CoreCascade -scene penumbra -method path_tracing_parallel -output assets/path_tracing_penumbra

echo "=== Path Tracing Center ==="
./CoreCascade -scene center -method path_tracing_parallel -output assets/path_tracing_center

echo "=== Path Tracing Pinhole ==="
./CoreCascade -scene pinhole -method path_tracing_parallel -output assets/path_tracing_pinhole

echo "=== Path Tracing Shadows ==="
./CoreCascade -scene shadows -method path_tracing_parallel -output assets/path_tracing_shadows

echo "=== Path Tracing Beam ==="
./CoreCascade -scene beam -method path_tracing_parallel -output assets/path_tracing_beam

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

(cd src && go build -o ../CoreCascade)

#path_tracing_batch
vanilla_radiance_cascade_batch
#./CoreCascade -method plot
#rotate_batch

