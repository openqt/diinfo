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
    ```
    $ ./diinfo list -v
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
