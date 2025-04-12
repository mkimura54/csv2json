# csv2json

CSVファイルをJSONに変換する。

# Usage

```
csv2json -a -t -b sample.csv
```

## オプション

|オプション|説明|
|-|-|
|-a|変換後のJSONを整形して表示するように指定|
|-t|型を判定して適切な形式でJSONに出力するように指定|
|-b|ブランクデータ部も出力するように指定|

## 入力例
```
head1,head2,head3
string,sample1,sample2
string,,blank
string,"co,mm,a",comma
int,128,256
bool,True,false
decimal,1.5,0.9.9
```

"co""mma" のような入力は不可とする。

## 出力例
```
[
  {
    "head1": "string",
    "head2": "sample1",
    "head3": "sample2"
  },
  {
    "head1": "string",
    "head2": "",
    "head3": "blank"
  },
  {
    "head1": "string",
    "head2": "co,mm,a",
    "head3": "comma"
  },
  {
    "head1": "int",
    "head2": 128,
    "head3": 256
  },
  {
    "head1": "bool",
    "head2": true,
    "head3": false
  },
  {
    "head1": "decimal",
    "head2": 1.5,
    "head3": "0.9.9"
  }
]
```
