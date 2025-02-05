FROM alpine

COPY ./bin/ /usr/bin/gook/

ENV PATH="${PATH}:/usr/bin/gook"
