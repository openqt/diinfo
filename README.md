# diinfo
Docker Image Information Inspector

## Configuration
- Search order:
    1. /etc/diinfo/diinfo.yml
    2. ./diinfo.yml

    - Content
    ```
    registry: http://t2cp.io:5000/
    json: false
    wide: 80
    all: true

    ```

## Docker
- Get from https://hub.docker.com/r/openqt/diinfo/
    ```
    docker pull openqt/diinfo
    ```

- Run
    In the docker there is a default configuration file in /diinfo.yml,
    You should mount yourself file, as below

    ```
    $ docker run -v $PWD/diinfo.yml:/diinfo.yml diinfo ls
    1.	alpine: [latest]
    2.	hello-world: [test latest]
    3.	nginx: [latest]
    4.	redis: [2 test latest]
    ```

## Examples
- help
    ```
    $ ./diinfo
    Docker Image Information Inspector

    Usage:
      diinfo [command]

    Available Commands:
      del         Delete image(s)
      help        Help about any command
      ls          List images in docker registry
      show        Show docker image internals

    Flags:
      -h, --help              help for diinfo
      -r, --registry string   Docker registry address
      -v, --verbose           Show logs

    Use "diinfo [command] --help" for more information about a command.

    ```

- list

List all images in the repository

    ```
    $ ./diinfo ls -v
    2018/05/25 17:46:13 Configuration: /opt/shift/src/github.com/openqt/diinfo/diinfo.yml
    {
      "Verbose": true,
      "Registry": "http://t2cp.io:5000/",
      "Username": "",
      "Password": ""
    }
    2018/05/25 17:46:13 registry.ping url=http://t2cp.io:5000/v2/
    2018/05/25 17:46:13 registry.repositories url=http://t2cp.io:5000/v2/_catalog
    2018/05/25 17:46:13 [hello-world nginx redis]
    1.      hello-world
    2.      nginx
    3.      redis
    ```

- info

