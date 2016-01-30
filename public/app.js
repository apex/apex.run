
// Backgrounds.
var backgrounds = [1,2,3,4]

// Highlighting.
hljs.initHighlighting()

// Storage.
var store = window.sessionStorage

// Background.
var background = store.getItem('background')

if (!background) {
  var i = Math.random() * backgrounds.length | 0
  background = backgrounds[i]
  console.log('setting background to %s', background)
  store.setItem('background', background)
}

// Body class for background.
var el = document.getElementById('header-overlay')
el.style.backgroundImage = 'url(public/images/'+background+'.jpg)'
