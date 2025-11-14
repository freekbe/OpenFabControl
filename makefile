COMPOSE_FILE = docker-compose.yml
COMPOSE_COMMAND = docker-compose

COLIMA_START_OPTIONS = start

# Targets
.PHONY: up down ps logs start_colima set_context

up:
	$(COMPOSE_COMMAND) up --build

down:
	$(COMPOSE_COMMAND) down

ps:
	$(COMPOSE_COMMAND) ps

logs:
	$(COMPOSE_COMMAND) logs -f $(SERVICE_NAME)

start_colima:
	colima $(COLIMA_START_OPTIONS)

set_context:
	docker context use colima

# Default target (optional - can be "up")
all: up
