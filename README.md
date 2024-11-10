# go-micro
Bunch of loosely coupled simple microservices I created to learn more about inter-service communications

Best way to test the code out would be to pull the repo and run "make up_build", it will take care of docker container initializations.
Following that you can check out the webpage on localhost:80

Will be deploying this with kubernetes on the internet when I find time to learn more about container orchestration.

## About the individual services ->
1. auth: Uses postgres for storing user info and golang/x/crypto for hashing
2. broker: Single point of entry for requests, uses chi library for multiplexer
3. frontend: Basic frontend nothing special here
4. logger: Uses mongo for storage, can communicate with other services either through the broker or through rpc/grpc
5. listener: Not very useful at the moment, uses rabbitmq to send events to consumers, more of a learning experiment
6. mail: Uses mailhog to send emails, also not used much but can be tested nonetheless, the mail server is on port 8025
7. Caddy: Reverse proxy for when I'll deploy on the internet, not in use right now.
8. project: Not a service, stores db data, makefile and docker-compose and docker swarm file