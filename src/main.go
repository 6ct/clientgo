package main

import (
        "github.com/webview/webview"
        // "fmt"
)

func main() {
        debug := true
        w := webview.New(debug)
        defer w.Destroy()
        w.SetTitle("Chief Client Go")
        w.SetSize(1280, 720, webview.HintNone)
        // fmt.Println(w.Window())
        w.Init(`
window.open = (...args) => {
        window.location = args[0] 
}

let adBlockCSS = "#aHolder {display:none !important;}"

document.documentElement.appendChild(document.createElement("style")).innerHTML = adBlockCSS
`)
        w.Navigate(`data:text/html,<style>
        @font-face {
                font-family: "GameFont";
                src: url("https://krunker.io/css/fonts/font2.ttf");
        }
        * {
                font-family: GameFont;
        }
        body {
                background-color: #1c1c1c;
                color: #fff;
                margin: 0;
        }
        #main {
                display: flex;
                justify-content: center;
                align-items: center;
                height: 100vh;
                width: 100vw;
        }
        button {
                padding: 15px;
                border-radius: 8px;
                color: #white;
                border: none;
                margin: 7.5px;
                font-size: 20px;
        }
        button:active {
                opacity: 0.8;
        }
        #proxyBtn {
                background-color: #ed4242;
        }
        #normalBtn {
                background-color: #b2f252;
        }
        </style>
        <div id="main">
        <a href="https://krunker.io"><button id="normalBtn">Load Game</button></a>
        <a href="https://browserfps.com"><button id="proxyBtn">Load Game Proxy</button></a>
        </div>
        `)
        w.Run()
}
