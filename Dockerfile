# Start from scratch : the docker container only contains the binary file to be executed
FROM scratch
MAINTAINER Nassim Kachroud 

# Copy the binary file and set it as entrypoint
COPY comsab-app /
ENTRYPOINT ["/comsab-app"]

# The service listens on port 5000 by default.
EXPOSE 7070
