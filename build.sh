set -e
set -o pipefail
set -u

(cd src/2D && go build -o ../../CoreCascade2D)
(cd src/3D && go build -o ../../CoreCascade3D)

#goreleaser release --snapshot --clean