import os
import re
from pathlib import Path
from typing import Dict, List, Optional, Tuple
from xml.etree import ElementTree as ET
from xml.etree.ElementTree import ParseError

# Input and output directories
BASE_DIR = "phosphor-icons/SVGs"  # Update with your SVG folder
OUTPUT_DIR = "./phosphor"  # Output directory for .templ files

# Ensure the output directory exists
os.makedirs(OUTPUT_DIR, exist_ok=True)


class IconProcessingError(Exception):
    """Custom exception for icon processing errors."""
    pass


def to_camel_case(snake_str: str) -> str:
    """
    Convert file names to CamelCase.

    Args:
        snake_str: The input string in kebab-case

    Returns:
        The string converted to CamelCase
    """
    if not snake_str:
        return ""

    # Remove leading/trailing hyphens and handle consecutive hyphens
    cleaned = re.sub(r'-+', '-', snake_str.strip('-'))
    components = cleaned.split("-")
    return "".join(x.title() for x in components if x)


def clean_icon_name(file_name: str) -> str:
    """
    Remove suffixes from the file name (e.g., -fill, -bold, etc.).
    More carefully handles cases where the suffix might be part of the actual name.

    Args:
        file_name: The name of the file without extension

    Returns:
        Cleaned icon name without variant suffixes
    """
    # Only remove the suffix if it's actually at the end and preceded by a hyphen
    suffixes = ["fill", "bold", "duotone", "thin", "regular", "light"]

    for suffix in suffixes:
        if file_name.endswith(f"-{suffix}"):
            return file_name[:-len(suffix) - 1]  # remove suffix and hyphen

    return file_name


def get_icon_group(file_name: str) -> str:
    """
    Get the group name based on the first word in the file name.

    Args:
        file_name: The name of the file without extension

    Returns:
        The group name for the icon
    """
    return file_name.split('-')[0]


def extract_svg_contents(svg_content: str) -> str:
    """
    Extract contents inside the <svg> tags using XML parsing.
    Removes namespace prefixes from the output.

    Args:
        svg_content: The full SVG file content

    Returns:
        The inner content of the SVG with cleaned namespace prefixes

    Raises:
        IconProcessingError: If SVG parsing fails or content is invalid
    """
    try:
        # Parse the SVG content
        root = ET.fromstring(svg_content)

        # Remove namespace prefixes from the tags
        for elem in root.iter():
            if '}' in elem.tag:
                elem.tag = elem.tag.split('}', 1)[1]

        # Get all child elements and their text without namespace prefixes
        inner_content = ''.join(ET.tostring(child, encoding='unicode') for child in root)

        # Clean up any remaining namespace declarations
        inner_content = re.sub(r'\sxmlns:[\w\d]+="[^"]*"', '', inner_content)
        inner_content = re.sub(r'\sxmlns="[^"]*"', '', inner_content)

        return inner_content.strip()
    except ParseError as e:
        raise IconProcessingError(f"Failed to parse SVG content: {e}")


def generate_templ_function(icon_name: str, variants_data: Dict[str, str]) -> str:
    """
    Generate a templ function with if/else conditions for variants.
    Regular variant (or first available variant) is used as the default case.

    Args:
        icon_name: The base name of the icon
        variants_data: Dictionary mapping variant names to SVG content

    Returns:
        Generated templ function as a string
    """
    function_name = to_camel_case(icon_name)

    # Ensure we have at least one variant
    if not variants_data:
        raise IconProcessingError(f"No variants found for icon: {icon_name}")

    # Extract regular variant for the default case (or first available variant as fallback)
    variants_for_conditions = variants_data.copy()
    default_content = variants_for_conditions.pop("Regular", None)
    if default_content is None:
        # If no Regular variant, use the first available one
        _, default_content = variants_for_conditions.popitem()

    # Generate the if/else conditions for remaining variants
    variant_cases = []
    for variant, svg_content in variants_for_conditions.items():
        if not variant_cases:  # First condition
            case = f"""
        if props.Variant == "{variant}" {{
            {svg_content}
        }}"""
        else:
            case = f"""
        else if props.Variant == "{variant}" {{
            {svg_content}
        }}"""
        variant_cases.append(case)

    # Only add conditions block if we have variants other than the default
    conditions_block = "".join(variant_cases) + """
        else {
            """ + default_content + """
        }""" if variant_cases else """
        """ + default_content

    return f"""templ {function_name}(props Props) {{
    @svg(props) {{{conditions_block}
    }}
}}
"""