Show image details with layer(s) information.

    - show solid layers
    ```
    $ ./diinfo show redis
    No.	      Size	Command [/bin/sh -c]
    ---	      ----	--------------------
    1.	  30114519	#(nop) ADD file:e7ac45803c3ab9b7023933b75f5a88eda1f3edca97c7e462401860777cf312f7 ...
    2.	      2086	groupadd -r redis && useradd -r -g redis redis
    3.	    981699	set -ex; fetchDeps='ca-certificates wget'; apt-get update; apt-get install -y -- ...
    4.	   8289389	set -ex; buildDeps=' wget gcc libc6-dev make '; apt-get update; apt-get install  ...
    5.	        97	mkdir /data && chown redis:redis /data
    6.	       402	#(nop) COPY file:9c29fbe8374a97f9c2d953c9c8b7224554607eeb7a610a930844f2bec678265 ...
    ------------------
    redis image size: 37.56 MB
    ```

    - show all layers
    ```
    $ ./diinfo show redis -a
    No.	      Size	Command [/bin/sh -c]
    ---	      ----	--------------------
    1.	  30114519	#(nop) ADD file:e7ac45803c3ab9b7023933b75f5a88eda1f3edca97c7e462401860777cf312f7 ...
    2.	         0	#(nop) CMD ["bash"]
    3.	      2086	groupadd -r redis && useradd -r -g redis redis
    4.	         0	#(nop) ENV GOSU_VERSION=1.10
    5.	    981699	set -ex; fetchDeps='ca-certificates wget'; apt-get update; apt-get install -y -- ...
    6.	         0	#(nop) ENV REDIS_VERSION=4.0.6
    7.	         0	#(nop) ENV REDIS_DOWNLOAD_URL=http://download.redis.io/releases/redis-4.0.6.tar. ...
    8.	         0	#(nop) ENV REDIS_DOWNLOAD_SHA=769b5d69ec237c3e0481a262ff5306ce30db9b5c8ceb14d102 ...
    9.	   8289389	set -ex; buildDeps=' wget gcc libc6-dev make '; apt-get update; apt-get install  ...
    10.	        97	mkdir /data && chown redis:redis /data
    11.	         0	#(nop) VOLUME [/data]
    12.	         0	#(nop) WORKDIR /data
    13.	       402	#(nop) COPY file:9c29fbe8374a97f9c2d953c9c8b7224554607eeb7a610a930844f2bec678265 ...
    14.	         0	#(nop) ENTRYPOINT ["docker-entrypoint.sh"]
    15.	         0	#(nop) EXPOSE 6379/tcp
    16.	         0	#(nop) CMD ["redis-server"]
    ------------------
    redis image size: 37.56 MB
    ```

    - show all information
    ```
    $ ./diinfo info redis -j
    ---------- Manifest ----------
    {
       "schemaVersion": 2,
       "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
       "config": {
          "mediaType": "application/vnd.docker.container.image.v1+json",
          "size": 5800,
          "digest": "sha256:1e70071f4af45af2cc9e1d1300c675c1ce37ee25a8a5cef1f375db5ed461dbab"
       },
       "layers": [
          {
             "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
             "size": 30114519,
             "digest": "sha256:c4bb02b17bb4b034c95a948c99c762cf0486a45f45441a052208d7750f1b413b"
          },
          {
             "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
             "size": 2086,
             "digest": "sha256:58638acf67c5d1d65732b562d5a7f5525b9788155cc10d4cd96c3d5380fadf04"
          },
          ...
       ]
    }
    ---------- Configuration ----------
    {
      "architecture": "amd64",
      "config": {
        "Hostname": "",
        "Domainname": "",
        "User": "",
        "AttachStdin": false,
        "AttachStdout": false,
        "AttachStderr": false,
        "ExposedPorts": {
          "6379/tcp": {}
        },
        "Tty": false,
        "OpenStdin": false,
        "StdinOnce": false,
        "Env": [
          "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
          "GOSU_VERSION=1.10",
          "REDIS_VERSION=4.0.6",
          "REDIS_DOWNLOAD_URL=http://download.redis.io/releases/redis-4.0.6.tar.gz",
          "REDIS_DOWNLOAD_SHA=769b5d69ec237c3e0481a262ff5306ce30db9b5c8ceb14d1023491ca7be5f6fa"
        ],
        "Cmd": [
          "redis-server"
        ],
        "ArgsEscaped": true,
        "Image": "sha256:03e433d6dc5431c878f412f8a65064556876c6dbad3393bd2d9cfff6ee66b794",
        "Volumes": {
          "/data": {}
        },
        "WorkingDir": "/data",
        "Entrypoint": [
          "docker-entrypoint.sh"
        ],
        "OnBuild": [],
        "Labels": null
      },
      "container": "cf72799c32e80417751c6657f680d42cef1f5a03befe10074860e9fdfa8d9709",
      "container_config": {
        "Hostname": "cf72799c32e8",
        "Domainname": "",
        "User": "",
        "AttachStdin": false,
        ...
      },
      "created": "2017-12-12T07:15:58.901779813Z",
      "docker_version": "17.06.2-ce",
      "history": [
        {
          "created": "2017-12-12T01:41:34.77099551Z",
          "created_by": "/bin/sh -c #(nop) ADD file:e7ac45803c3ab9b7023933b75f5a88eda1f3edca97c7e462401860777cf312f7 in / "
        },
        {
          "created": "2017-12-12T01:41:35.030304389Z",
          "created_by": "/bin/sh -c #(nop)  CMD [\"bash\"]",
          "empty_layer": true
        },
        ...
      ],
      "os": "linux",
      "rootfs": {
        "type": "layers",
        "diff_ids": [
          "sha256:cfce7a8ae6322bbbd827e1d7b401abbc81ab1663fd20b2037895cc3eff2aec6f",
          "sha256:bf0b6dc2d2d784ee1da4ab9b62d8f058a8814cf2a781f768ebb8e6cd72514127",
          ...
        ]
      }
    }
    ```

- del

The delete command is not fully implemented by the backend registry:2.
It just deletes image manifest file.
