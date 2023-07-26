# PXTL - Pixel art tools

## Installation

If you have Go installed:
```
go install github.com/zamsyt/pxtl/cmd/pxtl@latest
```

## Commands

## `downscale`

Automatically downscale screenshots of pixel art to 1:1 scale

```
pxtl downscale screenshot.png
```

Use the `tolerance` flag to adjust the detection of tile edges (0-255. Default: 5)

```
pxtl downscale screenshot.png --tolerance 40
```
