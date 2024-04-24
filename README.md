This is a wrapper for rsync.

The problem it solves, is when you rsync files to a target server, you want to run chown afterwards to set the user and group owners.

This is a particular problem for Jetbrains IDEs which provide no way to run a command after doing a deploy/upload.

Here is an example of usage:

`./rsyncchown  --chown=root:appusergroup:/opt/authserver -e "ssh -p 22 " --exclude=.svn --exclude=.cvs --exclude=.idea --exclude=.DS_Store --exclude=.git --exclude=.hg --exclude=.hprof --exclude=.pyc auth_server.py ubuntu@mysshhostname:/opt/authserver///auth_server.py`

The arguments are all passed intact to rsync EXCEPT for the --chown=root:appusergroup:/opt/authserver argument which is NOT passed to rsync.

Instead, --chown=root:appusergroup:/opt/authserver is not put on the rsync command line and instead is used to run a recursive chown command on the target directory via ssh, AFTER the rsync command has run.

rsyncchown is written for my specific requirements - don't trust it from any perspective including security or its behaviour - use at your own risk.



