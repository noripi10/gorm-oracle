# gorm + Oracle DB で簡単なCRUD + Mail処理

- Oracle clientは事前にインストール

  Version: 11.2.0.4

- Table Name

```text
GORMは設定よりも規約を優先します。
デフォルトでは、GORMは主キーとしてIDを使用し、テーブル名としてsnake_cases、カラム名としてsnake_caseに構造体名を複数化し、作成/更新時間を追跡するためにCreatedAt、UpdatedAtを使用します。

モデル名 -> テーブル名
WkData -> wk_data
```

- Tablerインターフェースでテーブル名を変更することも可能