FROM debian

# Copy the api executable to the container image
RUN mkdir /api
COPY api/build/grocery-data-api /api

# Copy the minified web app files to the container image
COPY app/build/ /var/www/html

# Install ca-certificates for SSL cert trust
RUN apt-get update \
 && apt-get install -y --no-install-recommends ca-certificates

# Install and configure nginx
RUN apt-get install nginx -y
COPY nginx.cfg /etc/nginx/sites-available/default

# Configure startup
COPY api/.env /api/.env
COPY entrypoint.sh /api/entrypoint.sh
ENTRYPOINT [ "sh", "/api/entrypoint.sh" ]
