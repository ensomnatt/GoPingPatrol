# Description

GoPingPatrol - is a service for checking health of your services. If something wrong with one of them - GoPingPatrol will notify you in telegram.

# Install

`git clone https://github.com/ensomnatt/GoPingPatrol`

After it, you should go to *config* directory and create *config.toml* file. You can find an example in the same directory.

Then, launch the project with docker compose. If you want, you can modify *docker-compose.yml*.

`docker compose up --build -d`

This will run 4 containers: checker, bot, bot's db and rabbitmq.

# Tech stack
**languages and frameworks (libraries):**
- go for checker
- typescript (puregram) for bot
- several python lines (flask) for test server (only for dev)

**db**: postgresql  
**message broker:** rabbitmq  
**deploy:** docker (compose)
