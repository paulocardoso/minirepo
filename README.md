# Minirepo - Micro Artifact Manager 

Mini Repo is a simple Maven Artifact Manager Server, with Mini Repo, you can easily deploy and download your own Java binaries.

# Supported tags and respective Dockerfile links

[Docker File](https://github.com/paulocardoso/minirepo/blob/master/Dockerfile)


#What is Mini Repo?

Mini Repo is a simple Maven Artifact Manager Server, with Mini Repo, you can easily deploy and download your own Java binaries.

#QuickStart with Mini Repo Server and Docker

Here is how to get a single node Couchbase Server cluster running on Docker containers:

Step - 1 : Run Minirepo Server docker container

docker run -d --name minirepo -p 8080:8080 paulinhocru/minirepo:latest

Step - 2 : Next, visit http://localhost:8080/repo on the host machine to see the Web Console.

## Environenment Variables
SERVER_PORT : Port of the Minirepo server will run
LIB_PATH : Folder where the artefacts will be stored inside of the container e.g `/tmp/minirepo`

