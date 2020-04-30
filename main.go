// Command text is a chromedp example demonstrating how to extract text from a
// specific element.
package main

import (
	"context"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
)

func main() {
	
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ProxyServer(`http://64.225.24.13:3128`),
	
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create context
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.whatismyip.com/`),
		chromedp.Text(`.list-group-item py-1 font-weight-bold`, &res, chromedp.NodeVisible, chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(strings.TrimSpace(res))
}
