'use strict';

var client = new EventSource('http://localhost:8080/events');
client.onmessage = function(msg) {
  var main = document.getElementById('main');
  var html = '<ul class="siteList">';
  var data = JSON.parse(msg.data);

  for(var i=0; i<data.length; i++) {
    html += '<li class="site">'
    + '<span>' + data[i].Site.Name + '</span>'
    + '<span>Status: ' + data[i].Status + '</span>'
    + '<span>Responsetime: ' + String(data[i].ResponseTime).substr(0, 5) + '</span>'
    + '<span>Updated: ' + data[i].Updated + '</span>'
    + '</li>'
  }
  html += '</ul>';
  main.innerHTML = html;
};