# diinfo
Docker Image Information Inspector

## Configuration
- Search order:
    1. /etc/diinfo/diinfo.yml
    2. ./diinfo.yml

- Content
    ```
    verbose: false
    registry: http://t2cp.io:5000/
    ```

## Examples

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
        {
          "created": "2017-12-12T07:11:57.725222924Z",
          "created_by": "/bin/sh -c groupadd -r redis \u0026\u0026 useradd -r -g redis redis"
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

