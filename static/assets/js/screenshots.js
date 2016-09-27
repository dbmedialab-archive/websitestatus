var fs = require('fs');

var page = require('webpage').create();
page.settings.javascriptEnabled = false;
page.viewportSize = {
  width: 1900,
  height: 1200
};

fs.changeWorkingDirectory('../../../');
var stream = fs.open('sites.json', 'r');
var json = '';
while(!stream.atEnd()) {
  json += stream.readLine();
}
stream.close();
fs.changeWorkingDirectory('static/assets/files/');

var j = JSON.parse(json);
var URLS = [];
for(var i=0; i<j.length; i++) {
  var url = j[i].url;
  URLS.push(url);
}

var index = 1;
function handle_page(file) {
  console.log(file);
  page.open(file, function() {
    page.clipRect = {
      top: 0,
      left: 0,
      width: 1900,
      height: 1200
    };

    page.render(index++ + '.png');
    setTimeout(next_page, 5000);
  });
}

function next_page() {
  var file = URLS.shift();
  if(!file) {
    phantom.exit(0);
  }
  handle_page(file);
}

next_page();