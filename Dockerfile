FROM ghcr.io/chmikata/incident-notification:0.3.1

COPY entrypoint.sh /app
RUN chmod +x /app/entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]
CMD ["--help"]
