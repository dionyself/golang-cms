var MobileMenu = React.createClass({displayName: "MobileMenu",
  render: function() {
    return (
      <div id="menu">
        <a href="#my-menu" className="btn" style={{position: 'relative', marginTop: '1%', border: '2px solid black', borderRadius: 7, height: 32, width: 36, padding: 6}}>
          <svg width="20" height="20">
            <path d="M0,1 20,1" stroke="#000" stroke-width="4"/>
            <path d="M0,8 20,8" stroke="#000" stroke-width="4"/>
            <path d="M0,14 20,14" stroke="#000" stroke-width="4"/>
          </svg>
        </a>
        <nav id="my-menu">
          <ul>
            <li><a href="/">Home</a></li>
            <li><a href="/about/">About us</a>
              <ul>
                <li><a href="/about/history/">History</a></li>
                <li><a href="/about/team/">The team</a></li>
                <li><a href="/about/address/">Our address</a></li>
              </ul>
            </li>
            <li><a href="/contact/">Contact</a></li>
          </ul>
        </nav>
      </div>
    );
  }
});

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
    React.render(<MobileMenu />, document.getElementById('mobile-header-content'));
});

$(document).ready(function() {
   $("#my-menu").mmenu();
});
