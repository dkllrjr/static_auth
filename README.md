# Static Authorization Middleware

`static_auth` is a simple program to serve a static site with very basic authentication. It is intended to be run in a [docker container](https://hub.docker.com/r/dkllrjr/static_auth), preferably behind some sort of reverse proxy.

## Example Usage

The program serves the static files found in the directory of your choice, assuming the directory is mounted to the `static_auth/website` directory within the docker container. In this example, we have our static files in the `website` directory within the directory where we run `docker-compose`:

```
.
+-- docker-compose.yml
+-- website
|   +-- index.html
|   +-- posts
|       +-- example_post.html
```

A `docker-compose.yml` file has been provided as an example of how to run the program:

```
version: "3.9"

services:

  static_auth:
    image: static_auth:0.1.0
    restart: unless-stopped
    environment:
      SA_SIGNUP: "${SA_SIGNUP}"
      SA_DATABASE: "${SA_DATABASE}"
      SA_URL: "${SA_URL}"
    volumes:
      - ./website:/static_auth/website
    ports:
      - 8080:8080
```

With a `.env` file, set the following variables (all are required):

- `SA_SIGNUP`: Set to `True` or `False`. This variable determines whether or not to prompt the user at startup to insert a new user.
- `SA_DATABASE`: Set to the name of your [`bbolt` database](.https://github.com/etcd-io/bbolt).
- `SA_URL`: Set to your desired URL.

Here's an example:

```
export SA_SIGNUP=True
export SA_DATABASE=users.db
export SA_URL=0.0.0.0:8080
```

Now, if you run `docker-compose up` with the above `docker-compose.yml` and `.env` files, you should be able to login with the user you created on startup of the container at `localhost:8080`.

## __Note__

This project is just at the very rough cut stage and may or may not go further. There are some leftovers from initial shaping that have yet to be cleaned up and the program is expected to be changed if I get to it in the future.

## Support

Consider supporting me if you want to see more:

<script type='text/javascript' src='https://storage.ko-fi.com/cdn/widget/Widget_2.js'></script><script type='text/javascript'>kofiwidget2.init('Support Me on Ko-fi', '#29abe0', 'K3K2JFGO4');kofiwidget2.draw();</script> 
