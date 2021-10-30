package main

import (
	"fmt"
	"image"
	"sync"
)

// 并发的情况下 某一个函数只执行一次 sync.Once

var icons map[string]image.Image
var wg sync.WaitGroup
var loadIconsOnce sync.Once

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}

	//访问时 代码可能被重排如下
	//icons = make(map[string]image.Image)
	//icons["left"] = loadIcon("left.png")
	//icons["up"] = loadIcon("up.png")
	//icons["right"] = loadIcon("right.png")
	//icons["down"] = loadIcon("down.png")
}

func loadIcon(s string) image.Image {
	return nil //方便起见 给个nil
}

// Icon 被多个goroutine调用时不是并发安全的
//重排后 某个goroutine 执行了make(map[string]image.Image) 下面的还没执行，
//另外的goroutine访问时 icons != nil 直接去访问icons[name] 这时候就有问题了
func Icon(name string) image.Image {
	if icons == nil {
		loadIcons()
	}
	wg.Done()
	return icons[name]
}

// OnceIcon 是并发安全的
func OnceIcon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

func main()  {
	for i:=0;i<3;i++ {
		wg.Add(1)
		//go Icon("right")
		go OnceIcon("right")
	}
	wg.Wait()
	fmt.Println("sync once test")
}
