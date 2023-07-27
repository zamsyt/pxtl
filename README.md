# PXTL - Pixel art tools

## Installation

Download: https://github.com/zamsyt/pxtl/releases

Alternatively, installation from source:
```
go install github.com/zamsyt/pxtl/cmd/pxtl@latest
```

## Commands

### `downscale`

Automatically downscale screenshots of pixel art to 1:1 scale

```
pxtl downscale screenshot.png
```

Use the `tolerance` flag to adjust the detection of tile edges (0-255. Default: 5)

```
pxtl downscale screenshot.png --tolerance 40
```

#### Notes

Keep in mind that this functionality is currently quite limited and experimental. I am planning to add improvements. Feedback is appreciated.

- **Full repeated lines are currently removed without warning.** (TODO)
- Images are expected to have a mostly consistent grid pattern throughout. It's not ideal to have misaligned background grids or UI elements visible.
- Some images won't work with the default low tolerance. I've been able to downscale some jpeg compressed images with a tolerance around 60.
