package main

// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	proxy := "http://118.99.127.22:8080"
	log.Println("proxy: ", proxy)

	// create chrome instance
	o := append(chromedp.DefaultExecAllocatorOptions[:],
		//... any options here
		chromedp.ProxyServer(proxy),
	)
	// create a timeout
	ctx1, cancel1 := context.WithTimeout(context.Background(), 300 * time.Second)
	defer cancel1()
	//ctx1 := context.Background()

	ctx2, cancel2 := chromedp.NewExecAllocator(ctx1, o...)
	defer cancel2()

	ctx, cancel := chromedp.NewContext(ctx2, chromedp.WithLogf(log.Printf))
	defer cancel()



	//u := `https://golang.org/pkg/time/`
	u := `https://ifconfig.me/`
	selector := `#ip_address`
	log.Println("requesting", u)
	log.Println("selector", selector)

	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(u),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitReady(selector),
		// retrieve the value of the textarea
		chromedp.Text(selector, &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ip:\n%s", example)
}
