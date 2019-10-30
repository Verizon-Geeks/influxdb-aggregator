FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/influxdb-aggregator \
    USER_UID=1001 \
    USER_NAME=influxdb-aggregator

# install operator binary
COPY build/_output/bin/influxdb-aggregator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
