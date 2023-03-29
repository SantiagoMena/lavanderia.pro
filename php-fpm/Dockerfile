FROM php:8.1-fpm

RUN apt update 
# \
#     && apt install -y zlib1g-dev g++ git libicu-dev zip libzip-dev zip \
#     && apt-get install -y libpng-dev libfreetype6-dev libjpeg62-turbo-dev libgd-dev libpng-dev libjpeg-dev libfreetype6-dev \
#     && docker-php-ext-install intl opcache pdo pdo_mysql gd \
#     && pecl install apcu \
#     && docker-php-ext-enable apcu \
#     && docker-php-ext-configure zip \
#     && docker-php-ext-install zip \
#     && docker-php-ext-install exif \
#     && docker-php-ext-configure gd --with-jpeg \
#     && docker-php-ext-install -j$(nproc) gd sockets \
#     && apt-get install -y librabbitmq-dev libssh-dev \
#     && docker-php-ext-install opcache bcmath sockets \
#     && pecl install amqp \
#     && docker-php-ext-enable amqp

WORKDIR /var/www/app

# RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

EXPOSE 9000