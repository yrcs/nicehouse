# the Windows `find.exe` is different from `find` in Linux bash/shell.
# to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
# changed to use /usr/bin/find (inside Cygwin64) to run find cli.
.PHONY: *
init proto-gen swag-gen clean tidy \
wire test build-amd64-app build help:
	/usr/bin/find app -mindepth 1 -maxdepth 1 -type d -print | xargs -L 1 /usr/bin/bash -c 'cd "$$0" && pwd && ${MAKE} $@'

.DEFAULT_GOAL := help