#!/bin/sh

# Replace placeholder in index.html with actual environment variable
# If VITE_API_URL is missing, it will remain as empty string or we can set default
API_URL=${VITE_API_URL:-http://localhost:8080/api/amiibos}

echo "Injecting API URL: $API_URL"

sed -i "s|__VITE_API_URL_PLACEHOLDER__|$API_URL|g" /usr/share/nginx/html/index.html

# Start Nginx
exec nginx -g "daemon off;"
