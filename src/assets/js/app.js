'use strict';

var React = require('react');
var ReactDOM = require('react-dom');

/*var oReq = new XMLHttpRequest();
oReq.open('GET', 'http://localhost:8080/sites');
oReq.responseType = 'json';
oReq.send();
oReq.onload = function() {*/
var client = new EventSource('http://localhost:8080')
client.onmessage = function (data) {
  console.log(data);
  /*ReactDOM.render(
    <SiteList sites={data} />,
    document.getElementById('main')
  );*/
} 
  


var SiteItem = React.createClass({
  render: function() {
    return <li className="site">
      <p>{this.props.sitename}</p>
      <p>{this.props.status}</p>
      <p>{this.props.responsetime}</p>
    </li>
  }
});

var SiteList = React.createClass({
  render: function() {        
    return (
      <ul className="siteList">
        {this.props.sites.map(function(site) {
          return <SiteItem key={site.Id} sitename={site.Name} siteurl={site.Url} />;
        })}
      </ul>
    );
  }
});
