/*********
* Exercise: findlinks
* FindLinks prints the links found in an HTML document using recursion
*
 */


package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main(){
	//Get the root HTML element of the HTML doc in (as a Note type)
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _,link := range visit(nil, doc){
		fmt.Println(link)
	}
}


func visit(links []string, n *html.Node)[]string {
	if n.Type == html.ElementNode && n.Data == "html"{
		fmt.Println("Currently in the HTML node")
		fmt.Println("First Child: ", n.FirstChild.Data)
		if n.FirstChild.NextSibling.NextSibling == nil {
			fmt.Println("The next sibling is nil")
		} else {
			fmt.Println("The next sibling is not nil. Here it is: ", n.FirstChild.NextSibling.NextSibling.Data)
			if n.FirstChild.NextSibling.Type == html.TextNode {
				fmt.Println("Element Type: TextNode" )
			}

		}
		if n.LastChild == nil {
			fmt.Println("The last child is nil")
		} else {
			fmt.Println("The last child is not nil. Here it is: ", n.LastChild.Data)
		}
	}

	if n.Type == html.ElementNode && n.Data == "body"{
		fmt.Println("We have a body element")
	}

	if n.Type == html.ElementNode && n.Data == "a" { // if the node is an element and it's an anchor
		for _,a := range n.Attr { //loop through all of the attributes
			if a.Key == "href" { //if the attribute key is an hyperlink reference
				links = append(links, a.Val) //append the hyperlink to the links slice
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
