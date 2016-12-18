##########
GOLANG CMS
##########


Open source Content Management System based on the BeeGO framework, inspired by Django CMS.


********
Features
********

* Hierarchical categories
* Extensive support for multilingual websites.  #TODO
* Use the content blocks (placeholders) in your own Templates
* Edit content directly in the frontend on your pages.  #TODO
* Navigation rendering and extending from your apps.
* SEO friendly urls.
* Mobile support.

You can define editable areas, called placeholders, in your templates which you fill
with many different so called CMS content plugins.
A list of all the plugins will be found here:

`3rd party plugins <http://www.3party-cms.com/golang-cms/>`_ #TODO

Unable to find a suitable plugin for you needs? Writing your own is very simple.

More information on `our website <http://www.golang-cms-url.org>`_.  #TODO

***********
Quick Start
***********

To compile and run the installer execute. (you will need to use GO v1.7.4 or later)

- go get -u  github.com/dionyself/beego
- cd $GOPATH/src/github.com/dionyself/beego
- git checkout golang-cms
- go get -u  github.com/beego/bee
- go get -u  github.com/dionyself/golang-cms
- cd $GOPATH/src/github.com/dionyself/golang-cms
- bee run

Browse 127.0.0.1:8080 to see GolangCMS running.
Login details. user: test, password: test

Note: You will be running a pre-alpha version in testmode.

*************
Documentation
*************

The current state of the project 'pre-alpha' version,
Only Linux based OS are supported, please report any bug you find.
Please head over to our `documentation <http://docs.goland-cms.org/>`_ for all
the details on how to install, extend and use the golang CMS.

http://docs.golang-cms-url.org/en/latest/introduction/index.html  #TODO

****
Demo
****

You will need to run the demo locally.
If you are using Docker you can run a golang-cms instance on port 8080
just run:

- docker run -p 8080:8080 dionyself/golang-cms:demo
Browse 127.0.0.1:8080 to see the demo.

if you can't see the demo please contact me.

To run unittests.

 - goconvey $GOPATH/src/github.com/dionyself/golang-cms/tests/

Donations.

.. image:: https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif
   :height: 100px
   :width: 200 px
   :scale: 50 %
   :alt: alternate text
   :align: right
   :target: https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=L4H5TUWZTZERS
