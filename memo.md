### Cloud Funtion

#### HTTP トリガー

https://cloud.google.com/functions/docs/calling/http?hl=ja&cloudshell=false#functions-calling-http-go

#### ランタイム

https://cloud.google.com/sdk/gcloud/reference/functions/deploy?hl=ja#--runtime

### Cloud Vision API (OCR)

#### 画像内のテキストを検出(ラベリング)

https://cloud.google.com/vision/docs/ocr?hl=ja

#### 画像のアップロード(リモート)

https://cloud.google.com/vision/docs/ocr?hl=ja#detect_text_in_a_remote_image

ローカルはもちろん、リモートでもできる。

#### url か Base64 でエンコードしたものを指定する必要がある。

https://cloud.google.com/vision/docs/base64?hl=ja

### ライブラリ

Go のライブラリ
https://github.com/googleapis/google-cloud-go

### Twitter API

```bash
 curl "https://api.twitter.com/2/tweets/1533234443200397312?tweet.fields=attachments&expansions=attachments.media_keys&media.fields=url" -H "Authorization: Bearer <Your Bearer Token>" | jq
```

```json
{
  "data": {
    "id": "1533234443200397312",
    "text": "雨の予報なのに良い天気 https://t.co/q1EdDB5Mep",
    "attachments": {
      "media_keys": ["3_1533234439870095360"]
    }
  },
  "includes": {
    "media": [
      {
        "media_key": "3_1533234439870095360",
        "type": "photo",
        "url": "https://pbs.twimg.com/media/FUckmYHaIAAtzId.jpg"
      }
    ]
  }
}
```
