Run this command to create MySQL on docker;
    
    
    docker run --name USERLOGINsql \
    -e MYSQL_ROOT_USER=root \
    -e MYSQL_ROOT_PASSWORD=123456 \
    -e MYSQL_DATABASE=USERLOGINsql \
    -p 3306:3306 \
    -d mysql:8.0.23 \
    --default-authentication-plugin=mysql_native_password \
    --character-set-server=utf8mb4 \
    --collation-server=utf8mb4_general_ci
