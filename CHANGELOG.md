
<a name="EdgeX Random Integer Device Service (found in device-random) Changelog"></a>
## EdgeX Random Integer Device Service
[Github repository](https://github.com/edgexfoundry/device-random)

## [v2.0.0] Ireland - 2021-06-30  (Not Compatible with 1.x releases)
### Change Logs for EdgeX Dependencies
- [device-sdk-go](https://github.com/edgexfoundry/device-sdk-go/blob/v2.0.0/CHANGELOG.md)
- [go-mod-core-contracts](https://github.com/edgexfoundry/go-mod-core-contracts/blob/v2.0.0/CHANGELOG.md)

### Features ✨
- Enable using MessageBus as the default ([#f079ac5](https://github.com/edgexfoundry/device-random/commits/f079ac5))
- Add secure MessageBus capability ([#4b5482a](https://github.com/edgexfoundry/device-random/commits/4b5482a))
- Remove Logging configuration ([#ae9ff48](https://github.com/edgexfoundry/device-random/commits/ae9ff48))
### Bug Fixes 🐛
- Added missing InsecureSecrets section ([#f0ed56a](https://github.com/edgexfoundry/device-random/commits/f0ed56a))
### Code Refactoring ♻
- remove unimplemented InitCmd/RemoveCmd configuration ([#3121aad](https://github.com/edgexfoundry/device-random/commits/3121aad))
- Change PublishTopicPrefix value to be 'edgex/events/device' ([#aa0da4d](https://github.com/edgexfoundry/device-random/commits/aa0da4d))
- Update configuration for change to common ServiceInfo struct ([#f6aa74d](https://github.com/edgexfoundry/device-random/commits/f6aa74d))
    ```
    BREAKING CHANGE:
    Service configuration has changed
    ```
- Update to assign and uses new Port Assignments ([#847e930](https://github.com/edgexfoundry/device-random/commits/847e930))
    ```
    BREAKING CHANGE:
    Device Modbus default port number has changed to 59988
    ```
- Added go mod tidy to Dockerfile ([#b906c1a](https://github.com/edgexfoundry/device-random/commits/b906c1a))
- refactor: Update for new service key names and overrides for hyphen to underscore ([#e6bfccb](https://github.com/edgexfoundry/device-random/commits/e6bfccb))
    ```
    BREAKING CHANGE:
    Service key names used in configuration have changed.
    ```
- use v2 device-sdk ([#ac673e8](https://github.com/edgexfoundry/device-random/commits/ac673e8))
### Documentation 📖
- Add badges to readme ([#73b6e10](https://github.com/edgexfoundry/device-random/commits/73b6e10))
### Build 👷
- update build files for v2 ([#6aec653](https://github.com/edgexfoundry/device-random/commits/6aec653))
- update Dockerfiles to use go 1.16 ([#ecae5e3](https://github.com/edgexfoundry/device-random/commits/ecae5e3))
- update go.mod to go 1.16 ([#40e495b](https://github.com/edgexfoundry/device-random/commits/40e495b))
### Continuous Integration 🔄
- update local docker image names ([#7b55d68](https://github.com/edgexfoundry/device-random/commits/7b55d68))

<a name="v1.3.1"></a>
## [v1.3.1] - 2021-02-02
### Continuous Integration 🔄
- add semantic.yml for commit linting, update PR template to latest ([#db9b377](https://github.com/edgexfoundry/device-random/commits/db9b377))
- standardize dockerfiles ([#147cc11](https://github.com/edgexfoundry/device-random/commits/147cc11))

<a name="v1.3.0"></a>
## [v1.3.0] - 2020-11-18
### Code Refactoring ♻
- Upgrade SDK to v1.2.4-dev.34 ([#dd26f79](https://github.com/edgexfoundry/device-random/commits/dd26f79))
- update dockerfile to appropriately use ENTRYPOINT and CMD, closes[#115](https://github.com/edgexfoundry/device-random/issues/115) ([#4c9a6eb](https://github.com/edgexfoundry/device-random/commits/4c9a6eb))
### Build 👷
- Upgrade to Go1.15 ([#f50c3e8](https://github.com/edgexfoundry/device-random/commits/f50c3e8))
- **all:** Enable use of DependaBot to maintain Go dependencies ([#b13c74e](https://github.com/edgexfoundry/device-random/commits/b13c74e))

<a name="v1.2.2"></a>
## [v1.2.2] - 2020-08-19
### Doc
- update pr template to include dependencies section ([#365baf6](https://github.com/edgexfoundry/device-random/commits/365baf6))

<a name="v1.2.1"></a>
## [v1.2.1] - 2020-06-12
### Code Refactoring ♻
- upgrade SDK to v1.2.0 ([#dd2933e](https://github.com/edgexfoundry/device-random/commits/dd2933e))
### Documentation 📖
- update pr template ([#b1922a9](https://github.com/edgexfoundry/device-random/commits/b1922a9))
### Build 👷
- Update relevant files in device-random for Go 1.13. ([#05f3a3d](https://github.com/edgexfoundry/device-random/commits/05f3a3d))
