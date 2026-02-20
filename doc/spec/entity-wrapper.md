# Entity Wrapper Spec

Pattern for wrapping entities from external systems with CLI rendering support.

## Package Structure

```
pkg/<service>/
├── client.go                 # Client struct with http.Client and base URL
├── new.go                    # Client factory
├── <entity>.go               # Client method to fetch single entity
├── <entities>.go             # Client method to fetch entity list
├── constant/
│   └── constant.go           # Service-level constants (e.g., Host)
├── <entity>/
│   ├── <entity>.go           # Pure struct definition
│   ├── constant.go           # Fallback display constants
│   ├── format.go             # Main Format(f *option.Format) method
│   ├── format_<field>.go     # Per-field formatters (private)
│   ├── from_<source>.go      # Parsing factories for different HTML sources
│   └── new.go                # Simple factory (if needed)
└── example/
    └── read.go               # CLI usage example
```

## Example: Bookstore API

### Entity Struct (`book/book.go`)

Pure data, no client references:

```go
package book

import "example.com/bookstore/author"

type Book struct {
    Title     string
    ISBN      string
    Pages     string
    Price     string
    Author    *author.Author
    Published string
    URL       string
}
```

### Constants (`book/constant.go`)

Fallback values for missing data:

```go
package book

const (
    NoPages     = "no pages"
    NoPrice     = "no price"
    NoPublished = "no date"
)
```

### Main Formatter (`book/format.go`)

Composes field formatters into a status line:

```go
package book

import (
    "github.com/funtimecoding/go-library/pkg/console/status"
    "github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (b *Book) Format(f *option.Format) string {
    s := status.New(f).String(
        b.formatTitle(f),
        b.formatAuthor(f),
        b.formatPublished(f),
        b.formatPages(f),
        b.formatPrice(f),
    ).RawList(b)

    return s.Format()
}
```

### Field Formatter (`book/format_title.go`)

Each field has its own formatter with color support:

```go
package book

import (
    "github.com/funtimecoding/go-library/pkg/console"
    "github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (b *Book) formatTitle(f *option.Format) string {
    if f.UseColor {
        return console.Cyan("%s", b.Title)
    }

    return b.Title
}
```

### Field Formatter with Fallback (`book/format_pages.go`)

Handle missing values gracefully:

```go
package book

import (
    "github.com/funtimecoding/go-library/pkg/console"
    "github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (b *Book) formatPages(f *option.Format) string {
    if b.Pages == "" {
        if f.UseColor {
            return console.Yellow("%s", NoPages)
        }

        return NoPages
    }

    return b.Pages
}
```

### Parsing Factory (`book/from_detail_page.go`)

Parse from a detail page HTML document:

```go
package book

import (
    "strings"

    "github.com/PuerkitoBio/goquery"
    "example.com/bookstore/author"
)

func FromDetailPage(doc *goquery.Document, bookURL string) *Book {
    b := &Book{URL: bookURL}

    b.Title = strings.TrimSpace(doc.Find("h1.book-title").Text())
    b.ISBN = strings.TrimSpace(doc.Find(".isbn").Text())
    b.Pages = strings.TrimSpace(doc.Find(".page-count").Text())
    b.Price = strings.TrimSpace(doc.Find(".price").Text())
    b.Published = strings.TrimSpace(doc.Find(".pub-date").Text())

    authorLink := doc.Find(".author-link a").First()
    authorLocator, _ := authorLink.Attr("href")
    b.Author = author.New(authorLocator)

    return b
}
```

### Parsing Factory (`book/from_list_item.go`)

Parse from a list item in a search/listing page:

```go
package book

import (
    "strings"

    "github.com/PuerkitoBio/goquery"
    "example.com/bookstore/author"
)

func FromListItem(s *goquery.Selection, baseURL string, a *author.Author) *Book {
    titleLink := s.Find(".title a").First()
    title := titleLink.Text()
    bookURL, exists := titleLink.Attr("href")

    if !exists || bookURL == "" {
        return nil
    }

    if !strings.HasPrefix(bookURL, "http") {
        bookURL = baseURL + bookURL
    }

    return &Book{
        Title:     strings.TrimSpace(title),
        URL:       bookURL,
        Pages:     strings.TrimSpace(s.Find(".pages").Text()),
        Price:     strings.TrimSpace(s.Find(".price").Text()),
        Published: strings.TrimSpace(s.Find(".date").Text()),
        Author:    a,
    }
}
```

### Related Entity (`author/author.go`)

```go
package author

type Author struct {
    Name    string
    Locator string
}
```

### Related Entity Factory (`author/new.go`)

```go
package author

func New(locator string) *Author {
    return &Author{Name: LocatorToName(locator), Locator: locator}
}
```

### Client Method (`bookstore/book.go`)

Client handles HTTP, delegates parsing to entity package:

```go
package bookstore

import (
    "github.com/funtimecoding/go-library/pkg/errors"
    "github.com/funtimecoding/go-library/pkg/hypertext"
    "github.com/funtimecoding/go-library/pkg/web"
    "github.com/funtimecoding/go-library/pkg/web/constant"
    "example.com/bookstore/book"
)

func (c *Client) Book(locator string) *book.Book {
    r := web.NewGet(locator)
    r.Header.Set(constant.UserAgent, "Mozilla/5.0")
    p := web.Send(c.client, r)
    defer errors.LogClose(p.Body)
    panicOnStatus(p)

    return book.FromDetailPage(hypertext.Document(p.Body), locator)
}
```

### CLI Usage (`example/read.go`)

```go
package example

import (
    "fmt"

    "github.com/funtimecoding/go-library/pkg/argument"
    "github.com/funtimecoding/go-library/pkg/console/status/option"
    "example.com/bookstore"
)

func Read() {
    argument.ParseBind()
    client := bookstore.New()
    f := option.ExtendedColor.Copy()

    book := client.Book(argument.RequiredPositional(0, "BOOK_URL"))
    fmt.Println(book.Format(f))

    fmt.Println("More from author:")
    for _, b := range client.Books(book.Author) {
        fmt.Printf("- %s\n", b.Format(f))
    }
}
```

## Key Principles

1. **Structs are pure data** - no client references, no HTTP logic
2. **Entities own their presentation** - `Format(f)` method on the entity
3. **Client does HTTP only** - fetches response, delegates parsing
4. **Parsing is source-specific** - `from_detail_page.go`, `from_list_item.go`, etc.
5. **Color is optional** - controlled by `option.Format.UseColor`
6. **Missing data has fallbacks** - constants like `NoPages`, `NoPrice`
7. **Related entities use factories** - `author.New(locator)` extracts name from locator
