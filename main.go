package main

// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("main")
	// create chrome instance
	o := append(chromedp.DefaultExecAllocatorOptions[:],
		//... any options here
		chromedp.ProxyServer("http://118.99.127.22:8080"),
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), o...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 300 * time.Second)
	defer cancel()

	//u := `https://golang.org/pkg/time/`
	u := `https://www.myip.com/`
	selector := `#ip`
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
