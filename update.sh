sudo docker build -t handler .
docker stop $(docker ps -aq)
docker rm $(docker ps --filter status=exited -q)
sudo docker run -d -p 8080:1234 handler:latest
