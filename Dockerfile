FROM alpine:3.11
RUN wget -O /usr/bin/omo-msa-account https://github.com/xtech-cloud/omo-msa-account/releases/download/v1.1.0/omo-msa-account
RUN chmod +x /usr/bin/omo-msa-account
ENV MSA_REGISTRY_PLUGIN consul
ENV MSA_REGISTRY_ADDRESS 127.0.0.1:8500
ENV MSA_CONFIG_DEFINE {"source":"file", "prefix":"/etc/msa/", "key":"account.yaml"}
ENV MSA_MODE release
EXPOSE 9600
ENTRYPOINT [ "omo-msa-account" ]
