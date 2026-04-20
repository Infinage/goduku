# Use BusyBox for the smallest possible footprint
FROM --platform=$TARGETPLATFORM busybox:latest AS build-env
WORKDIR /var/www

# Copy only the UI assets. 
COPY ui/index.html .
COPY ui/main.wasm .
COPY ui/wasm_exec.js .

# This maps the .wasm extension to the required MIME type
RUN printf ".wasm:application/wasm\n" > /var/www/httpd.conf

EXPOSE 80

# Run httpd pointing to the config file with -c
# -v: verbose, -f: foreground, -h: home directory, -c: config file
CMD ["httpd", "-v", "-f", "-h", "/var/www", "-c", "/var/www/httpd.conf"]
