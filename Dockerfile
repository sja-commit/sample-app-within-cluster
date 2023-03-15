FROM ubuntu

COPY ./sample-app-within-cluster ./sample-app-within-cluster
ENTRYPOINT [ "./sample-app-within-cluster" ]


