FROM stakater/base-centos:7
MAINTAINER "Stakater Team"

COPY Scaler .

ENTRYPOINT ["./Scaler"]