##########
GOLANG CMS
##########


Open source Content Management System based on the BeeGO framework.

********
Features
********

* Hierarchical categories
* Extensive support for multilingual websites.  #TODO
* Use the content blocks (& placeholders) in your own Templates
* Edit content directly in the frontend on your pages.  #TODO
* Navigation rendering and extending from your apps.
* SEO friendly urls.
* Mobile support.
* Editable areas & ads support

****
Demo
****

You will need to run the demo locally (Docker engine is required).
Run a golang-cms instance on port 8080:

- docker run -p 8080:8080 dionyself/golang-cms:latest

Browse 127.0.0.1:8080 to see GolangCMS running.
Login details. user: test, password: test

- To create new articles visit http://127.0.0.1:8080/article/0/edit
- To view an article visit http://127.0.0.1:8080/article/<article_id>/show
- ex. http://127.0.0.1:8080/article/2/show

Note: You will be running a pre-alpha version in testmode.
Only Linux based OS are supported, please report any bug you find.
if you can't see the demo please contact me.

*****************************************************
Setting a development environment and/or contributing
*****************************************************

 Download, develop, compile and contribute! (requires a golang IDE, git and GO v1.19.4 or later)

- git clone https://github.com/dionyself/golang-cms.git
- cd golang-cms
- go get github.com/beego/bee/v2
- go install github.com/beego/bee/v2
- bee run

To run unittests, integration tests and Selenium Automation Testing.

 - go test ./...
 - goconvey ./integration_tests
 - webdriver ./automated_tests

.. |bitcoin| image:: https://raw.githubusercontent.com/dionyself/golang-cms/master/static/img/btttcc.png
   :height: 230px
   :width: 230 px
   :alt: Donate with Bitcoin

.. |xmr| image:: https://raw.githubusercontent.com/dionyself/golang-cms/master/static/img/xmmr.jpeg
   :height: 250px
   :width: 250 px
   :alt: Donate with Monero
   
.. |paypal| image:: https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif
   :height: 100px
   :width: 200 px
   :alt: Donate with Paypal
   :target: https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=L4H5TUWZTZERS

+------------------------------+
| Donate to this project       |
+-----------+-------+----------+
| Bitcoin   |  XMR  | Paypal   |
+-----------+-------+----------+
| |bitcoin| + |xmr| + |paypal| +
+-----------+-------+----------+
