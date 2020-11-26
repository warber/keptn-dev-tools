FROM python:3.9-slim-buster
COPY psu.sh /psu.sh
RUN chmod +x /psu.sh
RUN apt-get update && apt-get install -y wget
ENTRYPOINT ["/psu.sh"]
