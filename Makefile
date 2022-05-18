release_windows_amd64:
	git describe --tags `git rev-list --tags --max-count=1` | xargs ./scripts/release-windows-amd64.sh

release_linux_amd64:
	git describe --tags `git rev-list --tags --max-count=1` | xargs ./scripts/release-linux-amd64.sh

release_all:
	git describe --tags `git rev-list --tags --max-count=1` | xargs ./scripts/release-windows-amd64.sh
	git describe --tags `git rev-list --tags --max-count=1` | xargs ./scripts/release-linux-amd64.sh

release_all_parallel:
	make release_windows_amd64 release_linux_amd64 -j2

.PHONY: release_windows_amd64 release_linux_amd64 release_all