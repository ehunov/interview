### ---------------------
###	Shortcuts
### ---------------------
de-php := docker-compose exec php sh -c

### ---------------------
###	Init
### ---------------------
init: up composer-install

composer-install:
	$(de-php) 'composer install'

### ---------------------
###	Develop
### ---------------------
up:
	docker-compose up -d

down:
	docker-compose down

sh:
	$(de-php) 'bash'

test:
	$(de-php) 'vendor/bin/phpunit'
