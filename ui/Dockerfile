FROM nginx

RUN mkdir /usr/share/nginx/html/css /usr/share/nginx/html/js
COPY dist/* /usr/share/nginx/html/
COPY deploy/nginx.conf /etc/nginx/conf.d/default.conf

# this feels like a hack but we'll keep it here for now if it works
RUN mv /usr/share/nginx/html/*.css /usr/share/nginx/html/css
RUN mv /usr/share/nginx/html/*.js /usr/share/nginx/html/js
RUN mv /usr/share/nginx/html/*.js.map /usr/share/nginx/html/js
