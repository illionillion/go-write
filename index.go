package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // ファイルを作成
    file, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("ファイルを作成できませんでした:", err)
        return
    }
    defer file.Close()

    // キーボードからの入力を受け取るためのScannerを作成
    scanner := bufio.NewScanner(os.Stdin)

    // ユーザーに入力を促すメッセージを表示
    fmt.Println("テキストを入力してください。終了するには Ctrl+D (Unix) または Ctrl+Z (Windows) を押してください。")

    // 入力をファイルに書き込む
    for scanner.Scan() {
        // 入力を取得
        input := scanner.Text()

        // 入力をファイルに書き込む
        _, err := file.WriteString(input + "\n")
        if err != nil {
            fmt.Println("書き込みエラー:", err)
            return
        }
    }

    // エラーチェック
    if err := scanner.Err(); err != nil {
        fmt.Println("読み取りエラー:", err)
        return
    }

    fmt.Println("入力をファイルに書き込みました。")
}
