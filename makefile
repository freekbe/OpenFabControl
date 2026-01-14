COMPOSE_FILE = docker-compose.yml
COMPOSE_COMMAND = docker-compose

COLIMA_START_OPTIONS = start

# Read DEV from .env file
DEV := $(shell grep -E '^DEV=' .env 2>/dev/null | cut -d '=' -f2 | tr -d ' ' || echo "")

# Targets
.PHONY: up down ps logs start_colima set_context

all: up

up:
	@if [ "$(DEV)" = "true" ]; then \
		$(COMPOSE_COMMAND) --profile dev up --build; \
	else \
		$(COMPOSE_COMMAND) up --build; \
	fi

down:
	@if [ "$(DEV)" = "true" ]; then \
		$(COMPOSE_COMMAND) --profile dev down; \
	else \
		$(COMPOSE_COMMAND) down; \
	fi

ps:
	$(COMPOSE_COMMAND) ps

logs:
	$(COMPOSE_COMMAND) logs -f $(SERVICE_NAME)

start_colima:
	colima $(COLIMA_START_OPTIONS)

set_context:
	docker context use colima

reset_db: down
	docker volume rm openfabcontrol_pgdata
