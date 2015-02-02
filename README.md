# go-netops-tutorials
Simple Code Snippet Examples for networkstatic.net Go Blog Posts

-Post at http://networkstatic.net/golang-network-ops/

### Example usage and output

*go run hello_world.go*

	Hello World

*go-netops-tutorials # go run exec_example.go*

	2015/02/02 15:18:01 The results are ----->  PING 8.8.8.8 (8.8.8.8): 56 data bytes
	64 bytes from 8.8.8.8: icmp_seq=0 ttl=56 time=38.111 ms
	64 bytes from 8.8.8.8: icmp_seq=1 ttl=56 time=37.963 ms
	
	--- 8.8.8.8 ping statistics ---
	2 packets transmitted, 2 packets received, 0.0% packet loss
	round-trip min/avg/max/stddev = 37.963/38.037/38.111/0.074 ms
	
	2015/02/02 15:18:01 The date is --> Mon Feb  2 15:18:01 EST 2015
	
*go run resolve_dns_example.go*

	docker.io resolves to -->  162.242.195.84
	
	badH0stN@me.c0m doesnt resolve and prints a nil pointer, thats uglies -->  <nil>
	
	2015/02/02 15:18:08 lookup failed for [ badH0stN@me.c0m ]
	
	openvswitch.org resolves to -->  69.56.251.103

*go-netops-tutorials # go run struct_examples.go*

	### struct as a pointer ###
	{
		"Network": "10.1.1.0",
		"Mask": 24
	}
	### struct as a literal ###
	{
		"Network": "10.1.2.0",
		"Mask": 30
	}
	### struct as a pointer ###
	{
		"Network": "10.1.1.0",
		"Mask": 28
	}


**Note:** In these examples, each go file contains a main() function which is the starting point of execution that enables each go file to be run individually your IDE may complain about there being 

If you run:

*go build ./*

You will get the following complaints:

	 github.com/nerdalert/go-netops-tutorials
	./hello_world.go:5: main redeclared in this block
		previous declaration at ./exec_example.go:6
	./resolve_dns_example.go:9: main redeclared in this block
		previous declaration at ./hello_world.go:5
	./struct_examples.go:14: main redeclared in this block
		previous declaration at ./resolve_dns_example.go:9

To get rid of those errors you can rename some of the main() functions to a new name. Those functions could then be called from a single main() function in the root directory. I think starting out having a single file to experiment with is the simplest approach to hack for beginners.

Also, feel free to send pull requests in to learn how to do pull requests. Here is a quick how-to for pull requests. There are quite a few ways to do so, but it can be daunting for first timers so hopefully an example helps a bit. Also checkout the [Github Pull Request Help](https://help.github.com/articles/using-pull-requests/)

### How to create a Pull Request

Clone this repository

git clone git@github.com:socketplane/socketplane

	Click the fork button at the top of the page.
 
To create a fork using ssl, [setup your ssh keys](https://help.github.com/articles/generating-ssh-keys/) and then clone using the following. Be sure to replace my username w/your username. 

	git clone git@github.com:<your github username>/go-netops-tutorials.git

Alternatively you can use https to clone:

	https://github.com/nerdalert/go-netops-tutorials.git

Create a branch for your work. This creates a new branch that is basically a copy of the worksapce:


	# For bugs
	git checkout -b bug/42
	# For long-lived feature branches
	git checkout -b feature/something-cool

View the branch

	git branch -a
	
Make a change in a file and view changes you have made:

	git status
	git diff <branchname>
	example: git diff master

Make your changes and commit

	git add --all
	git commit -s
	or individually per file with:
	git add <changed file name>

Push your changes to your GitHub fork

	git push <github-user> <branch-name>

Example: git push git@github.com:<your username>/go-netops-tutorials.git new_feature/awesome

Raise a Pull Request

	git pull-request

or

Use the [Github Pull Request Help](https://help.github.com/articles/using-pull-requests/)

If you need to make changes to your pull request in response to commets etc...

Checkout your working branch

	git checkout <branch-name-with-changes>

Make changes and then commit

	git add --all
	git commit --amend
	git push --force


### Example Pull Request CLI

Here is the CLI from updating the README w/ instructions above

	go-netops-tutorials # git checkout -b update/README
	M	README.md
	Switched to a new branch 'update/README'

	go-netops-tutorials # git status
	On branch update/README
	Changes to be committed:
	  (use "git reset HEAD <file>..." to unstage)
	
		modified:   README.md
	
	go-netops-tutorials # git add README.md

	go-netops-tutorials # git commit -m"updated README"  --signoff
	[update/README 433b6b5] updated README
	 1 file changed, 133 insertions(+)

	go-netops-tutorials # git push git@github.com:nerdalert/go-netops-tutorials.git 'update/README'
	Counting objects: 5, done.
	Delta compression using up to 8 threads.
	Compressing objects: 100% (3/3), done.
	Writing objects: 100% (3/3), 1.94 KiB | 0 bytes/s, done.
	Total 3 (delta 1), reused 0 (delta 0)
	To git@github.com:nerdalert/go-netops-tutorials.git
	 * [new branch]      update/README -> update/README

	
Lastly, you can then go to your fork and Github will ask you to make a pull request. You simply compare your branch/fork to the destination branch such as 'master' and click create pull request.

Thanks for stopping by and have fun hacking! Go is really a game changer in my mind for netops dev with really good performance.
