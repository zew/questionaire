 
 
[ ![GoDoc](http://godoc.org/github.com/zew/go-questionaire?status.svg)          ](https://godoc.org/github.com/zew/go-questionaire) [ ![Travis Build](https://travis-ci.org/zew/go-questionaire.svg?branch=master)  ](https://travis-ci.org/zew/go-questionaire) [ ![Report Card](https://goreportcard.com/badge/github.com/zew/go-questionaire) ](https://goreportcard.com/report/github.com/zew/go-questionaire) [ ![code-coverage](http://gocover.io/_badge/github.com/zew/go-questionaire) ](http://gocover.io/github.com/zew/go-questionaire) 


# Go-Questionaire 

A http(s) webserver serving questionaires.

## Usage

Install and setup [golang](https://golang.org/doc/install)

    cd $HOME/go/src/github.com/zew
    go get -u github.com/zew/go-questionaire
    cd go-questionaire
    mv config-example.json  config.json  # adapt to your purposes
    mv logins-example.json  logins.json  # dito
    touch ./templates/site.css           # put your site's styles here
    go build
    ./go-questionaire                    # under windows: go-questionaire.exe

More info in [deploy on linux/unix](./static/doc/linux-instructions.md)




## Semantics

* Package `generators` contains programs for creating various questionaires.  
Questionaires are encoded as JSON file serving as data entry template.

* Different questionaires are separated by URL path.

* Survey results are pulled in by the independent command `transferrer`. 

## Data thrift

* Surveys contain no personal data - only a user ID, the questions and the answers.

* The `transferrer` pulls in the responses from an internet server.

* Once inside your organization, the results are fed into any JSON reading application.


## Technical design principles

* All content and all results are driven  
by __JSON files__.

* No database required.

* Server side validation.

* Client side JS validation is deliberately omitted;  
   [a would-be JS client lib](http://www.javascript-coder.com/html-form/form-validation.phtml)


* Package `systemtest` performs full circle filling out a questionaire and compares the 
resulting JSON file.

* Column width for any label or form element can be set individually (`ColSpanLabel` and `ColSpanControl`)

* Each label or form element can be styled additionally (`CSSLabel` and `CSSControl`)


At inception we envisioned a JSON schema validator  
and questionaire creation by directly editing of JSON files  
but that remains as elusive as it did with XML.


### Layout concept

#### Accepted Solution

We chose fixed table layout.

We need full fledged markup, since mere CSS classes such as `<div style='display: table/table-row/table-cell'` do not support colspan or rowspan functionality. 

Page width can be adjusted for each page. 
Squeezing or stretching all rows equally.
Page remains horizontally _centered_.

![Page width](./static/img/doc/page-width.png)

Each control group width can be adjusted.
The control group remains left-aligned.

![Group width](./static/img/doc/group-width.png)

Group property `OddRowsColoring` to activate alternating background

![Group width](./static/img/doc/odd-rows-coloring.png)

The table border can be set via ./templates/site.css  
`table.bordered td { myBorderCSS }`

Page property `NoNavigation` decouples the page from the
navigational workflow. 
Such pages can be jumped to by setting submit buttons to their index value.
Useful for greeting- and goodbye-pages.


#### Rejected Solutions

Inline block suffers from the disadvantage, that 
the white space between inline block elements subtracts from the total width.
The column width computation must be based on a compromise slack of i.e. 97.5 percent.

Stacking cells wit `float: left` takes away the nice vertical middle alignment of the cells.


## About Go-App-Tpl

* Go-Questionaire is based on Go-App-Tpl

* Go-App-Tpl is a number of packages for building go web applications.  


It features

  * Http router with safe settings and optional https encryption

  * Session package by Alex Edwards

  * Configurable url prefix for running multiple instances on same server:port

  * Middleware for logging, access restrictions etc.

  * Middleware catches request handler panics

  * Static file handlers
  
  * JSON config file with reloadable settings 

  * JSON logins file, also reloadable
  
  * Handlers for login, changing password, login by hash

  * Site layout template with jQuery from CDN cache; fallback to localhost 

  * Multi language strings

  * Templates having access to session and request

  * Stack of dynamic subtemplate calls 
  
  * Template pre-parsing, configurable for development or production

  * Markdown file handler, rewriting image links 
  
  * Multi language markdown files
  
  * Shell script to control application under Linux

  * CSRF and XSS defence

  

## Technical design guidelines

* Subpackaging is done by concern, neither too amorphous nor too atomic. 

* Go-App-Tpl has no "hooks" or interfaces for isolation of "framework" code.  
Just copy it and add your handlers. 

Future updates can be merged.

