FROM public.ecr.aws/lambda/provided:latest
ARG SOPS_VERSION=v3.11.0
ARG SOPS_ARCHITECTURE=arm64

WORKDIR /var/runtime/
COPY bootstrap bootstrap
RUN chmod 755 bootstrap

RUN dnf --nodocs install -y jq

RUN curl -o /usr/local/bin/sops -L "https://github.com/getsops/sops/releases/download/${SOPS_VERSION}/sops-${SOPS_VERSION}.linux.${SOPS_ARCHITECTURE}" \
 && chmod +x /usr/local/bin/sops

WORKDIR /var/task/
COPY function.sh function.sh
RUN chmod 755 function.sh

# N.B. the entrypoint is part of the base image
ENTRYPOINT ["/lambda-entrypoint.sh"]
CMD ["function.handler"]
