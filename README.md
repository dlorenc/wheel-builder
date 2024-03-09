# WheelBuilder: A Dagger Function To Build Python Wheels

This is a Dagger function that builds Python wheels.
It uses the `build` Python package and `wolfi` images.

You can optionally specify extra Wolfi packages to install before the build.
By default, `python`, `pip`, and `build-base` are installed.

## Usage

Build the `requests` package from GitHub into a wheel:
```
dagger call build-wheel --src=https://github.com/psf/requests export --path=./dist
```

Specify a different python version:

```
dagger call build-wheel --src=https://github.com/psf/requests --python-version=3.11 export --path=./dist
```