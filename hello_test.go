package main // テストファイルは、テストしたい関数と同じパッケージに属している必要がある

import "testing"

func TestHello(t *testing.T) {
	got := Hello()         // テストしたい関数（Hello）を呼び出して、その結果を「got」という変数に入れる
	want := "Hello, world" // 期待する結果を「want」という変数に入れる

	if got != want { //「got」と「want」を比較して、同じかどうかをチェック
		t.Errorf("got %q want %q", got, want) // もし違っていたら、エラーメッセージを表示
	}
}

// テストを書く際のルール
// 1. xxx_test.goのような名前のファイルにある必要があります。
// 2. テスト関数はTestという単語で始まる必要があります。
// 3. テスト関数は1つの引数のみをとります。 t *testing.T
// 4. *testing.T 型を使うには、他のファイルの fmt と同じように import "testing" が必要。
