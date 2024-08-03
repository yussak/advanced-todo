import { describe, it } from "vitest";

// TODO:テストケースは仮なので書き換える
describe("todos", () => {
  describe("index todo", () => {
    // 正常系
    it.todo("個数が正しい");
    it.todo("無い場合のテキストが出る");
    it.todo("title, bodyが表示されている");

    // 異常系
    it.todo("処理に失敗した時エラーが出る");
  });

  describe("create todo", () => {
    // 正常系
    it.todo("title, bodyを入力して追加できる");
    it.todo("バリデーションが正しく動く");

    // 異常系
    it.todo("どちらかが不正値の場合送信できない");
    it.todo("バリデーションが正しく動く");
    it.todo("処理に失敗した時エラーが出る");
  });

  describe("delete todo", () => {
    // 正常系
    it.todo("正しいidを指定して削除できる");
    it.todo("不正なidを指定して削除できない");

    // 異常系
    it.todo("idが指定されない時エラーが出る");
    it.todo("処理に失敗した時エラーが出る");
  });

  describe("update todo", () => {
    // 正常系
    it.todo("バリデーションが正しく動く");

    it.todo("正しいidを指定して更新できる");

    // 異常系
    it.todo("バリデーションが正しく動く");
    it.todo("不正なidを指定して更新できない");
    it.todo("処理に失敗した時エラーが出る");
  });

  describe("show todo", () => {
    // 正常系

    it.todo("正しいIDで取得できる");

    // 異常系
    it.todo("存在しないIDでnot found");
    it.todo("処理に失敗した時エラーが出る");
  });
});
