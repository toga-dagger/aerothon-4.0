# BOOTSTRAP:


Download Bootstrap to get the compiled CSS and JavaScript, source code, or include it with your favorite package managers like npm, RubyGems, and more.
(it will be done by our website)
we can download bootstrap in two ways.
The most preferred one is compiled version.This is because it takes less space and fast execution.And it is:

Download ready-to-use compiled code for Bootstrap v5.0.2 to easily drop into your project, which includes:

Compiled and minified CSS bundles (see CSS files comparison)
Compiled and minified JavaScript plugins (see JS files comparison)
This doesn’t include documentation, source files, or any optional JavaScript dependencies like Popper.

reference link:

https://getbootstrap.com/docs/3.4/getting-started/

# Ruby:

Requirements if windows os:
A system running Windows 10, updated to version 2004 (May 2020)
A user account with administrator privileges
Access to the command line / powershell
(click Start > type “cmd” > right-click > Run as administrator OR Start > type “powershell” >right-click > Run as administrator)

Installing Ruby Using the RubyInstaller Tool

### Step 1: Download the RubyInstaller Tool
Downloading the [Ruby installer tool.](https://drive.google.com/file/d/1umA3isoeBV_drdyRU6Az-5-Cd8krgs7V/view?usp=sharing)

Use a web browser to navigate to the Download page (linked above). If no specific version is needed, select the bolded option:

Save the file and remember its location.

### Step 2: Run the Ruby Installer
> 1. Browse to the location of the RubyInstaller tool, and double-click.

> 2. A Setup dialog launches and displays the License Agreement. Review it, tick the box to accept the agreement, then click Next.

installing ruby from install tool
> 3. The installer asks to adjust the installation location and associated file types. Leave the defaults selected, unless the system requires differently, then click Install.

Set up destination of ruby installation
> 4. The system prompts to select components. Leave the default settings, and click Next.

Select components for Ruby installation.
> 5. The Setup dialog runs for several minutes. When finished, it prompts to “Run `ridk install` to setup MSYS2 and development toolchain.” Leave this checked and click Finish.

ruby install is finished
> 6. A black command-line-style window appears, labeled RubyInstaller2 for Windows. Select the components to install. Press Enter to install the default tools.

Completing Ruby installation for Windows.
> 7. Once it finishes, press Enter again for the installer to close.

Step 3: Verify the Ruby Installation
If you have any terminal windows open, close them. Open a new terminal window, and enter the following:

```
ruby –v
```
The system displays the current version of Ruby installed on Windows 10.


# Ruby installation using linux:


## Installation on Linux
We are installing Ruby On Rails on Linux using rbenv. It is a lightweight Ruby Version Management Tool. The rbenv provides an easy installation procedure to manage various versions of Ruby, and a solid environment for developing Ruby on Rails applications.

Follow the steps given below to install Ruby on Rails using rbenv tool.

### Step 1: Install Prerequisite Dependencies
First of all, we have to install git - core and some ruby dependences that help to install Ruby on Rails. Use the following command for installing Rails dependencies using yum.
```
tp> sudo yum install -y git-core zlib zlib-devel gcc-c++ patch readline readline-devel libyaml-devel libffi-devel openssl-devel make bzip2 autoconf automake libtool bison curl sqlite-devel
```
### Step 2: Install rbenv
Now we will install rbenv and set the appropriate environment variables. Use the following set of commands to get rbenv for git repository.
```
tp> git clone git://github.com/sstephenson/rbenv.git .rbenv
tp> echo 'export PATH = "$HOME/.rbenv/bin:$PATH"' >> ~/.bash_profile
tp> echo 'eval "$(rbenv init -)"' >> ~/.bash_profile
tp> exec $SHELL
tp> git clone git://github.com/sstephenson/ruby-build.git ~/.rbenv/plugins/ruby-build
tp> echo 'export PATH = "$HOME/.rbenv/plugins/ruby-build/bin:$PATH"' << ~/.bash_profile
tp> exec $SHELL
```
### Step 3: Install Ruby
Before installing Ruby, determine which version of Ruby you want to install. We will install Ruby 2.2.3. Use the following command for installing Ruby.
```
tp> rbenv install -v 2.2.3
Use the following command for setting up the current Ruby version as default.

tp> rbenv global 2.2.3
Use the following command to verify the Ruby version.

tp> ruby -v

ruby 2.2.3p173 (2015-08-18 revivion 51636) [X86_64-linux]
Ruby provides a keyword gem for installing the supported dependencies; we call them gems. If you don't want to install the documentation for Ruby-gems, then use the following command.

tp> echo "gem: --no-document" > ~/.gemrc
Thereafter, it is better to install the Bundler gem, because it helps to manage your application dependencies. Use the following command to install bundler gem.

tp> gem install bundler
Step 4: Install Rails
Use the following command for installing Rails version 4.2.4.

tp> install rails -v 4.2.4
Use the following command to make Rails executable available.

tp> rbenv rehash
Use the following command for checking the rails version.

tp> rails -v
Output

tp> Rails 4.2.4
Ruby on Rails framework requires JavaScript Runtime Environment (Node.js) to manage the features of Rails. Next, we will see how we can use Node.js to manage Asset Pipeline which is a Rails feature.
```

### Step 5: Install JavaScript Runtime
Let us install Node.js from the Yum repository. We will take Node.js from EPEL yum repository. Use the following command to add the EPEL package to the yum repository.
```
tp> sudo yum -y install epel-release
Use the following command for installing the Node.js package.

tp> sudo yum install nodejs
Congratulations! You are now on Rails over Linux.
```

### Step 6: Install Database
By default, Rails uses sqlite3, but you may want to install MySQL, PostgreSQL, or other RDBMS. This is optional; if you have the database installed, then you may skip this step and it is not mandatory that you have a database installed to start the rails server. For this tutorial, we are using PostgreSQL database. Therefore use the following commands to install PostgreSQL.
```
tp> sudo yum install postgresql-server postgresql-contrib
Accept the prompt, by responding with a y. Use the following command to create a PostgreSQl database cluster.

tp> sudo postgresql-setup initdb
Use the following command to start and enable PostgreSQL.

tp> sudo systemctl start postgresql
tp> sudo systemctl enable postgresql
Keeping Rails Up-to-Date
Assuming you have installed Rails using RubyGems, keeping it up-to-date is relatively easy. We can use the same command in both Windows and Linux platform. Use the following command −

tp> gem update rails
Output
```

### GEM Update
This will automatically update your Rails installation. The next time you restart your application, it will pick up this latest version of Rails. While using this command, make sure you are connected to the internet.

### Installation Verification
You can verify if everything is set up according to your requirements or not. Use the following command to create a demo project.
```
tp> rails new demo
Output
```
### Rails New Demo
It will generate a demo rail project; we will discuss about it later. Currently we have to check if the environment is set up or not. Next, use the following command to run WEBrick web server on your machine.
```
tp> cd demo
tp> rails server
```

It will generate auto-code to start the server

### Rails Server
Now open your browser and type the following −

`http://localhost:3000`
It should display a message, something like, "Welcome aboard" or "Congratulations".

refernece link https://www.tutorialspoint.com/ruby-on-rails/rails-installation.htm