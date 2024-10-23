# Environment variables:
# POSTGRES_NAME: The name of the PostgreSQL database to create
# POSTGRES_USER: The username for the PostgreSQL database
# POSTGRES_PASSWORD: The password for the PostgreSQL database
# POSTGRES_IP_ADDRESS: The IP address for the PostgreSQL container
# POSTGRES_HOST_PORT: The port on the host for PostgreSQL
# POSTGRES_CONTAINER_PORT: The port inside the container for PostgreSQL

postgres:
	docker run \
	--name auth_postgres17 \
	-e POSTGRES_USER=${POSTGRES_USER} \
	-e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
	-p ${POSTGRES_HOST_PORT}:${POSTGRES_CONTAINER_PORT} \
	-h ${POSTGRES_IP_ADDRESS} \
	-v postgres_data:/var/lib/postgresql/data \
	--restart unless-stopped \
	--health-cmd="pg_isready -U ${POSTGRES_USER}" \
	--health-interval=10s \
	--health-timeout=5s \
	--health-retries=5 \
	--network=bridge \
	postgres:17-alpine
.PHONY: postgres

createdb:
	docker exec -it auth_postgres17 createdb \
	--username=${POSTGRES_USER} \
	--owner=${POSTGRES_USER} \
	${POSTGRES_NAME}
.PHONY: createdb