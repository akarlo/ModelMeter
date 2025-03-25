#!/usr/bin/env python3
import re
import subprocess
from decimal import Decimal


def get_ollama_total():
    """
    Calculate the total storage used by all Ollama models.
    Returns a string with the total size in the most appropriate unit.
    """
    try:
        # Run the ollama list command
        result = subprocess.run(
            ["ollama", "list"], capture_output=True, text=True, check=True
        )

        # Get the output
        output = result.stdout

        # Skip first line (header)
        lines = output.strip().split("\n")[1:]

        # Regular expression to match size values like "2.7 GB", "17 GB", "800 MB", etc.
        size_pattern = r"(\d+(?:\.\d+)?) (MB|GB|TB)"

        total_bytes = Decimal("0")

        # Process each line to extract and sum the sizes
        for line in lines:
            match = re.search(size_pattern, line)
            if match:
                size_value = Decimal(match.group(1))
                size_unit = match.group(2)

                # Convert to bytes for consistent calculation
                if size_unit == "MB":
                    total_bytes += size_value * Decimal("1000000")
                elif size_unit == "GB":
                    total_bytes += size_value * Decimal("1000000000")
                elif size_unit == "TB":
                    total_bytes += size_value * Decimal("1000000000000")

        # Convert back to the most appropriate unit
        if total_bytes >= Decimal("1000000000000"):
            return f"{(total_bytes / Decimal('1000000000000')).quantize(Decimal('0.1'))} TB"
        elif total_bytes >= Decimal("1000000000"):
            return (
                f"{(total_bytes / Decimal('1000000000')).quantize(Decimal('0.1'))} GB"
            )
        elif total_bytes >= Decimal("1000000"):
            return f"{(total_bytes / Decimal('1000000')).quantize(Decimal('0.1'))} MB"
        else:
            return f"{total_bytes} B"
    except subprocess.CalledProcessError as e:
        return f"Error running ollama list: {e}"
    except Exception as e:
        return f"Unexpected error: {e}"


if __name__ == "__main__":
    print(get_ollama_total())
