# Phosphor Templ Icons

<img src="docs/preview.png" alt="Preview">

This project provides a collection of [Phosphor Icons](https://phosphoricons.com/) wrapped as templ components for easy
use in Go-based web applications. Each icon is available as a reusable component, generated with `templ` version
**v0.3.819**, ensuring compatibility with the latest templ features.

## Features

- **Lightweight**: Each icon is rendered as an inline SVG, minimizing the need for external assets.
- **Customizable**: Modify size, variant, and other attributes directly via the `Props` struct.
- **Comprehensive**: Includes the full set of Phosphor Icons, with consistent support for all variants (e.g., `thin`,
  `light`, `bold`, `duotone`).

## Installation

1. Install the `templ` command-line tool:
   ```bash
   go install github.com/a-h/templ/cmd/templ@v0.3.819
   ```
2. Add this module to your project:
   ```bash
   go get github.com/iota-uz/icons
   ```

## Usage

### Adding Icons

Import the `icons` package and use the provided templ components with the `Props` struct for customization.

```templ
package main

import "github.com/iota-uz/icons/phosphor"

templ Page() {
  <html>
    <body>
      <h1>Example Page</h1>
      @phosphor.ArrowArcLeft(phosphor.Props{
        Variant: phosphor.DuoTone,
        Size:    "24",
        Class:   "rotate",
      })
      @phosphor.IconHeart(phosphor.Props{
        Variant: phosphor.Bold,
        Size:    "32",
        Class:   "icon-heart",
      })
    </body>
  </html>
}
```

### Props Struct

Use the `Props` struct to customize icons. The struct includes:

- **`Size`**: The size of the icon (e.g., `"24"`, `"48"`).
- **`Class`** *(optional)*: Additional CSS classes to apply to the icon.
- **`Variant`**: The style of the icon. Options:
    - `Regular`
    - `Filled`
    - `DuoTone`
    - `Thin`
    - `Bold`
    - `Light`
- **`Attributes`** *(optional)*: A `templ.Attributes` map for further customization.

### Example: Custom Attributes

```templ
@icons.IconArrowRight(icons.Props{
  Variant: icons.Filled,
  Size:    "32",
  Class:   "navigate-icon",
  Attributes: templ.Attributes{
    "data-action": "navigate",
  },
})
```

## Development

This project includes a `process.py` script to transform icons downloaded from the Phosphor Icons website into templ
components.

### Steps to Generate Components

1. **Download the Phosphor Icons ZIP file**:
    - Visit [Phosphor Icons](https://phosphoricons.com/assets/phosphor-icons.zip).
    - Download the `phosphor-icons.zip` file.

2. **Unzip the downloaded file**:
   ```bash
   unzip phosphor-icons.zip -d phosphor-icons
   ```

3. **Run the `process.py` script**:
    - Transform the unzipped assets into `.templ` files:
    ```bash
    python process.py
    ```

4. **Compile the `.templ` files into Go components**:
   ```bash
   templ generate
   ```

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.
