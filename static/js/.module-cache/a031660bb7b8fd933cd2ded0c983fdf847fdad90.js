var MobileMenu = React.createClass({displayName: "MobileMenu",
  render: function() {
    return (
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
   $("#my-menu").mmenu({
      // options
   }, {
      // configuration
   });
});
$(document).ready(function() {
   $("#my-menu").mmenu();
});
