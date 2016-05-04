FROM centurylink/ca-certs
MAINTAINER Albin Gilles <gilles.albin@gopex.be>
ENV REFRESHED_AT 2016-05-01

COPY . /.
EXPOSE 2015

CMD ["./server"]
