### Proto
- Репозиторий halal-screen-proto содержит `.proto` файлы
- Репозиторий halal-screen-proto подключён в этот проект с использованием git subtree.
- Репозиторий halal-screen-proto добавлен как удалённый репозиторий в этом проекте с таким же названием  
  Выполнено командой:
  ```bash
    $ git remote add halal-screen-proto git@github.com:ravshan01/halal-screen-proto.git
  ```

  #### Subtree Control
    - Add
      ```bash
      $ git subtree add --prefix proto halal-screen-proto master --squash
      ``` 
    - Fetch
      ```bash
      $ git fetch halal-screen-proto master
      ```
    - Pull 
      ```bash
      $ git subtree pull --prefix proto halal-screen-proto master --squash
      ```
      ```bash
      $ :qa
      ```
    - Fetch and pull
      ```bash
      $ git fetch halal-screen-proto master && git subtree pull --prefix proto halal-screen-proto master --squash
      ```
      ```bash
      $ :qa
      ```

  #### Generate
    ```shell 
    protoc --proto_path=proto --go_out=./proto --go_opt=paths=source_relative --go-grpc_out=./proto --go-grpc_opt=paths=source_relative ./proto/blur.proto
    ```
