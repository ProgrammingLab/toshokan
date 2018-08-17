# toshokan
プロラボの図書館

## Usage(for docker-compose)
### Install
```
$ git clone https://github.com/ProgrammingLab/toshokan.git
$ cd toshokan
$ cp toshokan.sample.toml toshokan.toml
$ cat > .env
TOSHOKAN_DB_PASSWORD="password"
^C
$ 
```
`toshokan.toml`を開き、`TOSHOKAN_DB_PASSWORD`に入れたパスワードを`dataSourceName`に設定する。

### Run the backend server
```
$ docker-compose up
```
