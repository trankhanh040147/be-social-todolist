./build_script.sh 1
docker run -d --name social-todo -p 3000:3000 -e ITEM_SERVICE_URL="http://localhost:9091"  -e MYSQL_GORM_DB_TYPE=mysql -e MYSQL_GORM_DB_SOURCE="root:my-secret-pw@tcp(my_mysql:3306)/social-todo-list?charset=utf8mb4&parseTime=True&loc=Local" -e SECRET=iTaskSecret2024 --network social-net social-todo-service
docker run -d --name social-todo-item-service -p 9091:9091 -e GINPORT=9091  -e MYSQL_GORM_DB_TYPE=mysql -e MYSQL_GORM_DB_SOURCE="root:my-secret-pw@tcp(my_mysql:3306)/social-todo-list?charset=utf8mb4&parseTime=True&loc=Local" -e SECRET=iTaskSecret2024 --network social-net social-todo-service

# reverse proxy
docker run -d --network social-net -p 80:80 -v /var/run/docker.sock:/tmp/docker.sock:ro nginxproxy/nginx-proxy
#docker run --detach --network social-net --name nginx-proxy --publish 80:80 --volume /var/run/docker.sock:/tmp/docker.sock:ro nginxproxy/nginx-proxy:1.4

# items
docker run -d --name social-todo -p 3000:3000
-e ITEM_SERVICE_URL="http://social-todo-item:9091" -e MYSQL_GORM_DB_TYPE=mysql -e MYSQL_GORM_DB_SOURCE="root:my-secret-pw@tcp(my_mysql:3306)/social-todo-list?charset=utf8mb4&parseTime=True&loc=Local" -e SECRET=iTaskSecret2024  \
-e VIRTUAL_HOST="social-todo.localhost" -e VIRTUAL_PATH="/api/v1/items" -e VIRTUAL_PORT=3000 --network social-net --expose 3000 social-todo-service

#login
docker run -d --name social-todo-login
-e MYSQL_GORM_DB_TYPE=mysql -e MYSQL_GORM_DB_SOURCE="root:my-secret-pw@tcp(my_mysql:3306)/social-todo-list?charset=utf8mb4&parseTime=True&loc=Local" -e SECRET=iTaskSecret2024  \
-e VIRTUAL_HOST="social-todo.localhost" -e VIRTUAL_PATH="/api/v1/login" -e VIRTUAL_PORT=3000 --network social-net --expose 3000 social-todo-service

# items
docker run -d --name social-todo-item -e GINPORT=9091
-e MYSQL_GORM_DB_TYPE=mysql -e MYSQL_GORM_DB_SOURCE="root:my-secret-pw@tcp(my_mysql:3306)/social-todo-list?charset=utf8mb4&parseTime=True&loc=Local" -e SECRET=iTaskSecret2024  \
-e VIRTUAL_HOST="social-todo.localhost" -e VIRTUAL_PORT=9091 --network social-net --expose 9091 social-todo-service
