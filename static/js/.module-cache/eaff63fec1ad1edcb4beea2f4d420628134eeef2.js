var MobileMenu = React.createClass({displayName: "MobileMenu",
  render: function() {
    return (
      React.createElement("a", {href: "#my-menu", class: "btn", style: "position: relative; margin-top:1%; border: 2px solid black; border-radius: 7px; height: 32px; width:36px; padding: 6px;"}, 
        React.createElement("svg", {width: "20", height: "20"}, 
          React.createElement("path", {d: "M0,1 20,1", stroke: "#000", "stroke-width": "4"}), 
          React.createElement("path", {d: "M0,8 20,8", stroke: "#000", "stroke-width": "4"}), 
          React.createElement("path", {d: "M0,14 20,14", stroke: "#000", "stroke-width": "4"})
        )
      )
    );
  }
});

var HelloWorld = React.createClass({displayName: "HelloWorld",
  render: function() {
    return (
      React.createElement("p", null, 
        "Hello, ", React.createElement("input", {type: "text", placeholder: "Your name here"}), "!" + ' ' +
        "It is ", this.props.date.toTimeString()
      )
    );
  }
});

setInterval(function() {
  React.render(
    React.createElement(HelloWorld, {date: new Date()}),
    document.getElementById('example')
  );
}, 500);

$(document).ready(function() {
    React.render(React.createElement(MobileMenu, null), document.getElementById('menu'));
});

$(document).ready(function() {
   $("#my-menu").mmenu();
});
