FROM ubuntu:14.04
RUN apt-get update
RUN apt-get install -y nginx
ADD ./nginx.conf /etc/nginx/sites-available/default
RUN service nginx restart

# automatically copies the package source,
# fetches the application dependencies
# builds the program, and configures it to run on startup.
FROM golang:onbuild
EXPOSE 8080
