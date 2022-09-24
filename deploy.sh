#!/bin/bash

#fool's man CI/CD
pwd

echo "LOCAL: Build and save image"
docker build . -f docker/Dockerfile --build-arg enviro=staging -t bookstore-users-api
#docker build . -t bookstore-users-api
docker save -o image.tar bookstore-users-api

echo "LOCAL: Copy image over scp"

if ! scp -P 30005 image.tar dev@codethusiast.my.id:/home/dev
then
  echo "scp failed"
  exit
fi

echo "REMOTE STAGING"

echo "STAGING: Stop and remove container bookstore-users-api & image bookstore-users-api"
ssh dev@codethusiast.my.id -p 30005 'sudo -S docker container stop bookstore-users-api'
ssh dev@codethusiast.my.id -p 30005 'sudo -S docker container rm bookstore-users-api'
ssh dev@codethusiast.my.id -p 30005 'sudo -S docker image rm bookstore-users-api'

echo "STAGING: Run bookstore-users-api image as bookstore-users-api"
ssh dev@codethusiast.my.id -p 30005 'sudo -S docker load -i image.tar'
#ssh dev@codethusiast.my.id -p 30005 'docker run -d --net docker_job2go_test_db --mount src=/home/files,target=/upload,type=bind --name bookstore-users-api bookstore-users-api'
# ssh dev@codethusiast.my.id -p 30005 'sudo -S docker run --name bookstore-users-api -p 34040:9393 bookstore-users-api'

# create container
echo "create docker container"
ssh dev@103.186.30.178 -p 30005 'sudo -S docker container create --network mysql_backend --name bookstore-users-api -e PORT=9393 -e mysql_users_username=root -e mysql_users_password=root -e mysql_users_host=my-own-mysql -e mysql_users_schema=udemy_users_db -e INSTANCE_ID="my first instance" -p 34040:9393 bookstore-users-api'
echo "end create docker container"
# run  container
echo "run docker container"
ssh dev@103.186.30.178 -p 30005 'sudo -S docker container start bookstore-users-api'
echo "end run docker container"
# 
#ssh dev@103.186.30.178 -p 30005 'docker run -d --net docker_job2go_test_db web-server-golang --name web-server-golang'


echo "STAGING: Cleanup"
ssh dev@codethusiast.my.id -p 30005 'rm image.tar'

echo "LOCAL: Cleanup"
docker image rm bookstore-users-api

rm image.tar

echo "ok"

