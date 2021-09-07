## fd_question-box

福ちゃんの質問箱のサーバーサイド

## セットアップ

```
$ docker-compouse up
```

## API ドキュメント

https://d-undefined.github.io/fd_question-box-server/

Swagger を github pages でホスティングしています。

### 更新方法

#### ReDoc のインストール

```
npm install -g redoc-cli
```

```
$ redoc-cli bundle docs/swagger.yaml -o docs/index.html
```
