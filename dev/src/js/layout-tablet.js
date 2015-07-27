var HelloWorld = React.createClass({
  render: function() {
    return (
      <p>
        Hello, <input type="text" placeholder="Your name here" />!
        It is {this.props.date.toTimeString()}
      </p>
    );
  }
});

setInterval(function() {
  React.render(
    <HelloWorld date={new Date()} />,
    document.getElementById('example')
  );
}, 500);

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
