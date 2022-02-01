GitHub GraphQL Client
======================

## 記録
### 認証の準備
[ここ](https://docs.github.com/ja/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)にアクセスして、個人アクセストークンを発行する。パスワードの代わり。

### エンドポイント
https://api.github.com/graphql

### curlアクセス
```
curl -H "Authorization: bearer token" -X POST -d " \ 
 { \
   \"query\": \"query { viewer { login }}\" \
 } \
" https://api.github.com/graphql
{"data":{"viewer":{"login":"tagryo"}}}
```

> "query"の文字列値は、改行文字をエスケープしていなければならず、そうなっていなかった場合にはスキーマが正しくパースできません。 POSTのボディについては、外側をダブルクオートで囲み、内部のダブルクオートはエスケープしてください。

だそうです。


## Explorer
GraphiQL でのアクセス。
https://github.com/skevy/graphiql-app/releases
ここで .exe の GraphiQL アプリをダウンロード、インストール。

https://docs.github.com/ja/graphql/guides/using-the-explorer
ここのとおり、
- Edit HTTP Headers で、 Key に Authorization と Value に Bearer ${token} を設定
- Endpoint に https://api.github.com/graphql を設定

### クエリとミューテーション
クエリの構造の基本。
```
query {
  JSON objects to return
}
```
ミューテーション
```
mutation {
  mutationName(input: {MutationNameInput!}) {
    MutationNamePayload
  }
}
```
mutationName: 実行したい変更の種類
MutationNameInput: サーバーに送信したいデータ
MutationNamePayload: サーバーから返してほしいデータ

変数は variables で定義。
```
variables {
   "number_of_repos": 3
}
```
query の引数に渡す。
```
query($number_of_repos:Int!) {
    ...
```
つかうとこで指定
```
...
    reositories(last: $number_of_repos) {
        ...
    }
...
```

今回はクエリしか使わないかな。ミューテーションは書き込みの指示。

## クエリの例
mercariの runner リポジトリに対し、 main ブランチの2022-01-22以降のコミットを取得するクエリ。
```
query($refPrefix:String!) {
  repository(name: "runner", owner: "mercari") {
    name
    refs(refPrefix: $refPrefix, query:"main", first: 2) {
      edges {
        cursor
        node {
          name
          prefix
          target {
            ... on Commit {
              id
              history(since: "2022-01-22T11:35:50-05:00") {
                totalCount
                pageInfo {
                  endCursor
                  startCursor
                }
                edges {
                  cursor
                  node {
                    oid
                    messageHeadline
                    changedFiles
                    committedDate
                    author {
                      name
                      email
                      date
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

ブランチの数変わらないからこうでもいい
```
query($branchName:String!) {
  repository(name: "runner", owner: "mercari") {
    name
    ref(qualifiedName: $branchName) {
      name
      prefix
      target {
        ... on Commit {
          id
          history(since: "2022-01-22T11:35:50-05:00") {
            totalCount
            pageInfo {
              endCursor
              startCursor
            }
            edges {
              cursor
              node {
                oid
                messageHeadline
                changedFiles
                committedDate
                author {
                  name
                  email
                  date
                }
              }
            }
          }
        }
      }
    }
  }
}
```
