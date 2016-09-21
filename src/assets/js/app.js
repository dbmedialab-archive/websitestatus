'use strict';
Vue.config.delimiters = ['${', '}']

var client = new EventSource('http://localhost:8080/events');
client.onmessage = function(msg) {
  var main = document.getElementById('main');
  var html = '<ul class="siteList">';
  var data = JSON.parse(msg.data);
  console.log(data);
  for(var i=0; i<data.length; i++) {
    html += '<li class="site">'
    + '<span>' + data[i].Site.Name + '</span>'
    + '<span>Status: ' + data[i].Status + '</span>'
    + '<span>Loading time: ' + data[i].ResponseTime + '</span>'
    + '</li>'
  }
  html += '</ul>';
  main.innerHTML = html;
};