FROM scratch
COPY bin/run /run
ENTRYPOINT [ "/run" ]