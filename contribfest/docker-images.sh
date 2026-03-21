docker pull kindest/node:v1.35.0@sha256:452d707d4862f52530247495d180205e029056831160e22870e37e3f6c1ac31f
docker pull quay.io/metallb/controller:v0.13.7
docker pull quay.io/metallb/speaker:v0.13.7
docker pull ghcr.io/npolshakova/kgateway:v1.0.0-ci1
docker pull ghcr.io/npolshakova/sds:v1.0.0-ci1
docker pull ghcr.io/npolshakova/envoy-wrapper:v1.0.0-ci1
docker pull ghcr.io/npolshakova/dummy-idp:0.0.1
docker pull ghcr.io/npolshakova/extproc-server:0.0.1
docker pull gcr.io/k8s-staging-gateway-api/echo-basic:v20240412-v1.0.0-394-g40c666fd
docker pull registry.k8s.io/coredns/coredns:v1.12.2

docker save -o node.tar kindest/node:v1.35.0@sha256:452d707d4862f52530247495d180205e029056831160e22870e37e3f6c1ac31f
docker save -o controller.tar quay.io/metallb/controller:v0.13.7
docker save -o speaker.tar quay.io/metallb/speaker:v0.13.7
docker save -o kgateway.tar ghcr.io/npolshakova/kgateway:v1.0.0-ci1
docker save -o sds.tar ghcr.io/npolshakova/sds:v1.0.0-ci1
docker save -o envoy.tar ghcr.io/npolshakova/envoy-wrapper:v1.0.0-ci1
docker save -o idp.tar ghcr.io/npolshakova/dummy-idp:0.0.1
docker save -o extproc.tar ghcr.io/npolshakova/extproc-server:0.0.1
docker save -o echo.tar gcr.io/k8s-staging-gateway-api/echo-basic:v20240412-v1.0.0-394-g40c666fd
docker save -o coredns.tar registry.k8s.io/coredns/coredns:v1.12.2
