FROM centurylink/ca-certs
MAINTAINER Albin Gilles <gilles.albin@gopex.be>
ENV REFRESHED_AT 2016-04-30

COPY . /.
EXPOSE 4030

CMD ["./server"]
