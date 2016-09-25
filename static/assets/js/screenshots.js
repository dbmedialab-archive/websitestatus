var page = require('webpage').create();
var fs = require('fs');

var stream = fs.open('./sites.json', 'r');
var json = '';
while(!stream.atEnd()) {
  json += stream.readLine(); 
}
stream.close();

var j = JSON.parse(json);
for(var i=0; i<j.length; i++) {
  var name = j[i].name;
  var url = j[i].url;
  renderImage(name, url);
}


function renderImage(name, url) {  
  page.viewportSize = {
      width: 800,
      height: 600
  };
  
  page.open(url, function() {
    page.render('../img/' + name + '.png');
  });
}