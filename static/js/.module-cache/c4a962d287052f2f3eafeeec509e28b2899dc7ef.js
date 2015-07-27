var MobileMenu = React.createClass({displayName: "MobileMenu",
  render: function() {
    return (
      React.createElement("div", {id: "menu"}, 
        React.createElement("a", {href: "#my-menu", className: "btn", style: {position: 'relative', marginTop: '1%', border: '2px solid black', borderRadius: 7, height: 32, width: 36, padding: 6}}, 
          React.createElement("svg", {width: "20", height: "20"}, 
            React.createElement("path", {d: "M0,1 20,1", stroke: "#000", "stroke-width": "4"}), 
            React.createElement("path", {d: "M0,8 20,8", stroke: "#000", "stroke-width": "4"}), 
            React.createElement("path", {d: "M0,14 20,14", stroke: "#000", "stroke-width": "4"})
          )
        ), 
        React.createElement("nav", {id: "my-menu"}, 
          React.createElement("ul", null, 
            React.createElement("li", null, React.createElement("a", {href: "/"}, "Home")), 
            React.createElement("li", null, React.createElement("a", {href: "/about/"}, "About us"), 
              React.createElement("ul", null, 
                React.createElement("li", null, React.createElement("a", {href: "/about/history/"}, "History")), 
                React.createElement("li", null, React.createElement("a", {href: "/about/team/"}, "The team")), 
                React.createElement("li", null, React.createElement("a", {href: "/about/address/"}, "Our address"))
              )
            ), 
            React.createElement("li", null, React.createElement("a", {href: "/contact/"}, "Contact"))
          )
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
