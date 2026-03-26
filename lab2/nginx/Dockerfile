FROM nginx:alpine

# Удаляем дефолтную конфигурацию
RUN rm /etc/nginx/conf.d/default.conf

# Копируем наш шаблон конфигурации
COPY templates/default.conf.template /etc/nginx/templates/

# Nginx автоматически обработает шаблоны из /etc/nginx/templates/
# и создаст конфигурационные файлы в /etc/nginx/conf.d/

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
