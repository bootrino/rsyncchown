This is a wrapper for rsync.

This repo includes the golang source code and a binary compiled for MacOS Arm64 - use it at your own risk, or compile the source code yourself - it's not hard just go do it.

The problem it solves, is when you rsync files to a target server, you want to run chown afterwards to set the user and group owners.

There seems to be options for setting ownership within rsync but macos has an ancient version of rsync and even current versions of rsync don't seem to work in any way that makes sense to me.

Fuck knows why rsync can't do this and fuck knows why Macos is not using a new version of rsync and I'm not spending any more time trying to work it out - instead I wrote this wrapper.

This is a particular problem for Jetbrains IDEs which provide no way to run a command after doing a deploy/upload via rsync.

Here is an example of usage:

`./rsyncchown  --chown=root:appusergroup:/opt/authserver -e "ssh -p 22 " --exclude=.svn --exclude=.cvs --exclude=.idea --exclude=.DS_Store --exclude=.git --exclude=.hg --exclude=.hprof --exclude=.pyc auth_server.py ubuntu@mysshhostname:/opt/authserver///auth_server.py`

The arguments are all passed intact to rsync EXCEPT for the --chown=root:appusergroup:/opt/authserver argument which is NOT passed to rsync.

Instead, --chown=root:appusergroup:/opt/authserver is not put on the rsync command line and instead is used to run a recursive chown command on the target directory via ssh, AFTER the rsync command has run.

No doubt there is better ways to do this but rsyncchown is written for my specific requirements and gets the job done that I need - don't trust it from any perspective including security or its behaviour - use at your own risk.

Here is how you would use it in Jetbrains IDEs:


![Screenshot 2024-04-24 at 10 59 41 am](https://github.com/bootrino/rsyncchown/assets/22624099/e61ab1fb-fc7c-464b-be0f-49bd9eb7478f)



