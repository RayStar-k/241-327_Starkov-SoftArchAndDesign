FROM nginx:alpine

RUN rm /etc/nginx/conf.d/default.conf

COPY templates/default.conf.template /etc/nginx/templates/

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
