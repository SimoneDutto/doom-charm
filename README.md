# Doom K8s

> [!NOTE]  
> Super experimental, the code is straight up bad, the charm is precisely how it comes from the tutorial.
> This can be better of course, but it is not the idea behind this project.

This is a very experimental charm to play doom in k8s via the terminal.

I've built the Doom terminal binary from https://github.com/cryptocode/terminal-doom,
embedded the binary into the go binary.

Then, following these guides:
- [Create a k8s charm](https://canonical-charmcraft.readthedocs-hosted.com/en/stable/tutorial/write-your-first-kubernetes-charm-for-a-go-app/)
- [Publish a charm](https://ops.readthedocs.io/en/latest/tutorial/from-zero-to-hero-write-your-first-kubernetes-charm/publish-your-charm-on-charmhub.html)


## Test it on your microk8s

- Install microk8s and juju: https://canonical-juju.readthedocs-hosted.com/en/latest/user/howto/manage-your-deployment/manage-your-deployment-environment/#manage-your-deployment-environment
- Deploy the charm to microk8s using juju: https://charmhub.io/doom-k8s
- [Install Ghostty](https://snapcraft.io/ghostty) or any other terminal in this [list](https://github.com/cryptocode/terminal-doom?tab=readme-ov-file#where-does-it-run)
- `juju status` -> get the unit's ip
- `ssh <ip> -p 2223` -> enjoy Doom in your terminal
