import os
import logging
from typing import List

# Set up logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

def get_files(directory: str) -> List[str]:
    """Return a list of files in the given directory."""
    return [f for f in os.listdir(directory) if os.path.isfile(os.path.join(directory, f))]

def main():
    # Define the directory to search for files
    directory = '/path/to/files'

    # Check if the directory exists
    if not os.path.exists(directory):
        logger.error(f"Directory '{directory}' does not exist.")
        return

    # Get a list of files in the directory
    files = get_files(directory)

    # Print the list of files
    logger.info(f"Found {len(files)} files in '{directory}':")
    for file in files:
        logger.info(file)

if __name__ == "__main__":
    main()