def process_svg_files() -> None:
    """
    Process SVG files and generate a single templ file with all icons.

    Raises:
        IconProcessingError: If there are issues processing the files
    """
    # Map of variants to their folder names
    variants = {
        "Filled": "fill",
        "DuoTone": "duotone",
        "Bold": "bold",
        "Thin": "thin",
        "Regular": "regular",
        "Light": "light",
    }

    # Initialize tracking variables
    icon_variants: Dict[str, Dict[str, str]] = {}
    total_files_found = 0
    skipped_files = 0
    processed_files = 0
    failed_files = 0

    # Process all SVG files
    for variant, folder in variants.items():
        variant_path = Path(BASE_DIR) / folder
        if not variant_path.exists():
            print(f"Skipping non-existent folder: {folder}")
            continue

        svg_files = list(variant_path.glob("*.svg"))
        total_files_found += len(svg_files)
        print(f"Found {len(svg_files)} SVG files in {folder}")

        try:
            for file in svg_files:
                raw_icon_name = file.stem
                icon_name = clean_icon_name(raw_icon_name)

                # Check for case-sensitive duplicates
                normalized_name = icon_name.lower()
                if any(existing.lower() == normalized_name and existing != icon_name
                       for existing in icon_variants):
                    print(f"Warning: Duplicate icon name with different case found: {icon_name}")
                    skipped_files += 1
                    continue

                try:
                    with open(file, "r", encoding="utf-8") as svg_file:
                        svg_content = svg_file.read()
                except UnicodeDecodeError:
                    print(f"Warning: Unicode decode error in file {file}. Trying with latin-1 encoding.")
                    with open(file, "r", encoding="latin-1") as svg_file:
                        svg_content = svg_file.read()

                # Extract contents inside <svg>
                try:
                    svg_inner_content = extract_svg_contents(svg_content)
                    if not svg_inner_content:
                        print(f"Warning: No content found in {file}")
                        skipped_files += 1
                        continue

                    # Store inner contents for this icon and variant
                    if icon_name not in icon_variants:
                        icon_variants[icon_name] = {}
                    icon_variants[icon_name][variant] = svg_inner_content

                except IconProcessingError as e:
                    print(f"Error processing {file}: {e}")
                    failed_files += 1
                    continue

        except Exception as e:
            raise IconProcessingError(f"Error processing folder {folder}: {e}")

    if not icon_variants:
        raise IconProcessingError("No valid icons found to process")

    # Group icons by their first word
    icon_groups: Dict[str, List[str]] = {}
    for icon_name in icon_variants:
        group_name = get_icon_group(icon_name)
        if group_name not in icon_groups:
            icon_groups[group_name] = []
        icon_groups[group_name].append(icon_name)

    # Write icons grouped by their first word
    try:
        # Process each group
        for group_name, icon_names in sorted(icon_groups.items()):
            try:
                file_name = f"{group_name}.templ"
                output_path = Path(OUTPUT_DIR) / file_name

                with open(output_path, "w", encoding="utf-8") as outfile:
                    # Add package name
                    outfile.write("package phosphor\n\n")

                    # Write all icons in this group
                    for icon_name in sorted(icon_names):
                        templ_function = generate_templ_function(icon_name, icon_variants[icon_name])
                        outfile.write(templ_function)
                        outfile.write("\n\n")

                processed_files += len(icon_names)
                print(f"Generated {group_name}.templ with {len(icon_names)} icons")

            except IconProcessingError as e:
                print(f"Error generating template for group {group_name}: {e}")
                failed_files += len(icon_names)
                continue
            except IOError as e:
                print(f"Failed to write file for group {group_name}: {e}")
                failed_files += len(icon_names)
                continue

        print(f"\nProcessing Summary:")
        print(f"Total SVG files found: {total_files_found}")
        print(f"Total groups generated: {len(icon_groups)}")
        print(f"Successfully processed icons: {processed_files}")
        print(f"Skipped files: {skipped_files}")
        print(f"Failed files: {failed_files}")
        print(f"\nOutput directory: {OUTPUT_DIR}")

    except IOError as e:
        raise IconProcessingError(f"Failed to write output files: {e}")


if __name__ == "__main__":
    try:
        process_svg_files()
    except IconProcessingError as e:
        print(f"Error: {e}")
        exit(1)