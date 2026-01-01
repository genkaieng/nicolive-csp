# nicolive-csp

ニコ生のコメントをリアルタイム取得するCSPです。コメントビュワーなどのクライアントアプリからCSPを子プロセスとして実行することで、ニコ生のコメントサーバーとの通信を抽象化することが出来ます。

*※ CSP = Comment Server Protocol*

## Get Started

### Installing and running Nicolive-CSP

Install global with go:

```
go install github.com/genkaieng/nicolive-csp@latest
```

Then simply run `nicolive-csp` to get started:

```
nicolive-csp <lvid>
```

## Dependencies

protobufは [n-air-app/nicolive-comment-protobuf](https://github.com/n-air-app/nicolive-comment-protobuf) から拝借

### protoのコンパイル

```
protoc --go_out=. --go_opt=module=github.com/genkaieng/nicolive-csp -I=proto proto/**/*.proto
```

## References

今回の実装の参考にしたリポジトリや記事など

- https://github.com/boxfish-jp/Nicolive-API
- https://qiita.com/DaisukeDaisuke/items/3938f245caec1e99d51e
