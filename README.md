# Main
- Run `install-tools` script to install all necessary tools
  ```shell
  $ scripts/install-tools.sh
  ```
  or for windows
  ```shell
  $ cd scripts && install-tools.cmd && cd ..
  ```


### Proto

#### Docs

- [Grpc documentation](https://grpc.io/docs/languages/go/quickstart/)
- [Protobuf documentation](https://protobuf.dev/getting-started/gotutorial/)

#### Repository

- Репозиторий halal-screen-proto содержит `.proto` файлы
- Репозиторий halal-screen-proto подключён в этот проект с использованием git subtree.
- Репозиторий halal-screen-proto добавлен как удалённый репозиторий в этом проекте с таким же названием  
  Выполнено командой:
  ```shell
    $ git remote add halal-screen-proto git@github.com:ravshan01/halal-screen-proto.git
  ```

  #### Subtree Control
  - Fetch and Pull script 
    ```shell
    # master - branch name
    $ scripts/proto-fetch-pull.sh master
    ```
    Or windows
    ```shell
    # master - branch name
    $ cd scripts && proto-fetch-pull.cmd master && cd ..
    ```
  - Add
    ```shell
    $ git subtree add --prefix proto halal-screen-proto master --squash
    ``` 
  - Fetch
    ```shell
    $ git fetch halal-screen-proto master
    ```
  - Pull
    ```shell
    $ git subtree pull --prefix proto halal-screen-proto master --squash
    ```
    ```shell
    $ :qa
    ```

#### Generate

- Protobuf
  ```shell 
  $ protoc --proto_path=proto --go_out=./proto --go_opt=paths=source_relative ./proto/blur.proto ./proto/detection.proto
  ```
- GRPC
  ```shell
   $ protoc --proto_path=proto --go-grpc_out=./proto --go-grpc_opt=paths=source_relative ./proto/blur.proto
  ```
