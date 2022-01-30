//// Created by EldersJavas(EldersJavas&gmail.com)
//
//package main
//
//import (
//	"github.com/hajimehoshi/ebiten/v2"
//	"os"
//)
//
//func main() {
//	err := createFile()
//	if err != nil {
//		return
//	}
//	ebiten.SetWindowResizable()
//	ebiten.SetCursorMode()
//	ebiten.SetRunnableOnUnfocused()
//	}
////调用os.MkdirAll递归创建文件夹
//func createFile(filePath string)  error  {
//	if !isExist(filePath) {
//		err := os.MkdirAll(filePath,os.ModePerm)
//		return err
//	}
//	return nil
//}
//
//// 判断所给路径文件/文件夹是否存在(返回true是存在)
//func isExist(path string) bool {
//	_, err := os.Stat(path)    //os.Stat获取文件信息
//	if err != nil {
//		if os.IsExist(err) {
//			return true
//		}
//		return false
//	}
//	return true
//}