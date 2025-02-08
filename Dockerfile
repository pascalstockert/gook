FROM alpine

COPY ./bin/ /usr/gook/

ENV PATH="${PATH}:/usr/gook"
