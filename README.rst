
.. image:: https://github.com/dionyself/golang-cms/actions/workflows/test.yml/badge.svg

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

************************************
Setting up a development environment
************************************

Quickstart — Docker (recommended)
=================================

The fastest way to run the demo locally. Requires only Docker.

.. code-block:: bash

   git clone https://github.com/dionyself/golang-cms.git
   cd golang-cms
   docker compose up --build

Browse http://127.0.0.1:8080 and log in with ``user: test, password: test``.

If you don't have ``docker compose`` (Compose v1), use ``docker-compose`` instead.

Quickstart — local Go toolchain
===============================

Requires Go **1.21+** and a working C compiler (for the ``mattn/go-sqlite3`` cgo dependency).

.. code-block:: bash

   git clone https://github.com/dionyself/golang-cms.git
   cd golang-cms
   # Note: `go get` for installing executables is deprecated since Go 1.17.
   # Use `go install` for the `bee` CLI.
   go install github.com/beego/bee/v2@latest
   export PATH="$PATH:$(go env GOPATH)/bin"
   bee run

Then open http://127.0.0.1:8080.

If you prefer to manage dependencies through the standard toolchain, use:

.. code-block:: bash

   go mod download
   go run .

Troubleshooting
===============

* ``go get: command not found`` (Go 1.17+)
   ``go get`` for installing executables was removed in Go 1.17. Use ``go install pkg@version`` instead.
* ``cgo: C compiler not found`` when building ``go-sqlite3``
   Install ``gcc`` (Linux: ``apt install build-essential`` / ``dnf install gcc`` ; macOS: ``xcode-select --install`` ; Windows: TDM-GCC or MinGW-w64).
* Demo container exits immediately
   Check ``docker logs <container_id>``; the most common cause is a port conflict on 8080.
* ``bee: command not found`` after ``go install``
   ``go install`` puts the binary in ``$(go env GOPATH)/bin``. Add that to your ``PATH`` (see the snippet above).

****************
Contributing
****************

* Fork the repository and create a feature branch.
* Make sure ``go test ./...`` passes locally before opening a PR.
* Run ``gofmt -s -w .`` and (optionally) ``golangci-lint run`` for style.
* Open a pull request against the ``master`` branch.

Integration tests live in ``integration_tests/`` and are run with the ``goconvey`` UI or ``go test``.

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
+-----------+----------+-------+
| Bitcoin   |  Paypal  | XMR   |
+-----------+----------+-------+
| |bitcoin| + |paypal| + |xmr| +
+-----------+----------+-------+
