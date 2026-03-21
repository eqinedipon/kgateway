FROM ghcr.io/npolshakova/devcontainer:kubeconeu-2026

RUN mkdir -p docker-images

COPY node.tar docker-images/node.tar
COPY controller.tar docker-images/controller.tar
COPY speaker.tar docker-images/speaker.tar
COPY kgateway.tar docker-images/kgateway.tar
COPY sds.tar docker-images/sds.tar
COPY envoy.tar docker-images/envoy.tar
COPY idp.tar docker-images/idp.tar
COPY extproc.tar docker-images/extproc.tar
COPY echo.tar docker-images/echo.tar
COPY coredns.tar docker-images/coredns.tar


