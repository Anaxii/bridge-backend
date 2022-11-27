echo "Pre-build steps"
aws ecr get-login-password --region us-west-2 | docker login --username AWS --password-stdin 072175416279.dkr.ecr.us-west-2.amazonaws.com
echo "Build steps"
docker build -t puffin-bridge-handler .
echo "tag"
docker tag puffin-bridge-handler:latest 072175416279.dkr.ecr.us-west-2.amazonaws.com/puffin-bridge-handler:latest
echo "push"
docker push 072175416279.dkr.ecr.us-west-2.amazonaws.com/puffin-bridge-handler:latest