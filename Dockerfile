# A minimal Docker image based on Alpine Linux with a complete package index and only 5 MB in size!
FROM alpine

# expose port
EXPOSE 8585

# Add executable into image
ADD build/app /

CMD ["/app"]

# execute command when docker launches / run
#CMD ["./app"]